// Code generated by thriftrw

package typedefs

import (
	"fmt"
	"github.com/thriftrw/thriftrw-go/wire"
	"strings"
)

type Event struct {
	UUID *UUID
	Time *Timestamp
}

func (v *Event) ToWire() wire.Value {
	var fields [2]wire.Field
	i := 0
	fields[i] = wire.Field{ID: 1, Value: v.UUID.ToWire()}
	i++
	if v.Time != nil {
		fields[i] = wire.Field{ID: 2, Value: v.Time.ToWire()}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]})
}
func _UUID_Read(w wire.Value) (*UUID, error) {
	var x UUID
	err := x.FromWire(w)
	return &x, err
}
func _Timestamp_Read(w wire.Value) (Timestamp, error) {
	var x Timestamp
	err := x.FromWire(w)
	return x, err
}
func (v *Event) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.UUID, err = _UUID_Read(field.Value)
				if err != nil {
					return err
				}
			}
		case 2:
			if field.Value.Type() == wire.TI64 {
				var x Timestamp
				x, err = _Timestamp_Read(field.Value)
				v.Time = &x
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
func (v *Event) String() string {
	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("UUID: %v", v.UUID)
	i++
	if v.Time != nil {
		fields[i] = fmt.Sprintf("Time: %v", *(v.Time))
		i++
	}
	return fmt.Sprintf("Event{%v}", strings.Join(fields[:i], ", "))
}

type _List_Event_ValueList []*Event

func (v _List_Event_ValueList) ForEach(f func(wire.Value) error) error {
	for _, x := range v {
		err := f(x.ToWire())
		if err != nil {
			return err
		}
	}
	return nil
}
func (v _List_Event_ValueList) Close() {
}
func _Event_Read(w wire.Value) (*Event, error) {
	var v Event
	err := v.FromWire(w)
	return &v, err
}
func _List_Event_Read(l wire.List) ([]*Event, error) {
	if l.ValueType != wire.TStruct {
		return nil, nil
	}
	o := make([]*Event, 0, l.Size)
	err := l.Items.ForEach(func(x wire.Value) error {
		i, err := _Event_Read(x)
		if err != nil {
			return err
		}
		o = append(o, i)
		return nil
	})
	l.Items.Close()
	return o, err
}

type EventGroup []*Event

func (v EventGroup) ToWire() wire.Value {
	x := ([]*Event)(v)
	return wire.NewValueList(wire.List{ValueType: wire.TStruct, Size: len(x), Items: _List_Event_ValueList(x)})
}
func (v *EventGroup) FromWire(w wire.Value) error {
	x, err := _List_Event_Read(w.GetList())
	*v = (EventGroup)(x)
	return err
}

type Pdf []byte

func (v Pdf) ToWire() wire.Value {
	x := ([]byte)(v)
	return wire.NewValueBinary(x)
}
func (v *Pdf) FromWire(w wire.Value) error {
	x, err := w.GetBinary(), error(nil)
	*v = (Pdf)(x)
	return err
}

type State string

func (v State) ToWire() wire.Value {
	x := (string)(v)
	return wire.NewValueString(x)
}
func (v *State) FromWire(w wire.Value) error {
	x, err := w.GetString(), error(nil)
	*v = (State)(x)
	return err
}

type Timestamp int64

func (v Timestamp) ToWire() wire.Value {
	x := (int64)(v)
	return wire.NewValueI64(x)
}
func (v *Timestamp) FromWire(w wire.Value) error {
	x, err := w.GetI64(), error(nil)
	*v = (Timestamp)(x)
	return err
}

type Transition struct {
	From   State
	To     State
	Events EventGroup
}

func (v *Transition) ToWire() wire.Value {
	var fields [3]wire.Field
	i := 0
	fields[i] = wire.Field{ID: 1, Value: v.From.ToWire()}
	i++
	fields[i] = wire.Field{ID: 2, Value: v.To.ToWire()}
	i++
	if v.Events != nil {
		fields[i] = wire.Field{ID: 3, Value: v.Events.ToWire()}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]})
}
func _State_Read(w wire.Value) (State, error) {
	var x State
	err := x.FromWire(w)
	return x, err
}
func _EventGroup_Read(w wire.Value) (EventGroup, error) {
	var x EventGroup
	err := x.FromWire(w)
	return x, err
}
func (v *Transition) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.From, err = _State_Read(field.Value)
				if err != nil {
					return err
				}
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				v.To, err = _State_Read(field.Value)
				if err != nil {
					return err
				}
			}
		case 3:
			if field.Value.Type() == wire.TList {
				v.Events, err = _EventGroup_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
func (v *Transition) String() string {
	var fields [3]string
	i := 0
	fields[i] = fmt.Sprintf("From: %v", v.From)
	i++
	fields[i] = fmt.Sprintf("To: %v", v.To)
	i++
	if v.Events != nil {
		fields[i] = fmt.Sprintf("Events: %v", v.Events)
		i++
	}
	return fmt.Sprintf("Transition{%v}", strings.Join(fields[:i], ", "))
}

type UUID I128

func (v *UUID) ToWire() wire.Value {
	x := (*I128)(v)
	return x.ToWire()
}
func (v *UUID) FromWire(w wire.Value) error {
	return (*I128)(v).FromWire(w)
}

type I128 struct {
	High int64
	Low  int64
}

func (v *I128) ToWire() wire.Value {
	var fields [2]wire.Field
	i := 0
	fields[i] = wire.Field{ID: 1, Value: wire.NewValueI64(v.High)}
	i++
	fields[i] = wire.Field{ID: 2, Value: wire.NewValueI64(v.Low)}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]})
}
func (v *I128) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TI64 {
				v.High, err = field.Value.GetI64(), error(nil)
				if err != nil {
					return err
				}
			}
		case 2:
			if field.Value.Type() == wire.TI64 {
				v.Low, err = field.Value.GetI64(), error(nil)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
func (v *I128) String() string {
	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("High: %v", v.High)
	i++
	fields[i] = fmt.Sprintf("Low: %v", v.Low)
	i++
	return fmt.Sprintf("I128{%v}", strings.Join(fields[:i], ", "))
}
