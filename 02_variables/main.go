package main

import "fmt"

// ---------------------------------------------------------
// TOPIC: Variables, Constants, and Shadowing
// ---------------------------------------------------------

// 1. PACKAGE LEVEL VARIABLES (Global to the package)
// Can be declared using 'var', but NOT with short declaration ':='.
var packageVar string = "I am available throughout the package"

func main() {
	fmt.Println("--- 1. Variable Declarations ---")

	// A) Standard Declaration (var name type)
	// Useful when you don't have an initial value yet.
	// Go assigns a "Zero Value" automatically (0 for int, "" for string, false for bool).
	var age int
	fmt.Println("Zero value of age:", age)
	age = 25
	fmt.Println("Assigned age:", age)

	// B) Type Inference (var name = value)
	// Go guesses the type based on the value (string in this case).
	var name = "Gopher"
	fmt.Println("Name:", name)

	// C) Short Variable Declaration (name := value)
	// The most common way in Go. Only works INSIDE functions.
	// It declares AND initializes.
	city := "Kyiv"
	fmt.Println("City:", city)

	// D) Multiple Declaration
	var (
		width  int = 100
		height int = 200
	)
	fmt.Println("Dimensions:", width, "x", height)

	// ---------------------------------------------------------

	fmt.Println("\n--- 2. Constants ---")

	// Constants are immutable. They cannot be changed after definition.
	// Calculated at compile time.
	const Pi = 3.14159
	const AppName = "MyGoApp"

	// AppName = "NewName" // COMPILER ERROR: cannot assign to AppName
	fmt.Println("Pi:", Pi)

	// ---------------------------------------------------------

	fmt.Println("\n--- 3. Variable Scope & Shadowing (Important!) ---")

	// SCOPE: Where a variable is visible.
	// SHADOWING: Declaring a variable with the same name in an inner scope.

	x := 10                                 // Outer 'x'
	fmt.Println("Outer x before block:", x) // Prints 10

	// Creating a new block (inner scope)
	{
		// ⚠️ SHADOWING HAPPENS HERE
		// We use ':=' which creates a NEW variable named 'x' specific to this block.
		// It "shadows" (hides) the outer 'x'.
		x := 50
		fmt.Println("Inner x inside block:", x) // Prints 50

		// If we used '=' instead of ':=', we would overwrite the outer variable.
		// x = 50
	}

	// Back in the outer scope
	fmt.Println("Outer x after block:", x) // Prints 10 (unchanged because of shadowing)
}

// ---------------------------------------------------------
// ⚠️ COMMON PITFALLS
// ---------------------------------------------------------
//
// 1. Unused Variables:
//    Go considers unused variables a specific error, not just a warning.
//    If you declare 'score := 10' and never read it, the code won't compile.
//
// 2. Short Declaration Re-assignment:
//    You cannot use ':=' twice on the exact same variable in the same scope.
//
//    x := 1
//    x := 2 // ERROR: no new variables on left side of :=
//    x = 2  // CORRECT: simple assignment
//
// 3. Accidental Shadowing:
//    Be careful with ':=' inside 'if' or 'for' blocks. You might think you are
//    updating an outer variable, but you are actually creating a temporary local one.
