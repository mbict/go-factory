package factory

import (
	"testing"
)

type (
	testStruct struct {
		ID   int
		Name string
	}
)

func createFactory() Factory {
	f := NewFactory()
	f.Definition((*testStruct)(nil), func(data Data) Data {
		return Data{
			"ID":   1,
			"Name": "foo",
		}
	})
	return f
}

//
// Generation tests
//

func TestCreateSingleDefinitionWithInitializedStruct(t *testing.T) {
	factory := createFactory()
	seed := testStruct{}

	err := factory.Create(&seed)

	if err != nil {
		t.Fatalf("Should not generate a error but got a error `%s`", err.Error())
	}

	if seed.ID != 1 {
		t.Errorf("Expected ID value `%d` but got `%d`", 1, seed.ID)
	}

	if seed.Name != "foo" {
		t.Errorf("Expected Name value `%s` but got `%s`", "foo", seed.Name)
	}
}

func TestCreateSingleDefinitionWithFixedValues(t *testing.T) {
	factory := createFactory()
	seed := testStruct{}

	err := factory.Create(&seed, Data{"ID": 2, "Name": "bar"})

	if err != nil {
		t.Fatalf("Should not generate a error but got a error `%s`", err.Error())
	}

	if seed.ID != 2 {
		t.Errorf("Expected ID value `%d` but got `%d`", 2, seed.ID)
	}

	if seed.Name != "bar" {
		t.Errorf("Expected Name value `%s` but got `%s`", "bar", seed.Name)
	}
}

func TestCreateSingleDefinitionWithNilPointer(t *testing.T) {
	factory := createFactory()
	seed := (*testStruct)(nil)

	err := factory.Create(&seed)

	if err != nil {
		t.Fatalf("Should not generate a error but got a error `%s`", err.Error())
	}

	if seed == nil {
		t.Error("Created model is nil instead of a initialized object")
	}

	if seed.ID != 1 {
		t.Errorf("Expected ID value `%d` but got `%d`", 1, seed.ID)
	}

	if seed.Name != "foo" {
		t.Errorf("Expected Name value `%s` but got `%s`", "foo", seed.Name)
	}
}

func TestCreateSingleDefinitionWithInitializedPointer(t *testing.T) {
	factory := createFactory()
	seed := &testStruct{}

	err := factory.Create(&seed)

	if err != nil {
		t.Fatalf("Should not generate a error but got a error `%s`", err.Error())
	}

	if seed == nil {
		t.Error("Created model is nil instead of a initialized object")
	}

	if seed.ID != 1 {
		t.Errorf("Expected ID value `%d` but got `%d`", 1, seed.ID)
	}

	if seed.Name != "foo" {
		t.Errorf("Expected Name value `%s` but got `%s`", "foo", seed.Name)
	}
}

func TestCreateMultipleDefinitionsWithInitializedArray(t *testing.T) {
	factory := createFactory()
	var seeds [3]testStruct

	err := factory.Create(&seeds)

	if err != nil {
		t.Fatalf("Should not generate a error but got a error `%s`", err.Error())
	}

	if len(seeds) != 3 {
		t.Errorf("Expected to have a slice length of `%d` but got a length `%d`", 3, len(seeds))
	}

	for _, seed := range seeds {
		if seed.ID != 1 {
			t.Errorf("Expected ID value `%d` but got `%d`", 1, seed.ID)
		}

		if seed.Name != "foo" {
			t.Errorf("Expected Name value `%s` but got `%s`", "foo", seed.Name)
		}
	}
}

func TestCreateMultipleDefinitionsWithInitializedSlice(t *testing.T) {
	factory := createFactory()
	var seeds = []testStruct{
		{},
		{},
		{},
	}

	err := factory.Create(&seeds)

	if err != nil {
		t.Fatalf("Should not generate a error but got a error `%s`", err.Error())
	}

	if len(seeds) != 3 {
		t.Errorf("Expected to have a slice length of `%d` but got a length `%d`", 3, len(seeds))
	}

	for _, seed := range seeds {
		if seed.ID != 1 {
			t.Errorf("Expected ID value `%d` but got `%d`", 1, seed.ID)
		}

		if seed.Name != "foo" {
			t.Errorf("Expected Name value `%s` but got `%s`", "foo", seed.Name)
		}
	}
}

func TestCreateMultipleDefinitionsWithInitializedArrayAndFixedValues(t *testing.T) {
	factory := createFactory()
	var seeds [3]testStruct

	err := factory.Create(&seeds, Data{"ID": 2, "Name": "bar"})

	if err != nil {
		t.Fatalf("Should not generate a error but got a error `%s`", err.Error())
	}

	if len(seeds) != 3 {
		t.Errorf("Expected to have a slice length of `%d` but got a length `%d`", 3, len(seeds))
	}

	for _, seed := range seeds {
		if seed.ID != 2 {
			t.Errorf("Expected ID value `%d` but got `%d`", 2, seed.ID)
		}

		if seed.Name != "bar" {
			t.Errorf("Expected Name value `%s` but got `%s`", "bar", seed.Name)
		}
	}
}

