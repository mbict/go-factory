package factory

import (
	"errors"
	"fmt"
	"reflect"
)

// Data represents the map structure for the data fixtures
type Data map[string]interface{}

// Definition is a function who wraps the data creation logic to fill a structure.
type Definition func(data Data) Data

type definitionList map[reflect.Type]Definition

// PersistHandler is a function who is called after a new model is created and needs to be stored somewhere.
type PersistHandler func(interface{})

// Factory is a model factory to create database models on the fly for testing.
type Factory interface {
	Create(model interface{}, data ...Data) error
	SetPersistHandler(handler PersistHandler)
	Definition(model interface{}, definition Definition) Factory
}

type factory struct {
	definitions    definitionList
	persistHandler PersistHandler
}

// NewFactory will create and initialize a Factory
func NewFactory() Factory {
	return &factory{
		definitions: make(definitionList),
	}
}

// SetPersistHandler will set the handler who is called for every
// model created. You can wire your database handler for persistence here
func (f *factory) SetPersistHandler(handler PersistHandler) {
	f.persistHandler = handler
}

// Create will create a new model based on the provided model
// You can provide the following types (by reference)
// - a single element type (yourStruct)
// - a pointer type (*yourStruct)
// - a nil pointer type (*yourStruct)(nil)
// - array/slices type ([]yourStruct)
// - array/slices pointer type ([]*yourStruct)
// when you provide a array/slice make sure it has elements (nil or initialized)
// create will fill all the elements
func (f factory) Create(model interface{}, data ...Data) error {
	//find the type
	t := getType(model)
	def, ok := f.definitions[t]
	if !ok {
		return fmt.Errorf("definition for type `%s` not found", t)
	}

	v := reflect.ValueOf(model)
	if v.Kind() != reflect.Ptr || !v.Elem().CanSet() {
		return errors.New("provided model is not by reference")
	}

	//reset element to zero variant
	v = v.Elem()
	if v.Kind() == reflect.Ptr && v.CanSet() {
		v.Set(reflect.New(v.Type().Elem()))
	}
	v = reflect.Indirect(v)

	if v.Kind() == reflect.Array || v.Kind() == reflect.Slice { //slice or array types
		for i := 0; i < v.Len(); i++ {
			vs := v.Index(i)
			if vs.Kind() != reflect.Ptr {
				vs = vs.Addr()
			}

			if vs.CanSet() {
				vs.Set(reflect.New(vs.Type().Elem()))
			}

			if err := populate(vs.Interface(), generate(def, data...)); err != nil {
				return err
			}

			if f.persistHandler != nil {
				f.persistHandler(vs.Interface())
			}
		}
	} else { //singular types
		if err := populate(model, generate(def, data...)); err != nil {
			return err
		}

		if f.persistHandler != nil {
			f.persistHandler(model)
		}
	}
	return nil
}

// Definition will add a new model to the definitions table
// The provided definition should be a function who accepts the data as the first param
// and should return the model
func (f *factory) Definition(model interface{}, definition Definition) Factory {
	t := getType(model)
	f.definitions[t] = definition
	return f
}

// Generate generates the data based on the definition function and overwrites data
// with the fixed data if provided
func generate(def Definition, data ...Data) Data {
	overrideValues := Data{}
	for _, values := range data {
		for key, value := range values {
			overrideValues[key] = value
		}
	}

	fillData := def(overrideValues)

	for key, value := range overrideValues {
		fillData[key] = value
	}

	return fillData
}

// getType will get the absolute type of a provided obj
func getType(obj interface{}) reflect.Type {
	t := reflect.TypeOf(obj)
	for t.Kind() == reflect.Ptr || t.Kind() == reflect.Slice || t.Kind() == reflect.Array {
		t = t.Elem()
	}
	return t
}

// populate will fill the provided struct with the provided data
func populate(obj interface{}, data Data) error {
	for name, value := range data {
		err := setField(obj, name, value)
		if err != nil {
			return err
		}
	}
	return nil
}

// setField will set a property of the struct
func setField(obj interface{}, name string, value interface{}) error {
	v := reflect.Indirect(reflect.ValueOf(obj).Elem())
	field := v.FieldByName(name)

	if !field.IsValid() {
		return fmt.Errorf("no such field: `%s` in obj", name)
	}

	if !field.CanSet() {
		return fmt.Errorf("cannot set `%s` field value", name)
	}

	structFieldType := field.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return fmt.Errorf("provided value type (%s) didn't match obj field type (%s) for field `%s`", val.Type(), structFieldType, name)
	}

	field.Set(val)
	return nil
}
