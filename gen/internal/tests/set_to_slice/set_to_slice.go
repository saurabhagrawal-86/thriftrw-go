// Code generated by thriftrw v1.21.0-dev. DO NOT EDIT.
// @generated

package set_to_slice

import (
	errors "errors"
	fmt "fmt"
	multierr "go.uber.org/multierr"
	thriftreflect "go.uber.org/thriftrw/thriftreflect"
	wire "go.uber.org/thriftrw/wire"
	zapcore "go.uber.org/zap/zapcore"
	strings "strings"
)

var ConstListStringList [][]string = [][]string{
	[]string{
		"hello",
	},
	[]string{
		"world",
	},
}

var ConstStringList []string = []string{
	"hello",
}

func _MyStringList_Read(w wire.Value) (MyStringList, error) {
	var x MyStringList
	err := x.FromWire(w)
	return x, err
}

type _Set_String_sliceType_Zapper []string

// MarshalLogArray implements zapcore.ArrayMarshaler, enabling
// fast logging of _Set_String_sliceType_Zapper.
func (s _Set_String_sliceType_Zapper) MarshalLogArray(enc zapcore.ArrayEncoder) (err error) {
	for _, v := range s {
		enc.AppendString(v)
	}
	return err
}

type AnotherStringList MyStringList

// ToWire translates AnotherStringList into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
func (v AnotherStringList) ToWire() (wire.Value, error) {
	x := (MyStringList)(v)
	return x.ToWire()
}

// String returns a readable string representation of AnotherStringList.
func (v AnotherStringList) String() string {
	x := (MyStringList)(v)
	return fmt.Sprint(x)
}

// FromWire deserializes AnotherStringList from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
func (v *AnotherStringList) FromWire(w wire.Value) error {
	x, err := _MyStringList_Read(w)
	*v = (AnotherStringList)(x)
	return err
}

// Equals returns true if this AnotherStringList is equal to the provided
// AnotherStringList.
func (lhs AnotherStringList) Equals(rhs AnotherStringList) bool {
	return (MyStringList)(lhs).Equals((MyStringList)(rhs))
}

func (v AnotherStringList) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	return ((_Set_String_sliceType_Zapper)(([]string)(v))).MarshalLogArray(enc)
}

type Bar struct {
	RequiredInt32ListField             []int32        `json:"requiredInt32ListField,required"`
	OptionalStringListField            []string       `json:"optionalStringListField,omitempty"`
	RequiredTypedefStringListField     StringList     `json:"requiredTypedefStringListField,required"`
	OptionalTypedefStringListField     StringList     `json:"optionalTypedefStringListField,omitempty"`
	RequiredFooListField               []*Foo         `json:"requiredFooListField,required"`
	OptionalFooListField               []*Foo         `json:"optionalFooListField,omitempty"`
	RequiredTypedefFooListField        FooList        `json:"requiredTypedefFooListField,required"`
	OptionalTypedefFooListField        FooList        `json:"optionalTypedefFooListField,omitempty"`
	RequiredStringListListField        [][]string     `json:"requiredStringListListField,required"`
	RequiredTypedefStringListListField StringListList `json:"requiredTypedefStringListListField,required"`
}

type _Set_I32_sliceType_ValueList []int32

func (v _Set_I32_sliceType_ValueList) ForEach(f func(wire.Value) error) error {
	for _, x := range v {
		w, err := wire.NewValueI32(x), error(nil)
		if err != nil {
			return err
		}

		if err := f(w); err != nil {
			return err
		}
	}
	return nil
}

func (v _Set_I32_sliceType_ValueList) Size() int {
	return len(v)
}

func (_Set_I32_sliceType_ValueList) ValueType() wire.Type {
	return wire.TI32
}

func (_Set_I32_sliceType_ValueList) Close() {}

type _Set_String_sliceType_ValueList []string

func (v _Set_String_sliceType_ValueList) ForEach(f func(wire.Value) error) error {
	for _, x := range v {
		w, err := wire.NewValueString(x), error(nil)
		if err != nil {
			return err
		}

		if err := f(w); err != nil {
			return err
		}
	}
	return nil
}

func (v _Set_String_sliceType_ValueList) Size() int {
	return len(v)
}

func (_Set_String_sliceType_ValueList) ValueType() wire.Type {
	return wire.TBinary
}

func (_Set_String_sliceType_ValueList) Close() {}

type _Set_Foo_sliceType_ValueList []*Foo

