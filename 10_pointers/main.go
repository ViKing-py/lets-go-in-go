package main

import "fmt"

// ---------------------------------------------------------
// TOPIC: Pointers
// ---------------------------------------------------------

func main() {
	// 1. BASICS: Address (&) and Type (*)
	fmt.Println("--- 1. Basics ---")

	var age int = 25
	fmt.Println("Original Value:", age)

	// The '&' operator generates a pointer to its operand.
	// 'ptr' holds the memory address where 'age' is stored.
	// The type of 'ptr' is *int (pointer to an integer).
	var ptr *int = &age

	fmt.Println("Address (ptr):", ptr)

	// 2. DEREFERENCING (*)
	// The '*' operator denotes the pointer's underlying value.
	// This is often called "dereferencing".
	fmt.Println("Value via pointer (*ptr):", *ptr)

	// We can change the value at that address through the pointer.
	*ptr = 30
	fmt.Println("New Value (age) after changing *ptr:", age) // age is now 30

	// 3. PASS BY VALUE VS PASS BY POINTER
	fmt.Println("\n--- 2. Function Arguments ---")

	number := 100

	// Case A: Pass by Value (Copy)
	modifyValue(number)
	fmt.Println("After modifyValue:", number) // Remains 100

	// Case B: Pass by Pointer (Reference)
	modifyPointer(&number)
	fmt.Println("After modifyPointer:", number) // Becomes 999

	// 4. NIL POINTERS
	fmt.Println("\n--- 3. Nil Pointers ---")

	// The zero value of a pointer is nil.
	var emptyPtr *int
	fmt.Println("Value of emptyPtr:", emptyPtr)

	// Safe check before usage:
	if emptyPtr == nil {
		fmt.Println("Pointer is nil, skipping dereference.")
	}
}

// modifyValue receives a COPY of the integer.
// Changing 'n' here does not affect the original variable.
func modifyValue(n int) {
	n = 0
}

// modifyPointer receives the MEMORY ADDRESS of the integer.
// Changing '*n' here changes the original variable.
func modifyPointer(n *int) {
	*n = 999
}

// ---------------------------------------------------------
// ⚠️ COMMON PITFALLS
// ---------------------------------------------------------
//
// 1. Dereferencing a Nil Pointer (The "Panic"):
//    If you try to read or write to a nil pointer, the program crashes.
//
//    var p *int // p is nil
//    *p = 10    // CRASH! runtime error: invalid memory address or nil pointer dereference
//
//    ALWAYS check if a pointer is nil if you are unsure if it has been initialized.
//
// 2. Confusing the Asterisk (*):
//    - In a type declaration (var p *int), '*' means "this is a pointer type".
//    - In code logic (*p = 10), '*' means "read/write the value at this address".
//
// 3. No Pointer Arithmetic:
//    Unlike C or C++, Go does not allow you to do things like 'ptr++' to move
//    to the next memory slot. Go prioritizes safety over this flexibility.
