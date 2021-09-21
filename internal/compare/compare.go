// Copyright (c) 2021 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package compare

import (
	"fmt"

	"go.uber.org/multierr"
	"go.uber.org/thriftrw/compile"
)

type addReqError struct {
	field string
	struc string
}

func (e addReqError) Error() string {
	return fmt.Sprintf("adding a required field %s to %s is not backwards compatible", e.field, e.struc)
}

type changOptToReqError struct {
	field string
	struc string
}

func (e changOptToReqError) Error() string {
	return fmt.Sprintf("changing an optional field %s in %s to required is not backwards compatible",
		e.field, e.struc)
}

type removeMethodError struct {
	method  string
	service string
}

func (e removeMethodError) Error() string {
	return fmt.Sprintf("removing method %s in service %s is not backwards compatible", e.method, e.service)
}

type deleteServiceError struct {
	service string
}

func (e deleteServiceError) Error() string {
	return fmt.Sprintf("deleting service %s is not backwards compatible", e.service)
}

// Files compares two full file paths.
func Files(fromFile, toFile string) error {
	toModule, err := compile.Compile(toFile)
	if err != nil {
		return err
	}
	fromModule, err := compile.Compile(fromFile)
	if err != nil {
		return err
	}

	return Modules(fromModule, toModule)
}

// Modules looks for removed methods and added required fields.
func Modules(fromModule, toModule *compile.Module) error {
	err := checkRemovedMethods(fromModule, toModule)

	return multierr.Append(err, checkRequiredFields(fromModule, toModule))
}

func checkRemovedMethods(fromModule, toModule *compile.Module) error {
	return services(fromModule, toModule)
}

func checkRequiredFields(fromModule, toModule *compile.Module) error {
	for n, spec := range toModule.Types {
		fromSpec, ok := fromModule.Types[n]
		if !ok {
			// This is a new Type, which is backwards compatible.
			continue
		}
		if s, ok := spec.(*compile.StructSpec); ok {
			// Match on Type names. Here we hit a limitation, that if someone
			// renames the struct and then adds a new field, we don't really have
			// a good way of tracking it.
			if fromStructSpec, ok := fromSpec.(*compile.StructSpec); ok {
				err := structSpecs(fromStructSpec, s)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// StructSpecs compares two structs defined in a Thrift file.
func structSpecs(from, to *compile.StructSpec) error {
	fields := make(map[int16]*compile.FieldSpec, len(from.Fields))
	// Assume that these two should be compared.
	for _, f := range from.Fields {
		// Capture state of all fields here.
		fields[f.ID] = f
	}

	var errs error
	for _, toField := range to.Fields {
		if fromField, ok := fields[toField.ID]; ok {
			fromRequired := fromField.Required
			toRequired := toField.Required
			if !fromRequired && toRequired {
				errs = multierr.Append(errs, changOptToReqError{toField.ThriftName(), to.ThriftName()})
			}
		} else if toField.Required {
			errs = multierr.Append(errs, addReqError{toField.ThriftName(), to.ThriftName()})
		}
	}

	return errs
}

// Services compares two service definitions.
func services(fromModule, toModule *compile.Module) error {
	var errs error
	for n, fromService := range fromModule.Services {
		toServ, ok := toModule.Services[n]
		if !ok {
			// Service was deleted, which is not backwards compatible.
			errs = multierr.Append(errs, deleteServiceError{n})
			// Do not need to check its functions since it was deleted.

			continue
		}
		for f := range fromService.Functions {
			if _, ok := toServ.Functions[f]; !ok {
				errs = multierr.Append(errs, removeMethodError{f, n})
			}
		}
	}

	return errs
}