func (v _Set_Foo_sliceType_ValueList) ForEach(f func(wire.Value) error) error {
	for _, x := range v {
		if x == nil {
			return fmt.Errorf("invalid set item: value is nil")
		}
		w, err := x.ToWire()
		if err != nil {
			return err
		}

		if err := f(w); err != nil {
			return err
		}
	}
	return nil
}

func (v _Set_Foo_sliceType_ValueList) Size() int {
	return len(v)
}

func (_Set_Foo_sliceType_ValueList) ValueType() wire.Type {
	return wire.TStruct
}

func (_Set_Foo_sliceType_ValueList) Close() {}

type _Set_Set_String_sliceType_sliceType_ValueList [][]string

func (v _Set_Set_String_sliceType_sliceType_ValueList) ForEach(f func(wire.Value) error) error {
	for _, x := range v {
		if x == nil {
			return fmt.Errorf("invalid set item: value is nil")
		}
		w, err := wire.NewValueSet(_Set_String_sliceType_ValueList(x)), error(nil)
		if err != nil {
			return err
		}

		if err := f(w); err != nil {
			return err
		}
	}
	return nil
}

func (v _Set_Set_String_sliceType_sliceType_ValueList) Size() int {
	return len(v)
}

func (_Set_Set_String_sliceType_sliceType_ValueList) ValueType() wire.Type {
	return wire.TSet
}

func (_Set_Set_String_sliceType_sliceType_ValueList) Close() {}

// ToWire translates a Bar struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Bar) ToWire() (wire.Value, error) {
	var (
		fields [10]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.RequiredInt32ListField == nil {
		return w, errors.New("field RequiredInt32ListField of Bar is required")
	}
	w, err = wire.NewValueSet(_Set_I32_sliceType_ValueList(v.RequiredInt32ListField)), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	if v.OptionalStringListField != nil {
		w, err = wire.NewValueSet(_Set_String_sliceType_ValueList(v.OptionalStringListField)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}
	if v.RequiredTypedefStringListField == nil {
		return w, errors.New("field RequiredTypedefStringListField of Bar is required")
	}
	w, err = v.RequiredTypedefStringListField.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 3, Value: w}
	i++
	if v.OptionalTypedefStringListField != nil {
		w, err = v.OptionalTypedefStringListField.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 4, Value: w}
		i++
	}
	if v.RequiredFooListField == nil {
		return w, errors.New("field RequiredFooListField of Bar is required")
	}
	w, err = wire.NewValueSet(_Set_Foo_sliceType_ValueList(v.RequiredFooListField)), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 5, Value: w}
	i++
	if v.OptionalFooListField != nil {
		w, err = wire.NewValueSet(_Set_Foo_sliceType_ValueList(v.OptionalFooListField)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 6, Value: w}
		i++
	}
	if v.RequiredTypedefFooListField == nil {
		return w, errors.New("field RequiredTypedefFooListField of Bar is required")
	}
	w, err = v.RequiredTypedefFooListField.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 7, Value: w}
	i++
	if v.OptionalTypedefFooListField != nil {
		w, err = v.OptionalTypedefFooListField.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 8, Value: w}
		i++
	}
	if v.RequiredStringListListField == nil {
		return w, errors.New("field RequiredStringListListField of Bar is required")
	}
	w, err = wire.NewValueSet(_Set_Set_String_sliceType_sliceType_ValueList(v.RequiredStringListListField)), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 9, Value: w}
	i++
	if v.RequiredTypedefStringListListField == nil {
		return w, errors.New("field RequiredTypedefStringListListField of Bar is required")
	}
	w, err = v.RequiredTypedefStringListListField.ToWire()
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 10, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Set_I32_sliceType_Read(s wire.ValueList) ([]int32, error) {
	if s.ValueType() != wire.TI32 {
		return nil, nil
	}

	o := make([]int32, 0, s.Size())
	err := s.ForEach(func(x wire.Value) error {
		i, err := x.GetI32(), error(nil)
		if err != nil {
			return err
		}

		o = append(o, i)
		return nil
	})
	s.Close()
	return o, err
}

func _Set_String_sliceType_Read(s wire.ValueList) ([]string, error) {
	if s.ValueType() != wire.TBinary {
		return nil, nil
	}

	o := make([]string, 0, s.Size())
	err := s.ForEach(func(x wire.Value) error {
		i, err := x.GetString(), error(nil)
		if err != nil {
			return err
		}

		o = append(o, i)
		return nil
	})
	s.Close()
	return o, err
}

func _StringList_Read(w wire.Value) (StringList, error) {
	var x StringList
	err := x.FromWire(w)
	return x, err
}

