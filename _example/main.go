package main

import (
	"github.com/mbict/fake"
	. "github.com/mbict/go-factory"

	"fmt"
	"math/rand"
)

type myStruct struct {
	Id   int
	Name string
}

func main() {
	factory := NewFactory()

	factory.Definition((*myStruct)(nil), func(data Data) Data {
		return Data{
			"Id":   rand.Intn(10),
			"Name": fake.FirstName(),
		}
	})

	factory.SetPersistHandler(func(m interface{}) {
		fmt.Printf("Persist called for model: %#v\n", m)
	})

	//
	// create a random seed object
	//

	//no fixed data
	seed := myStruct{}
	fmt.Printf("\nBefore create: %#v\n", seed)
	factory.Create(&seed, nil)
	fmt.Printf("After create: %#v\n", seed)

	//with fixed data
	seed_fixed := myStruct{}
	fmt.Printf("\nBefore create: %#v\n", seed_fixed)
	factory.Create(&seed_fixed, Data{"Name": "bar"})
	fmt.Printf("After create: %#v\n", seed_fixed)

	//
	// now with arrays
	//

	//no fixed data
	seeds := [3]myStruct{}
	fmt.Printf("\nBefore create: %#v\n", seeds)
	factory.Create(&seeds, nil)
	fmt.Printf("After create: %#v\n", seeds)

	//with fixed data
	seeds_fixed := [3]myStruct{}
	fmt.Printf("\nBefore create: %#v\n", seeds_fixed)
	factory.Create(&seeds_fixed, Data{"Name": "bar"})
	fmt.Printf("After create: %#v\n", seeds_fixed)

}
