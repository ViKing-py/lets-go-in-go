package main

import "fmt"

// ---------------------------------------------------------
// TOPIC: Functions
// ---------------------------------------------------------

// 1. BASIC PARAMETERS & TYPE OMISSION
// Standard function: takes arguments, returns a value.
// Note: When two or more consecutive named parameters share a type,
// you can omit the type from all but the last one.
// (x int, y int) -> (x, y int)
func add(x, y int) int {
	return x + y
}

// 2. MULTIPLE RETURN VALUES
// Go functions can return any number of results.
// This is often used to return a result AND an error (or a boolean check).
func divide(dividend, divisor int) (int, int) {
	quotient := dividend / divisor
	remainder := dividend % divisor
	return quotient, remainder
}

// 3. NAMED RETURN VALUES (NAKED RETURN)
// Return values can be named at the top of the function.
// They are treated as variables defined at the start of the function.
// A "return" statement without arguments returns the named values.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// "return" without arguments is a "naked" return.
	// It returns the current values of x and y automatically.
	return
}

// 4. VARIADIC FUNCTIONS
// A function that can be called with any number of trailing arguments.
// Inside the function, the param 'nums' becomes a slice of ints ([]int).
func sumAll(nums ...int) int {
	fmt.Printf("Received %d numbers to sum. Type of nums: %T\n", len(nums), nums)

	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println("--- 1. Basic Functions ---")
	result := add(42, 13)
	fmt.Println("Sum:", result)

	fmt.Println("\n--- 2. Multiple Return Values ---")
	// We capture both return values into variables q and r
	q, r := divide(17, 5)
	fmt.Printf("17 divided by 5 is %d with a remainder of %d\n", q, r)

	// ignoring one value using blank identifier "_"
	q2, _ := divide(10, 2)
	fmt.Println("Only interested in quotient:", q2)

	fmt.Println("\n--- 3. Named (Naked) Returns ---")
	x, y := split(17)
	fmt.Printf("Split 17 into: %d and %d\n", x, y)

	fmt.Println("\n--- 4. Variadic Functions ---")
	// Call with individual arguments
	t1 := sumAll(1, 2)
	t2 := sumAll(10, 20, 30, 40, 50)

	fmt.Println("Total 1:", t1)
	fmt.Println("Total 2:", t2)

	// Call with a slice
	// If you already have a slice, use "..." to spread it into the function
	numbers := []int{100, 200, 300}
	t3 := sumAll(numbers...)
	fmt.Println("Total from slice:", t3)
}

// ---------------------------------------------------------
// ⚠️ COMMON PITFALLS
// ---------------------------------------------------------
//
// 1. Naked Returns readability:
//    While named returns are cool, using naked returns in long functions
//    can harm readability. You don't know exactly what is being returned
//    unless you scroll back up to the signature.
//    BEST PRACTICE: Use them only in short functions.
//
// 2. Unused return values:
//    If a function returns multiple values, you MUST handle all of them.
//    You cannot assign 2 return values to 1 variable.
//
//    WRONG:
//    val := divide(10, 2) // Error: divide returns 2 values
//
//    CORRECT (if you want to ignore one):
//    val, _ := divide(10, 2)
//
// 3. Variadic Arguments vs Slices:
//    You cannot pass a slice directly to a variadic function without unpacking.
//
//    nums := []int{1, 2, 3}
//    sumAll(nums)    // Error: cannot use nums (type []int) as type int
//    sumAll(nums...) // Correct: expands the slice into individual arguments