func _Foo_Read(w wire.Value) (*Foo, error) {
	var v Foo
	err := v.FromWire(w)
	return &v, err
}

func _Set_Foo_sliceType_Read(s wire.ValueList) ([]*Foo, error) {
	if s.ValueType() != wire.TStruct {
		return nil, nil
	}

	o := make([]*Foo, 0, s.Size())
	err := s.ForEach(func(x wire.Value) error {
		i, err := _Foo_Read(x)
		if err != nil {
			return err
		}

		o = append(o, i)
		return nil
	})
	s.Close()
	return o, err
}

func _FooList_Read(w wire.Value) (FooList, error) {
	var x FooList
	err := x.FromWire(w)
	return x, err
}

func _Set_Set_String_sliceType_sliceType_Read(s wire.ValueList) ([][]string, error) {
	if s.ValueType() != wire.TSet {
		return nil, nil
	}

	o := make([][]string, 0, s.Size())
	err := s.ForEach(func(x wire.Value) error {
		i, err := _Set_String_sliceType_Read(x.GetSet())
		if err != nil {
			return err
		}

		o = append(o, i)
		return nil
	})
	s.Close()
	return o, err
}

func _StringListList_Read(w wire.Value) (StringListList, error) {
	var x StringListList
	err := x.FromWire(w)
	return x, err
}

// FromWire deserializes a Bar struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Bar struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Bar
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Bar) FromWire(w wire.Value) error {
	var err error

	requiredInt32ListFieldIsSet := false

	requiredTypedefStringListFieldIsSet := false

	requiredFooListFieldIsSet := false

	requiredTypedefFooListFieldIsSet := false

	requiredStringListListFieldIsSet := false
	requiredTypedefStringListListFieldIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TSet {
				v.RequiredInt32ListField, err = _Set_I32_sliceType_Read(field.Value.GetSet())
				if err != nil {
					return err
				}
				requiredInt32ListFieldIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TSet {
				v.OptionalStringListField, err = _Set_String_sliceType_Read(field.Value.GetSet())
				if err != nil {
					return err
				}

			}
		case 3:
			if field.Value.Type() == wire.TSet {
				v.RequiredTypedefStringListField, err = _StringList_Read(field.Value)
				if err != nil {
					return err
				}
				requiredTypedefStringListFieldIsSet = true
			}
		case 4:
			if field.Value.Type() == wire.TSet {
				v.OptionalTypedefStringListField, err = _StringList_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 5:
			if field.Value.Type() == wire.TSet {
				v.RequiredFooListField, err = _Set_Foo_sliceType_Read(field.Value.GetSet())
				if err != nil {
					return err
				}
				requiredFooListFieldIsSet = true
			}
		case 6:
			if field.Value.Type() == wire.TSet {
				v.OptionalFooListField, err = _Set_Foo_sliceType_Read(field.Value.GetSet())
				if err != nil {
					return err
				}

			}
		case 7:
			if field.Value.Type() == wire.TSet {
				v.RequiredTypedefFooListField, err = _FooList_Read(field.Value)
				if err != nil {
					return err
				}
				requiredTypedefFooListFieldIsSet = true
			}
		case 8:
			if field.Value.Type() == wire.TSet {
				v.OptionalTypedefFooListField, err = _FooList_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 9:
			if field.Value.Type() == wire.TSet {
				v.RequiredStringListListField, err = _Set_Set_String_sliceType_sliceType_Read(field.Value.GetSet())
				if err != nil {
					return err
				}
				requiredStringListListFieldIsSet = true
			}
		case 10:
			if field.Value.Type() == wire.TSet {
				v.RequiredTypedefStringListListField, err = _StringListList_Read(field.Value)
				if err != nil {
					return err
				}
				requiredTypedefStringListListFieldIsSet = true
			}
		}
	}

	if !requiredInt32ListFieldIsSet {
		return errors.New("field RequiredInt32ListField of Bar is required")
	}

	if !requiredTypedefStringListFieldIsSet {
		return errors.New("field RequiredTypedefStringListField of Bar is required")
	}

	if !requiredFooListFieldIsSet {
		return errors.New("field RequiredFooListField of Bar is required")
	}

	if !requiredTypedefFooListFieldIsSet {
		return errors.New("field RequiredTypedefFooListField of Bar is required")
	}

	if !requiredStringListListFieldIsSet {
		return errors.New("field RequiredStringListListField of Bar is required")
	}

	if !requiredTypedefStringListListFieldIsSet {
		return errors.New("field RequiredTypedefStringListListField of Bar is required")
	}

	return nil
}

// String returns a readable string representation of a Bar
// struct.
func (v *Bar) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [10]string
	i := 0
	fields[i] = fmt.Sprintf("RequiredInt32ListField: %v", v.RequiredInt32ListField)
	i++
	if v.OptionalStringListField != nil {
		fields[i] = fmt.Sprintf("OptionalStringListField: %v", v.OptionalStringListField)
		i++
	}
	fields[i] = fmt.Sprintf("RequiredTypedefStringListField: %v", v.RequiredTypedefStringListField)
	i++
	if v.OptionalTypedefStringListField != nil {
		fields[i] = fmt.Sprintf("OptionalTypedefStringListField: %v", v.OptionalTypedefStringListField)
		i++
	}
	fields[i] = fmt.Sprintf("RequiredFooListField: %v", v.RequiredFooListField)
	i++
	if v.OptionalFooListField != nil {
		fields[i] = fmt.Sprintf("OptionalFooListField: %v", v.OptionalFooListField)
		i++
	}
	fields[i] = fmt.Sprintf("RequiredTypedefFooListField: %v", v.RequiredTypedefFooListField)
	i++
	if v.OptionalTypedefFooListField != nil {
		fields[i] = fmt.Sprintf("OptionalTypedefFooListField: %v", v.OptionalTypedefFooListField)
		i++
	}
	fields[i] = fmt.Sprintf("RequiredStringListListField: %v", v.RequiredStringListListField)
	i++
	fields[i] = fmt.Sprintf("RequiredTypedefStringListListField: %v", v.RequiredTypedefStringListListField)
	i++

	return fmt.Sprintf("Bar{%v}", strings.Join(fields[:i], ", "))
}

