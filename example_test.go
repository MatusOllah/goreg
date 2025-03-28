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

func ExampleStandardRegistry_Reset() {
	type Thing string

	reg := goreg.NewStandardRegistry[Thing]()
	reg.Register("door", Thing("Door"))
	reg.Register("window", Thing("Window"))
	reg.Register("chair", Thing("Chair"))

	fmt.Println(reg.Len())

	reg.Reset()

	fmt.Println(reg.Len())

	// Output:
	// 3
	// 0
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

func ExampleOrderedRegistry_Register() {
	type Thing string

	reg := goreg.NewOrderedRegistry[Thing]()
	reg.Register("door", Thing("Door"))

	fmt.Println(reg.Get("door"))

	// Output:
	// Door true
}

func ExampleOrderedRegistry_Unregister() {
	type Thing string

	reg := goreg.NewOrderedRegistry[Thing]()
	reg.Register("door", Thing("Door"))

	fmt.Println(reg.Get("door"))

	reg.Unregister("door")
	fmt.Println(reg.Get("door"))

	// Output:
	// Door true
	//  false
}

func ExampleOrderedRegistry_GetIndex() {
	type Thing string

	reg := goreg.NewOrderedRegistry[Thing]()
	reg.Register("door", Thing("Door"))
	reg.Register("window", Thing("Window"))
	reg.Register("chair", Thing("Chair"))

	fmt.Println(reg.GetIndex(0))
	fmt.Println(reg.GetIndex(1))
	fmt.Println(reg.GetIndex(2))
	fmt.Println(reg.GetIndex(3))
	fmt.Println(reg.GetIndex(-1))

	// Output:
	// Door true
	// Window true
	// Chair true
	//  false
	//  false
}

func ExampleOrderedRegistry_Len() {
	type Thing string

	reg := goreg.NewOrderedRegistry[Thing]()
	reg.Register("door", Thing("Door"))
	reg.Register("window", Thing("Window"))
	reg.Register("chair", Thing("Chair"))

	fmt.Println(reg.Len())

	// Output:
	// 3
}

func ExampleOrderedRegistry_Reset() {
	type Thing string

	reg := goreg.NewOrderedRegistry[Thing]()
	reg.Register("door", Thing("Door"))
	reg.Register("window", Thing("Window"))
	reg.Register("chair", Thing("Chair"))

	fmt.Println(reg.Len())

	reg.Reset()

	fmt.Println(reg.Len())

	// Output:
	// 3
	// 0
}

func ExampleOrderedRegistry_Iter() {
	type Thing string

	reg := goreg.NewOrderedRegistry[Thing]()
	reg.Register("door", Thing("Door"))
	reg.Register("window", Thing("Window"))
	reg.Register("chair", Thing("Chair"))

	for id, obj := range reg.Iter() {
		fmt.Println(id, "=", obj)
	}

	// Output:
	// door = Door
	// window = Window
	// chair = Chair
}

func ExampleCollect() {
	type Thing string

	reg := goreg.NewOrderedRegistry[Thing]()
	reg.Register("door", Thing("Door"))
	reg.Register("window", Thing("Window"))
	reg.Register("chair", Thing("Chair"))

	theMap := goreg.Collect(reg)

	fmt.Println(theMap)

	// Output:
	// map[chair:Chair door:Door window:Window]
}

func ExampleCopy() {
	type Thing string

	reg1 := goreg.NewOrderedRegistry[Thing]()
	reg1.Register("door", Thing("Door"))
	reg1.Register("window", Thing("Window"))

	reg2 := goreg.NewOrderedRegistry[Thing]()
	reg2.Register("chair", Thing("Chair"))
	reg2.Register("microphone", Thing("Microphone"))

	goreg.Copy(reg1, reg2)

	fmt.Println(reg1)

	// Output:
	// [{door Door} {window Window} {chair Chair} {microphone Microphone}]
}

func ExampleEqual() {
	type Thing string

	reg1 := goreg.NewOrderedRegistry[Thing]()
	reg1.Register("door", Thing("Door"))
	reg1.Register("window", Thing("Window"))

	reg2 := goreg.NewOrderedRegistry[Thing]()
	reg2.Register("door", Thing("Door"))
	reg2.Register("window", Thing("Window"))

	fmt.Println(goreg.Equal(reg1, reg2))

	reg2.Register("chair", Thing("Chair"))
	fmt.Println(goreg.Equal(reg1, reg2))

	// Output:
	// true
	// false
}
