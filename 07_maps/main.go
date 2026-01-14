package main

import "fmt"

// ---------------------------------------------------------
// TOPIC: Arrays vs Slices
// ---------------------------------------------------------

func main() {
	// ==========================================
	// PART 1: ARRAYS (Fixed Size)
	// ==========================================
	fmt.Println("--- ARRAYS ---")

	// Declaration: [Size]Type
	// The size is part of the type! [2]int and [3]int are different types.
	var arr [3]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	// arr[3] = 40 // COMPILE ERROR: Index out of bounds

	fmt.Printf("Array: %v | Len: %d\n", arr, len(arr))

	// Arrays are "Value Types".
	// Assigning an array to a new variable COPIES the whole data.
	arrCopy := arr
	arrCopy[0] = 999

	fmt.Println("Original Array:", arr)     // [10 20 30] (Unchanged)
	fmt.Println("Copied Array:  ", arrCopy) // [999 20 30]

	// ==========================================
	// PART 2: SLICES (Dynamic Wrapper)
	// ==========================================
	fmt.Println("\n--- SLICES ---")

	// Declaration: []Type (No size inside brackets)
	// A slice is a "window" or a "view" onto an underlying array.
	var slice []int = []int{10, 20, 30}

	// APPENDING
	// Use 'append' to add elements. It handles memory resizing automatically.
	// You MUST reassign the result back to the slice variable.
	slice = append(slice, 40)
	slice = append(slice, 50)

	fmt.Printf("Slice: %v\n", slice)

	// LEN vs CAP
	// len: How many elements are in the slice right now.
	// cap: How many elements fit in the underlying array before Go needs to create a new, bigger one.
	fmt.Printf("Len: %d | Cap: %d\n", len(slice), cap(slice))

	// ==========================================
	// PART 3: SLICING (Creating a sub-slice)
	// ==========================================
	fmt.Println("\n--- SLICING SYNTAX ---")

	// syntax: slice[start_inclusive : end_exclusive]
	numbers := []int{0, 1, 2, 3, 4, 5}

	subSlice := numbers[1:4] // Grabs indices 1, 2, and 3
	fmt.Printf("Original: %v\n", numbers)
	fmt.Printf("SubSlice[1:4]: %v\n", subSlice)

	// ---------------------------------------------------------
	// ⚠️ COMMON PITFALLS (Crucial!)
	// ---------------------------------------------------------
	fmt.Println("\n--- PITFALLS ---")

	// PITFALL 1: Slices share the same memory (Backing Array)
	// Unlike arrays, slices are cheap references.
	// If you modify a sub-slice, the original slice changes too!

	subSlice[0] = 999 // We modify the sub-slice...

	fmt.Println("After modifying subSlice:")
	fmt.Println("SubSlice:", subSlice) // [999 2 3]
	fmt.Println("Original:", numbers)  // [0 999 2 3 4 5] -> CHANGED! ⚠️

	// FIX for Pitfall 1:
	// If you need independent data, use 'copy()' or construct a new slice.

	// PITFALL 2: Append return value
	// Beginners often write: append(slice, 10)
	// This does nothing to 'slice' if the capacity changes.
	// ALWAYS write: slice = append(slice, 10)
}