func _Set_I32_sliceType_Equals(lhs, rhs []int32) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for _, x := range lhs {
		ok := false
		for _, y := range rhs {
			if x == y {
				ok = true
				break
			}
		}
		if !ok {
			return false
		}
	}

	return true
}

func _Set_String_sliceType_Equals(lhs, rhs []string) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for _, x := range lhs {
		ok := false
		for _, y := range rhs {
			if x == y {
				ok = true
				break
			}
		}
		if !ok {
			return false
		}
	}

	return true
}

func _Set_Foo_sliceType_Equals(lhs, rhs []*Foo) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for _, x := range lhs {
		ok := false
		for _, y := range rhs {
			if x.Equals(y) {
				ok = true
				break
			}
		}
		if !ok {
			return false
		}
	}

	return true
}

func _Set_Set_String_sliceType_sliceType_Equals(lhs, rhs [][]string) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for _, x := range lhs {
		ok := false
		for _, y := range rhs {
			if _Set_String_sliceType_Equals(x, y) {
				ok = true
				break
			}
		}
		if !ok {
			return false
		}
	}

	return true
}

// Equals returns true if all the fields of this Bar match the
// provided Bar.
//
// This function performs a deep comparison.
func (v *Bar) Equals(rhs *Bar) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !_Set_I32_sliceType_Equals(v.RequiredInt32ListField, rhs.RequiredInt32ListField) {
		return false
	}
	if !((v.OptionalStringListField == nil && rhs.OptionalStringListField == nil) || (v.OptionalStringListField != nil && rhs.OptionalStringListField != nil && _Set_String_sliceType_Equals(v.OptionalStringListField, rhs.OptionalStringListField))) {
		return false
	}
	if !v.RequiredTypedefStringListField.Equals(rhs.RequiredTypedefStringListField) {
		return false
	}
	if !((v.OptionalTypedefStringListField == nil && rhs.OptionalTypedefStringListField == nil) || (v.OptionalTypedefStringListField != nil && rhs.OptionalTypedefStringListField != nil && v.OptionalTypedefStringListField.Equals(rhs.OptionalTypedefStringListField))) {
		return false
	}
	if !_Set_Foo_sliceType_Equals(v.RequiredFooListField, rhs.RequiredFooListField) {
		return false
	}
	if !((v.OptionalFooListField == nil && rhs.OptionalFooListField == nil) || (v.OptionalFooListField != nil && rhs.OptionalFooListField != nil && _Set_Foo_sliceType_Equals(v.OptionalFooListField, rhs.OptionalFooListField))) {
		return false
	}
	if !v.RequiredTypedefFooListField.Equals(rhs.RequiredTypedefFooListField) {
		return false
	}
	if !((v.OptionalTypedefFooListField == nil && rhs.OptionalTypedefFooListField == nil) || (v.OptionalTypedefFooListField != nil && rhs.OptionalTypedefFooListField != nil && v.OptionalTypedefFooListField.Equals(rhs.OptionalTypedefFooListField))) {
		return false
	}
	if !_Set_Set_String_sliceType_sliceType_Equals(v.RequiredStringListListField, rhs.RequiredStringListListField) {
		return false
	}
	if !v.RequiredTypedefStringListListField.Equals(rhs.RequiredTypedefStringListListField) {
		return false
	}

	return true
}

