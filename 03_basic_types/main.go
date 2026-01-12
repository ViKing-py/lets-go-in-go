package main

import "fmt"

// ---------------------------------------------------------
// TOPIC: Basic Types, Zero Values & Type Casting
// ---------------------------------------------------------

func main() {
	fmt.Println("=== 1. INTEGER TYPES ===")
	// 'int' is the most common type. Its size (32 or 64 bits) depends on your system.
	// explicit declaration:
	var age int = 25
	// short declaration (type inferred):
	items := 10

	fmt.Printf("Age: %d, Type: %T\n", age, age)
	fmt.Printf("Items: %d, Type: %T\n", items, items)

	// There are specific sizes: int8, int16, int32, int64
	// And unsigned types (positive only): uint8, uint16...
	var veryBigNumber int64 = 9223372036854775807
	fmt.Println("Big Int:", veryBigNumber)

	fmt.Println("\n=== 2. FLOAT TYPES ===")
	// Go has float32 and float64.
	// default inference is always float64 (more precision).
	price := 19.99
	fmt.Printf("Price: %f, Type: %T\n", price, price)

	fmt.Println("\n=== 3. BOOLEAN & STRING ===")
	// Bool: true or false
	isActive := true
	fmt.Println("Is Active?", isActive)

	// String: Double quotes "" are used for strings.
	// Strings in Go are immutable (you cannot change one character inside it).
	name := "Golang"
	fmt.Println("Name:", name)

	fmt.Println("\n=== 4. ZERO VALUES ===")
	// Crucial Concept: In Go, variables declared without an initial value
	// are NOT "undefined" or "null". They get a "Zero Value".
	var defaultInt int       // 0
	var defaultFloat float64 // 0.0
	var defaultBool bool     // false
	var defaultString string // "" (empty string)

	fmt.Printf("Zero Int: %d\n", defaultInt)
	fmt.Printf("Zero Float: %f\n", defaultFloat)
	fmt.Printf("Zero Bool: %v\n", defaultBool)
	fmt.Printf("Zero String: '%s'\n", defaultString)

	fmt.Println("\n=== 5. TYPE CASTING (CONVERSION) ===")
	// Go is a STATICALLY typed language with STRONG typing.
	// You cannot imply types. You must explicitly convert them using T(v).

	var a int = 10
	var b float64 = 5.5

	// WRONG: total := a + b (Compiler Error: mismatched types int and float64)

	// CORRECT: Convert int to float first
	total := float64(a) + b
	fmt.Println("Total:", total)

	// Converting float to int (Truncates the decimal part!)
	var c float64 = 9.99
	var d int = int(c)
	fmt.Println("9.99 converted to int is:", d) // Result is 9, not 10
}

// ---------------------------------------------------------
// ⚠️ COMMON PITFALLS
// ---------------------------------------------------------
//
// 1. Mismatched Types in Math:
//    You cannot add 'int' to 'int64' or 'int' to 'float64' without casting.
//    Even int and int64 are treated as completely different types.
//
// 2. Integer Division:
//    When dividing two integers, the result is an integer.
//    Example:
//    res := 3 / 2  -> Result is 1, not 1.5
//
//    Fix:
//    res := 3.0 / 2.0 -> Result is 1.5
//
// 3. Unused Variables:
//    If you declare 'var x int' and don't use 'x', code won't compile.
