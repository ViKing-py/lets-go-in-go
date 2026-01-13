package main

import "fmt"

// ---------------------------------------------------------
// TOPIC: Loops (The 'for' keyword)
// ---------------------------------------------------------

func main() {
	// Go has only one looping keyword: "for".
	// There are no "while" or "do-while" loops, but "for" can do it all.

	// 1. THE CLASSIC LOOP (C-Style)
	// Structure: for init; condition; post { ... }
	fmt.Println("--- 1. Classic Loop ---")
	for i := 0; i < 5; i++ {
		fmt.Printf("Count: %d\n", i)
	}

	// 2. THE WHILE-STYLE LOOP
	// Structure: for condition { ... }
	// We use this when we don't need initialization or post-steps.
	fmt.Println("\n--- 2. While-Style Loop ---")
	counter := 3
	for counter > 0 {
		fmt.Println("Countdown:", counter)
		counter-- // Decrement manually inside the loop
	}

	// 3. THE INFINITE LOOP
	// Structure: for { ... }
	// This runs forever until you explicitly 'break' out of it.
	// Commonly used for servers or listening to channels.
	fmt.Println("\n--- 3. Infinite Loop ---")
	sum := 0
	for {
		sum++ // infinite increment
		if sum == 5 {
			fmt.Println("Reached 5, breaking out!")
			break // Exits the loop immediately
		}
	}

	// 4. RANGE LOOP (Iterating over data)
	// Used for slices, arrays, maps, strings, and channels.
	// Returns two values: index and value.
	fmt.Println("\n--- 4. Range Loop ---")
	fruits := []string{"Apple", "Banana", "Cherry"}

	for index, value := range fruits {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	// If you only need the value, use the blank identifier (_) to ignore the index.
	fmt.Println("only values:")
	for _, value := range fruits {
		fmt.Printf("Item: %s\n", value)
	}
}

// ---------------------------------------------------------
// ⚠️ COMMON PITFALLS
// ---------------------------------------------------------
//
// 1. Modifying the 'value' in range loops:
//    The 'value' variable in a range loop is a COPY of the element.
//    Modifying it does NOT change the original array/slice.
//
//    numbers := []int{1, 2, 3}
//    for _, v := range numbers {
//        v = v * 10 // This changes the local copy 'v', not the slice 'numbers'!
//    }
//    // 'numbers' is still [1, 2, 3]
//
// 2. Braces are mandatory:
//    Unlike C or Java, you cannot skip braces for a single-line loop.
//    WRONG:   for i := 0; i < 3; i++ fmt.Println(i)
//    CORRECT: for i := 0; i < 3; i++ { fmt.Println(i) }