type _Set_I32_sliceType_Zapper []int32

// MarshalLogArray implements zapcore.ArrayMarshaler, enabling
// fast logging of _Set_I32_sliceType_Zapper.
func (s _Set_I32_sliceType_Zapper) MarshalLogArray(enc zapcore.ArrayEncoder) (err error) {
	for _, v := range s {
		enc.AppendInt32(v)
	}
	return err
}

type _Set_Foo_sliceType_Zapper []*Foo

// MarshalLogArray implements zapcore.ArrayMarshaler, enabling
// fast logging of _Set_Foo_sliceType_Zapper.
func (s _Set_Foo_sliceType_Zapper) MarshalLogArray(enc zapcore.ArrayEncoder) (err error) {
	for _, v := range s {
		err = multierr.Append(err, enc.AppendObject(v))
	}
	return err
}

type _Set_Set_String_sliceType_sliceType_Zapper [][]string

// MarshalLogArray implements zapcore.ArrayMarshaler, enabling
// fast logging of _Set_Set_String_sliceType_sliceType_Zapper.
func (s _Set_Set_String_sliceType_sliceType_Zapper) MarshalLogArray(enc zapcore.ArrayEncoder) (err error) {
	for _, v := range s {
		err = multierr.Append(err, enc.AppendArray((_Set_String_sliceType_Zapper)(v)))
	}
	return err
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Bar.
func (v *Bar) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	err = multierr.Append(err, enc.AddArray("requiredInt32ListField", (_Set_I32_sliceType_Zapper)(v.RequiredInt32ListField)))
	if v.OptionalStringListField != nil {
		err = multierr.Append(err, enc.AddArray("optionalStringListField", (_Set_String_sliceType_Zapper)(v.OptionalStringListField)))
	}
	err = multierr.Append(err, enc.AddArray("requiredTypedefStringListField", (_Set_String_sliceType_Zapper)(([]string)(v.RequiredTypedefStringListField))))
	if v.OptionalTypedefStringListField != nil {
		err = multierr.Append(err, enc.AddArray("optionalTypedefStringListField", (_Set_String_sliceType_Zapper)(([]string)(v.OptionalTypedefStringListField))))
	}
	err = multierr.Append(err, enc.AddArray("requiredFooListField", (_Set_Foo_sliceType_Zapper)(v.RequiredFooListField)))
	if v.OptionalFooListField != nil {
		err = multierr.Append(err, enc.AddArray("optionalFooListField", (_Set_Foo_sliceType_Zapper)(v.OptionalFooListField)))
	}
	err = multierr.Append(err, enc.AddArray("requiredTypedefFooListField", (_Set_Foo_sliceType_Zapper)(([]*Foo)(v.RequiredTypedefFooListField))))
	if v.OptionalTypedefFooListField != nil {
		err = multierr.Append(err, enc.AddArray("optionalTypedefFooListField", (_Set_Foo_sliceType_Zapper)(([]*Foo)(v.OptionalTypedefFooListField))))
	}
	err = multierr.Append(err, enc.AddArray("requiredStringListListField", (_Set_Set_String_sliceType_sliceType_Zapper)(v.RequiredStringListListField)))
	err = multierr.Append(err, enc.AddArray("requiredTypedefStringListListField", (_Set_Set_String_sliceType_sliceType_Zapper)(([][]string)(v.RequiredTypedefStringListListField))))
	return err
}

// GetRequiredInt32ListField returns the value of RequiredInt32ListField if it is set or its
// zero value if it is unset.
func (v *Bar) GetRequiredInt32ListField() (o []int32) {
	if v != nil {
		o = v.RequiredInt32ListField
	}
	return
}

// IsSetRequiredInt32ListField returns true if RequiredInt32ListField is not nil.
func (v *Bar) IsSetRequiredInt32ListField() bool {
	return v != nil && v.RequiredInt32ListField != nil
}

// GetOptionalStringListField returns the value of OptionalStringListField if it is set or its
// zero value if it is unset.
func (v *Bar) GetOptionalStringListField() (o []string) {
	if v != nil && v.OptionalStringListField != nil {
		return v.OptionalStringListField
	}

	return
}