func TestCreateMultipleDefinitionsWithUnintializedPtrArray(t *testing.T) {
	factory := createFactory()
	var seeds [3]*testStruct

	err := factory.Create(&seeds)

	if err != nil {
		t.Fatalf("Should not generate a error but got a error `%s`", err.Error())
	}

	if len(seeds) != 3 {
		t.Errorf("Expected to have a slice length of `%d` but got a length `%d`", 3, len(seeds))
	}

	for _, seed := range seeds {
		if seed.ID != 1 {
			t.Errorf("Expected ID value `%d` but got `%d`", 1, seed.ID)
		}

		if seed.Name != "foo" {
			t.Errorf("Expected Name value `%s` but got `%s`", "foo", seed.Name)
		}
	}
}

func TestCreateMultipleDefinitionsWithNilPtrSlice(t *testing.T) {
	factory := createFactory()
	var seeds = []*testStruct{
		nil,
		nil,
		nil,
	}

	err := factory.Create(&seeds)

	if err != nil {
		t.Fatalf("Should not generate a error but got a error `%s`", err.Error())
	}

	if len(seeds) != 3 {
		t.Errorf("Expected to have a slice length of `%d` but got a length `%d`", 3, len(seeds))
	}

	for _, seed := range seeds {
		if seed.ID != 1 {
			t.Errorf("Expected ID value `%d` but got `%d`", 1, seed.ID)
		}

		if seed.Name != "foo" {
			t.Errorf("Expected Name value `%s` but got `%s`", "foo", seed.Name)
		}
	}
}

func TestCreateMultipleDefinitionsWithIntializedPtrSlice(t *testing.T) {
	factory := createFactory()
	var seeds = []*testStruct{
		{ID: 5},
		{ID: 6},
		{ID: 7},
	}

	err := factory.Create(&seeds)

	if err != nil {
		t.Fatalf("Should not generate a error but got a error `%s`", err.Error())
	}

	if len(seeds) != 3 {
		t.Errorf("Expected to have a slice length of `%d` but got a length `%d`", 3, len(seeds))
	}

	for _, seed := range seeds {
		if seed.ID != 1 {
			t.Errorf("Expected ID value `%d` but got `%d`", 1, seed.ID)
		}

		if seed.Name != "foo" {
			t.Errorf("Expected Name value `%s` but got `%s`", "foo", seed.Name)
		}
	}
}

//
// Persist handler tests
//

func TestCallToPersistOnce(t *testing.T) {
	factory := createFactory()
	seed := testStruct{}
	calledTimes := 0
	factory.SetPersistHandler(func(m interface{}) {
		calledTimes++
		m, ok := m.(*testStruct)

		if !ok {
			t.Errorf("persist model cannot be converted to type `testStruct`")
		}

		if m != &seed {
			t.Errorf("persits model is not the same as the provided model to create")
		}
	})

	err := factory.Create(&seed)

	if err != nil {
		t.Fatalf("Should not generate a error but got a error `%s`", err.Error())
	}

	if calledTimes != 1 {
		t.Errorf("Expected persist function to called once, but is called %d times ", calledTimes)
	}
}

func TestCallToPersistForSlices(t *testing.T) {
	factory := createFactory()
	var seeds [3]*testStruct
	calledTimes := 0
	factory.SetPersistHandler(func(m interface{}) {
		m, ok := m.(*testStruct)

		if !ok {
			t.Errorf("persist model cannot be converted to type `testStruct`")
		}

		if m != seeds[calledTimes] {
			t.Errorf("persits model is not the same as the provided model to create")
		}

		calledTimes++
	})

	err := factory.Create(&seeds)

	if err != nil {
		t.Fatalf("Should not generate a error but got a error `%s`", err.Error())
	}

	if calledTimes != 3 {
		t.Errorf("Expected persist function to called 3 times, but is called %d times ", calledTimes)
	}
}

//
// Error tests
//

func TestThrowsErrorWhenCreateUnregisteredDefinitionType(t *testing.T) {
	factory := NewFactory()
	seed := (*testStruct)(nil)

	err := factory.Create(&seed)

	if err == nil {
		t.Fatalf("Should generate a error but got nothing")
	}

	if err.Error() != "definition for type `factory.testStruct` not found" {
		t.Fatalf("Expected error `%s` but got `%s`", "definition for type `factory.testStruct` not found", err.Error())
	}
}

func TestThrowsErrorWhenCreateModelNotByReference(t *testing.T) {
	factory := createFactory()
	seed := (*testStruct)(nil)

	err := factory.Create(seed)

	if err == nil {
		t.Fatalf("Should generate a error but got nothing")
	}

	if err.Error() != "provided model is not by reference" {
		t.Fatalf("Expected error `%s` but got `%s`", "provided model is not by reference", err.Error())
	}
}

func TestThrowsErrorWhenCreateModelWithNonExistingFields(t *testing.T) {
	factory := createFactory()
	seed := (*testStruct)(nil)

	err := factory.Create(&seed, Data{"Title": "foo"})

	if err == nil {
		t.Fatalf("Should generate a error but got nothing")
	}

	if err.Error() != "no such field: `Title` in obj" {
		t.Fatalf("Expected error `%s` but got `%s`", "no such field: `Title` in obj", err.Error())
	}
}

func TestThrowsErrorWhenCreateModelWithIncompatibleType(t *testing.T) {
	factory := createFactory()
	seed := (*testStruct)(nil)

	err := factory.Create(&seed, Data{"Name": int64(1)})

	if err == nil {
		t.Fatalf("Should generate a error but got nothing")
	}

	if err.Error() != "provided value type (int64) didn't match obj field type (string) for field `Name`" {
		t.Fatalf("Expected error `%s` but got `%s`", "provided value type (int64) didn't match obj field type (string) for field `Name`", err.Error())
	}
}
