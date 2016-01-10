[![Build Status](https://travis-ci.org/mbict/go-factory.png?branch=master)](https://travis-ci.org/mbict/go-factory)
[![GoDoc](https://godoc.org/github.com/mbict/go-factory?status.png)](http://godoc.org/github.com/mbict/go-factory)
[![GoCover](http://gocover.io/_badge/github.com/mbict/go-factory)](http://gocover.io/github.com/mbict/go-factory)

Factory
=======

The goal of this package is to enable the rapid creation of objects for the purpose of testing.

Why i created this
=====
In my tests i do a lot of acceptance testing and want to fill the database with random seeding data.
The creation of the seeding data does not need to clutter up all my tests but should be a simple line.
In case i want to use fixed values i need a simple dynamic way to override the generated values.

Install
=======
```
go get github.com/mbict/go-factory
```

Import
======
```GO
import (
    . "github.com/mbict/go-factory"
)
```

Usage
=====


### Introduction

In the following examples we assume you have a structure definition defined for myStruct who looks like:
```GO
type myStruct struct {
    Id int
    Name string
}
```

### Creation

Creation of a model is pretty simple just provide the the type you want.
The first param is a reference to the value type or array/slice.
The second param is a map with Data that overrides the random data the the definition will create.

#### example: Value type
```GO
seed := &myStruct{}
err := f.Create(&seed)
```

#### example: Value type with a fixed id of 1
The seed will have a fixed Id of 123, this is handy if the model definition generates random data and we want to override that behaviour
```GO
seed := &myStruct{}
err := f.Create(&seed, Data{"Id": 123})
```

#### example: Pointer to a initialized Value
```GO
seed := &myStruct{}
err := f.Create(&seed)
```

#### example: Nil pointer
Create will allocate and initialize the pointer and seed with data
```GO
seed := (*myStruct)(nil)
err := f.Create(&seed)
```

#### example: fixed array
Create will seed all the 3 structures in the array with data
```GO
var seeds [3]myStruct
err := f.Create(&seed)
```

#### example: fixed array with nil pointers
Create will allocate all 3 structures in the array, and seed them with data
```GO
var seeds [3]*myStruct
err := f.Create(&seed)
```

#### example: Slice
Slices ... no problem
```GO
var seeds = []*myStruct{
		{Id: 1},
		{Id: 2},
		{Id: 3},
	}
err := f.Create(&seed)
```

### Model Definitions
All object who could be created need to be defined first.
The definition is a some kind of a generator that should generate unique data for populating the object.

#### example
```GO
f := NewFactory()
f.Definition((*myStruct)(nil), func(data Data) Data {
    return Data{
        "Id":   rand.Intn(10), // <- using the random int from math
        "Name": fake.ProductName(), // <- or use a fake library (github.com/mbict/fake)
    }
})
```

### Persist Callbacks
When you need to store the created objects in for example a database you can register a persist handler.
```GO
f := NewFactory()
f.SetPersistHandler(func(m interface{}) {
    db.Save(m) // <- my implementation of a ORM (github.com/mbict/storm)
})
```

Author
======
Michael Boke