// IsSetOptionalStringListField returns true if OptionalStringListField is not nil.
func (v *Bar) IsSetOptionalStringListField() bool {
	return v != nil && v.OptionalStringListField != nil
}

// GetRequiredTypedefStringListField returns the value of RequiredTypedefStringListField if it is set or its
// zero value if it is unset.
func (v *Bar) GetRequiredTypedefStringListField() (o StringList) {
	if v != nil {
		o = v.RequiredTypedefStringListField
	}
	return
}

// IsSetRequiredTypedefStringListField returns true if RequiredTypedefStringListField is not nil.
func (v *Bar) IsSetRequiredTypedefStringListField() bool {
	return v != nil && v.RequiredTypedefStringListField != nil
}

// GetOptionalTypedefStringListField returns the value of OptionalTypedefStringListField if it is set or its
// zero value if it is unset.
func (v *Bar) GetOptionalTypedefStringListField() (o StringList) {
	if v != nil && v.OptionalTypedefStringListField != nil {
		return v.OptionalTypedefStringListField
	}

	return
}

// IsSetOptionalTypedefStringListField returns true if OptionalTypedefStringListField is not nil.
func (v *Bar) IsSetOptionalTypedefStringListField() bool {
	return v != nil && v.OptionalTypedefStringListField != nil
}

// GetRequiredFooListField returns the value of RequiredFooListField if it is set or its
// zero value if it is unset.
func (v *Bar) GetRequiredFooListField() (o []*Foo) {
	if v != nil {
		o = v.RequiredFooListField
	}
	return
}

// IsSetRequiredFooListField returns true if RequiredFooListField is not nil.
func (v *Bar) IsSetRequiredFooListField() bool {
	return v != nil && v.RequiredFooListField != nil
}

// GetOptionalFooListField returns the value of OptionalFooListField if it is set or its
// zero value if it is unset.
func (v *Bar) GetOptionalFooListField() (o []*Foo) {
	if v != nil && v.OptionalFooListField != nil {
		return v.OptionalFooListField
	}

	return
}

// IsSetOptionalFooListField returns true if OptionalFooListField is not nil.
func (v *Bar) IsSetOptionalFooListField() bool {
	return v != nil && v.OptionalFooListField != nil
}

// GetRequiredTypedefFooListField returns the value of RequiredTypedefFooListField if it is set or its
// zero value if it is unset.
func (v *Bar) GetRequiredTypedefFooListField() (o FooList) {
	if v != nil {
		o = v.RequiredTypedefFooListField
	}
	return
}

// IsSetRequiredTypedefFooListField returns true if RequiredTypedefFooListField is not nil.
func (v *Bar) IsSetRequiredTypedefFooListField() bool {
	return v != nil && v.RequiredTypedefFooListField != nil
}

// GetOptionalTypedefFooListField returns the value of OptionalTypedefFooListField if it is set or its
// zero value if it is unset.
func (v *Bar) GetOptionalTypedefFooListField() (o FooList) {
	if v != nil && v.OptionalTypedefFooListField != nil {
		return v.OptionalTypedefFooListField
	}

	return
}

// IsSetOptionalTypedefFooListField returns true if OptionalTypedefFooListField is not nil.
func (v *Bar) IsSetOptionalTypedefFooListField() bool {
	return v != nil && v.OptionalTypedefFooListField != nil
}

// GetRequiredStringListListField returns the value of RequiredStringListListField if it is set or its
// zero value if it is unset.
func (v *Bar) GetRequiredStringListListField() (o [][]string) {
	if v != nil {
		o = v.RequiredStringListListField
	}
	return
}

// IsSetRequiredStringListListField returns true if RequiredStringListListField is not nil.
func (v *Bar) IsSetRequiredStringListListField() bool {
	return v != nil && v.RequiredStringListListField != nil
}

// GetRequiredTypedefStringListListField returns the value of RequiredTypedefStringListListField if it is set or its
// zero value if it is unset.
func (v *Bar) GetRequiredTypedefStringListListField() (o StringListList) {
	if v != nil {
		o = v.RequiredTypedefStringListListField
	}
	return
}

// IsSetRequiredTypedefStringListListField returns true if RequiredTypedefStringListListField is not nil.
func (v *Bar) IsSetRequiredTypedefStringListListField() bool {
	return v != nil && v.RequiredTypedefStringListListField != nil
}

type Foo struct {
	StringField string `json:"stringField,required"`
}

// ToWire translates a Foo struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Foo) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.StringField), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Foo struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Foo struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Foo
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Foo) FromWire(w wire.Value) error {
	var err error

	stringFieldIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.StringField, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				stringFieldIsSet = true
			}
		}
	}

	if !stringFieldIsSet {
		return errors.New("field StringField of Foo is required")
	}

	return nil
}

