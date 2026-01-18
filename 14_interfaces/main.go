package main

import (
	"fmt"
	"math"
)

// ---------------------------------------------------------
// TOPIC: Interfaces
// ---------------------------------------------------------
// Interfaces in Go are collections of method signatures.
// They specify WHAT an object can do, not HOW it does it.
// If a type implements all methods of an interface, it satisfies that interface implicitly.

// 1. DEFINING AN INTERFACE
// Any type that has an Area() method returning a float64 satisfies the "Shape" interface.
type Shape interface {
	Area() float64
}

// 2. CONCRETE TYPES
type Rectangle struct {
	Width, Height float64
}

// Rectangle implements Shape because it has an Area() method
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

// Circle implements Shape
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// ---------------------------------------------------------
// MAIN EXECUTION
// ---------------------------------------------------------

func main() {
	fmt.Println("--- 1. Polymorphism ---")

	r := Rectangle{Width: 10, Height: 5}
	c := Circle{Radius: 5}

	// We can pass both Rectangle and Circle to this function
	// because they both satisfy the Shape interface.
	printShapeArea(r)
	printShapeArea(c)

	fmt.Println("\n--- 2. Empty Interface (interface{} or any) ---")
	// An empty interface has zero methods.
	// Since every type implements at least zero methods,
	// interface{} can hold a value of ANY type.

	var anything interface{} // Since Go 1.18, you can also use 'any' alias

	anything = "I am a string"
	fmt.Println(anything)

	anything = 42
	fmt.Println(anything)

	anything = true
	fmt.Println(anything)

	fmt.Println("\n--- 3. Type Assertion ---")
	// How to get the concrete value back from an interface?

	var data interface{} = "Hello, Go!"

	// Unsafe assertion (will panic if types don't match)
	str := data.(string)
	fmt.Println("Extracted string:", str)

	// Safe assertion (Comma-ok idiom) -> HIGHLY RECOMMENDED
	// We try to convert 'data' to an int.
	num, ok := data.(int)
	if !ok {
		fmt.Println("Assertion failed: data is not an integer")
	} else {
		fmt.Println("Extracted int:", num)
	}

	fmt.Println("\n--- 4. Type Switch ---")
	checkType(100)
	checkType("Golang")
	checkType(3.14)
	checkType(true)
}

// Helper function for Polymorphism example
// This function accepts the Interface, not a specific struct.
func printShapeArea(s Shape) {
	fmt.Printf("Area of type %T is: %.2f\n", s, s.Area())
}

// Helper function for Type Switch example
func checkType(i interface{}) {
	// 'switch v := i.(type)' is special syntax only valid in switches
	switch v := i.(type) {
	case int:
		fmt.Printf("It's an Integer: %d\n", v)
	case string:
		fmt.Printf("It's a String (len %d): %q\n", len(v), v)
	case float64:
		fmt.Printf("It's a Float: %.2f\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

// ---------------------------------------------------------
// ⚠️ COMMON PITFALLS
// ---------------------------------------------------------
//
// 1. "The Nil Interface" Trap (Very Common!)
//    An interface is a tuple of (type, value).
//    An interface is only equal to 'nil' if BOTH type and value are nil.
//
//    var r *Rectangle = nil   // r is a nil pointer to a Rectangle
//    var s Shape = r          // s holds (type=*Rectangle, value=nil)
//
//    if s == nil { ... }      // This will be FALSE! Even though the value inside is nil.
//
//    Correction: Always check if the concrete pointer is nil before assigning it to an interface,
//    or check using reflection (advanced).
//
// 2. Panic on Assertion
//    Doing `val := i.(int)` without checking `ok` will cause a runtime panic
//    if 'i' is not actually an int. Always use `val, ok := ...` unless you are 100% sure.
