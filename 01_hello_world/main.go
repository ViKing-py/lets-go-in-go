package main

// ---------------------------------------------------------
// TOPIC: Hello World & Basic Structure
// ---------------------------------------------------------

// 1. PACKAGE DECLARATION
// Every Go file must start with a package name.
// "package main" tells the Go compiler that this package should compile
// as an executable program rather than a shared library.
// Only the "main" package can contain the "main" function.

import "fmt"

// 2. IMPORTS
// We import the "fmt" package (short for format).
// It contains functions for formatting text, including printing to the console.

// 3. THE MAIN FUNCTION
// This is the entry point of the application.
// When you run the program, the code inside this function executes first.
func main() {
	// Calling a function from the "fmt" package.
	// "Println" prints the text and moves to a new line.
	fmt.Println("Hello, Go Developer!")

	fmt.Println("Welcome to the first lesson.")
}

// ---------------------------------------------------------
// ⚠️ COMMON PITFALLS (Watch out!)
// ---------------------------------------------------------
//
// 1. Missing '{' placement:
//    In Go, the opening brace '{' MUST be on the same line as the function declaration.
//    
//    WRONG:
//    func main()
//    {
//    }
//
//    CORRECT:
//    func main() {
//    }
//
// 2. Unused Imports:
//    If you import "fmt" but don't use it, Go will throw a compile-time error.
//    Go forces you to keep your code clean!