// String returns a readable string representation of a Foo
// struct.
func (v *Foo) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("StringField: %v", v.StringField)
	i++

	return fmt.Sprintf("Foo{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Foo match the
// provided Foo.
//
// This function performs a deep comparison.
func (v *Foo) Equals(rhs *Foo) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !(v.StringField == rhs.StringField) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Foo.
func (v *Foo) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	enc.AddString("stringField", v.StringField)
	return err
}

// GetStringField returns the value of StringField if it is set or its
// zero value if it is unset.
func (v *Foo) GetStringField() (o string) {
	if v != nil {
		o = v.StringField
	}
	return
}

type FooList []*Foo

// ToWire translates FooList into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
func (v FooList) ToWire() (wire.Value, error) {
	x := ([]*Foo)(v)
	return wire.NewValueSet(_Set_Foo_sliceType_ValueList(x)), error(nil)
}

// String returns a readable string representation of FooList.
func (v FooList) String() string {
	x := ([]*Foo)(v)
	return fmt.Sprint(x)
}

// FromWire deserializes FooList from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
func (v *FooList) FromWire(w wire.Value) error {
	x, err := _Set_Foo_sliceType_Read(w.GetSet())
	*v = (FooList)(x)
	return err
}

// Equals returns true if this FooList is equal to the provided
// FooList.
func (lhs FooList) Equals(rhs FooList) bool {
	return _Set_Foo_sliceType_Equals(([]*Foo)(lhs), ([]*Foo)(rhs))
}

func (v FooList) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	return ((_Set_Foo_sliceType_Zapper)(([]*Foo)(v))).MarshalLogArray(enc)
}

type MyStringList StringList

// ToWire translates MyStringList into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
func (v MyStringList) ToWire() (wire.Value, error) {
	x := (StringList)(v)
	return x.ToWire()
}

// String returns a readable string representation of MyStringList.
func (v MyStringList) String() string {
	x := (StringList)(v)
	return fmt.Sprint(x)
}

// FromWire deserializes MyStringList from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
func (v *MyStringList) FromWire(w wire.Value) error {
	x, err := _StringList_Read(w)
	*v = (MyStringList)(x)
	return err
}

// Equals returns true if this MyStringList is equal to the provided
// MyStringList.
func (lhs MyStringList) Equals(rhs MyStringList) bool {
	return (StringList)(lhs).Equals((StringList)(rhs))
}

func (v MyStringList) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	return ((_Set_String_sliceType_Zapper)(([]string)(v))).MarshalLogArray(enc)
}

type StringList []string

// ToWire translates StringList into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
func (v StringList) ToWire() (wire.Value, error) {
	x := ([]string)(v)
	return wire.NewValueSet(_Set_String_sliceType_ValueList(x)), error(nil)
}

// String returns a readable string representation of StringList.
func (v StringList) String() string {
	x := ([]string)(v)
	return fmt.Sprint(x)
}

// FromWire deserializes StringList from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
func (v *StringList) FromWire(w wire.Value) error {
	x, err := _Set_String_sliceType_Read(w.GetSet())
	*v = (StringList)(x)
	return err
}

// Equals returns true if this StringList is equal to the provided
// StringList.
func (lhs StringList) Equals(rhs StringList) bool {
	return _Set_String_sliceType_Equals(([]string)(lhs), ([]string)(rhs))
}

func (v StringList) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	return ((_Set_String_sliceType_Zapper)(([]string)(v))).MarshalLogArray(enc)
}

type StringListList [][]string

// ToWire translates StringListList into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
func (v StringListList) ToWire() (wire.Value, error) {
	x := ([][]string)(v)
	return wire.NewValueSet(_Set_Set_String_sliceType_sliceType_ValueList(x)), error(nil)
}

// String returns a readable string representation of StringListList.
func (v StringListList) String() string {
	x := ([][]string)(v)
	return fmt.Sprint(x)
}

// FromWire deserializes StringListList from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
func (v *StringListList) FromWire(w wire.Value) error {
	x, err := _Set_Set_String_sliceType_sliceType_Read(w.GetSet())
	*v = (StringListList)(x)
	return err
}

// Equals returns true if this StringListList is equal to the provided
// StringListList.
func (lhs StringListList) Equals(rhs StringListList) bool {
	return _Set_Set_String_sliceType_sliceType_Equals(([][]string)(lhs), ([][]string)(rhs))
}

