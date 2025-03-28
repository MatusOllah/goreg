package goreg_test

import (
	"fmt"

	"github.com/MatusOllah/goreg"
)

func ExampleStandardRegistry_Register() {
	type Thing string

	reg := goreg.NewStandardRegistry[Thing]()
	reg.Register("door", Thing("Door"))

	fmt.Println(reg.Get("door"))

	// Output:
	// Door true
}

func ExampleStandardRegistry_Unregister() {
	type Thing string

	reg := goreg.NewStandardRegistry[Thing]()
	reg.Register("door", Thing("Door"))

	fmt.Println(reg.Get("door"))

	reg.Unregister("door")
	fmt.Println(reg.Get("door"))

	// Output:
	// Door true
	//  false
}

func ExampleStandardRegistry_Len() {
	type Thing string

	reg := goreg.NewStandardRegistry[Thing]()
	reg.Register("door", Thing("Door"))
	reg.Register("window", Thing("Window"))
	reg.Register("chair", Thing("Chair"))

	fmt.Println(reg.Len())

	// Output:
	// 3
}

func ExampleStandardRegistry_Iter() {
	type Thing string

	reg := goreg.NewStandardRegistry[Thing]()
	reg.Register("door", Thing("Door"))
	reg.Register("window", Thing("Window"))
	reg.Register("chair", Thing("Chair"))

	for id, obj := range reg.Iter() {
		if id == "door" {
			fmt.Println("found", id, "=", obj)
		}
	}

	// Output:
	// found door = Door
}
