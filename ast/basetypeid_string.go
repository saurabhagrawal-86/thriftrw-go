// Code generated by "stringer -type=BaseTypeID"; DO NOT EDIT.

// Copyright (c) 2020 Uber Technologies, Inc.
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

package ast

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[BoolTypeID-1]
	_ = x[I8TypeID-2]
	_ = x[I16TypeID-3]
	_ = x[I32TypeID-4]
	_ = x[I64TypeID-5]
	_ = x[DoubleTypeID-6]
	_ = x[StringTypeID-7]
	_ = x[BinaryTypeID-8]
}

const _BaseTypeID_name = "BoolTypeIDI8TypeIDI16TypeIDI32TypeIDI64TypeIDDoubleTypeIDStringTypeIDBinaryTypeID"

var _BaseTypeID_index = [...]uint8{0, 10, 18, 27, 36, 45, 57, 69, 81}

func (i BaseTypeID) String() string {
	i -= 1
	if i < 0 || i >= BaseTypeID(len(_BaseTypeID_index)-1) {
		return "BaseTypeID(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _BaseTypeID_name[_BaseTypeID_index[i]:_BaseTypeID_index[i+1]]
}