func (v StringListList) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	return ((_Set_Set_String_sliceType_sliceType_Zapper)(([][]string)(v))).MarshalLogArray(enc)
}

type _Set_String_mapType_ValueList map[string]struct{}

func (v _Set_String_mapType_ValueList) ForEach(f func(wire.Value) error) error {
	for x := range v {
		w, err := wire.NewValueString(x), error(nil)
		if err != nil {
			return err
		}

		if err := f(w); err != nil {
			return err
		}
	}
	return nil
}

func (v _Set_String_mapType_ValueList) Size() int {
	return len(v)
}

func (_Set_String_mapType_ValueList) ValueType() wire.Type {
	return wire.TBinary
}

func (_Set_String_mapType_ValueList) Close() {}

func _Set_String_mapType_Read(s wire.ValueList) (map[string]struct{}, error) {
	if s.ValueType() != wire.TBinary {
		return nil, nil
	}

	o := make(map[string]struct{}, s.Size())
	err := s.ForEach(func(x wire.Value) error {
		i, err := x.GetString(), error(nil)
		if err != nil {
			return err
		}

		o[i] = struct{}{}
		return nil
	})
	s.Close()
	return o, err
}

func _Set_String_mapType_Equals(lhs, rhs map[string]struct{}) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for x := range rhs {
		if _, ok := lhs[x]; !ok {
			return false
		}
	}

	return true
}

type _Set_String_mapType_Zapper map[string]struct{}

// MarshalLogArray implements zapcore.ArrayMarshaler, enabling
// fast logging of _Set_String_mapType_Zapper.
func (s _Set_String_mapType_Zapper) MarshalLogArray(enc zapcore.ArrayEncoder) (err error) {
	for v := range s {
		enc.AppendString(v)
	}
	return err
}

type StringSet map[string]struct{}

// ToWire translates StringSet into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
func (v StringSet) ToWire() (wire.Value, error) {
	x := (map[string]struct{})(v)
	return wire.NewValueSet(_Set_String_mapType_ValueList(x)), error(nil)
}

// String returns a readable string representation of StringSet.
func (v StringSet) String() string {
	x := (map[string]struct{})(v)
	return fmt.Sprint(x)
}

// FromWire deserializes StringSet from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
func (v *StringSet) FromWire(w wire.Value) error {
	x, err := _Set_String_mapType_Read(w.GetSet())
	*v = (StringSet)(x)
	return err
}

// Equals returns true if this StringSet is equal to the provided
// StringSet.
func (lhs StringSet) Equals(rhs StringSet) bool {
	return _Set_String_mapType_Equals((map[string]struct{})(lhs), (map[string]struct{})(rhs))
}

func (v StringSet) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	return ((_Set_String_mapType_Zapper)((map[string]struct{})(v))).MarshalLogArray(enc)
}

// ThriftModule represents the IDL file used to generate this package.
var ThriftModule = &thriftreflect.ThriftModule{
	Name:     "set_to_slice",
	Package:  "go.uber.org/thriftrw/gen/internal/tests/set_to_slice",
	FilePath: "set_to_slice.thrift",
	SHA1:     "34a394ad873ba45745fa9002fe3625d8162d35e9",
	Raw:      rawIDL,
}

const rawIDL = "typedef set<string> StringSet\ntypedef set<string> (go.type = \"slice\") StringList\ntypedef set<Foo> (go.type = \"slice\") FooList\ntypedef StringList MyStringList\ntypedef MyStringList AnotherStringList\n\ntypedef set<set<string> (go.type = \"slice\")> (go.type = \"slice\") StringListList\n\nstruct Foo {\n    1: required string stringField\n}\n\nstruct Bar {\n    1: required set<i32> (go.type = \"slice\") requiredInt32ListField\n    2: optional set<string> (go.type = \"slice\") optionalStringListField\n    3: required StringList requiredTypedefStringListField\n    4: optional StringList optionalTypedefStringListField\n    5: required set<Foo> (go.type = \"slice\") requiredFooListField\n    6: optional set<Foo> (go.type = \"slice\") optionalFooListField\n    7: required FooList requiredTypedefFooListField\n    8: optional FooList optionalTypedefFooListField\n    9: required set<set<string> (go.type = \"slice\")> (go.type = \"slice\") requiredStringListListField\n    10: required StringListList requiredTypedefStringListListField\n}\n\nconst set<string> (go.type = \"slice\") ConstStringList = [\"hello\"]\nconst set<set<string>(go.type = \"slice\")> (go.type = \"slice\") ConstListStringList = [[\"hello\"], [\"world\"]]\n"
