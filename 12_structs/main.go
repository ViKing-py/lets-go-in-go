package main

import "fmt"

// ---------------------------------------------------------
// TOPIC: Defer, Panic, and Recover
// ---------------------------------------------------------

func main() {
	fmt.Println("--- PART 1: Basic Defer ---")
	basicDefer()

	fmt.Println("\n--- PART 2: Defer Stack (LIFO) ---")
	stackedDefers()

	fmt.Println("\n--- PART 3: Panic and Recover ---")
	// We wrap the risky code in a function to show that 'main' continues
	// even after the sub-function crashes.
	executeRiskyOperation()

	fmt.Println("\n✅ Main function reached the end gracefully.")
}

// 1. DEFER
// The 'defer' keyword postpones the execution of a function until the
// surrounding function returns. It is often used for cleanup (closing files, etc.).
func basicDefer() {
	// This line will print LAST, just before basicDefer exits.
	defer fmt.Println("   [Deferred]: This prints at the end of the function.")

	fmt.Println("1. Doing some work...")
	fmt.Println("2. Doing more work...")
}

// 2. DEFER STACK (LIFO)
// When multiple defer calls are used, they are pushed onto a stack.
// They execute in Last-In, First-Out order.
func stackedDefers() {
	fmt.Println("Counting down in defer:")

	for i := 1; i <= 3; i++ {
		// These will print: 3, then 2, then 1
		defer fmt.Println("   Deferred count:", i)
	}

	fmt.Println("Loop finished. Now defers will execute...")
}

// 3. PANIC & RECOVER
// Panic: Stops ordinary control flow. It's like throwing an exception.
// Recover: Regains control of a panicking goroutine. It captures the panic value.
func executeRiskyOperation() {
	// IMPORTANT: 'recover' must be called inside a 'defer' function.
	// If we don't recover, the whole program crashes.
	defer func() {
		// recover() returns nil if there was no panic.
		if r := recover(); r != nil {
			fmt.Println("   ⚠️ RECOVERED from panic!")
			fmt.Println("   Error message was:", r)
		}
	}()

	fmt.Println("Start risky operation...")

	// Triggering a manual panic
	// In real life, this could be an index out of range or nil pointer dereference.
	panic("Something went terribly wrong!")

	// This line will NEVER be reached because of the panic above.
	fmt.Println("This line will not print.")
}

// ---------------------------------------------------------
// ⚠️ COMMON PITFALLS (Watch out!)
// ---------------------------------------------------------
//
// 1. Placing 'recover' outside of 'defer':
//    'recover()' only works when called *inside* a deferred function.
//    If you call it directly in normal code, it does nothing.
//
//    WRONG:
//    func bad() {
//        panic("boom")
//        recover() // Won't catch anything, program crashes.
//    }
//
// 2. os.Exit ignores defers:
//    If you call 'os.Exit(1)', the program terminates immediately,
//    and deferred functions are NOT run.
//
// 3. Defer arguments evaluation:
//    Arguments to deferred functions are evaluated when the defer statement
//    is executed, not when the function actually runs.
//    Example:
//    i := 0
//    defer fmt.Println(i) // Will print 0, even if you change i later.
//    i++
