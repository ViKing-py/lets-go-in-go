package main

import (
	"fmt"
	"time"
)

// ---------------------------------------------------------
// TOPIC: Control Flow (If/Else & Switch)
// ---------------------------------------------------------

func main() {
	fmt.Println("--- 1. Standard If / Else ---")
	
	age := 18

	// Basic if-else structure
	// Note: You don't need parentheses ( ) around the condition.
	if age < 18 {
		fmt.Println("You are a minor.")
	} else if age == 18 {
		fmt.Println("You just became an adult!")
	} else {
		fmt.Println("You are an adult.")
	}

	fmt.Println("\n--- 2. If with Short Statement (Initialization) ---")
	
	// Go allows you to execute a short statement BEFORE the condition.
	// Syntax: if <statement>; <condition> { ... }
	// Common use case: Error handling or checking map keys.
	
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "is single digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
	
	// Note: 'num' is ONLY available inside this if/else block.
	// fmt.Println(num) // This would cause an error here!

	fmt.Println("\n--- 3. Basic Switch Statement ---")

	day := "Monday"

	// Unlike C or Java, you do NOT need 'break' statements.
	// Go breaks automatically after a match.
	switch day {
	case "Saturday", "Sunday": // Multiple values in one case
		fmt.Println("It's the weekend!")
	case "Monday":
		fmt.Println("It's the start of the work week.")
	default:
		fmt.Println("Just another work day.")
	}

	fmt.Println("\n--- 4. Tagless Switch (Cleaner If-Else) ---")

	// Switch without a variable acts like a long chain of if-else.
	// It's often cleaner to read than many if-else statements.
	hour := time.Now().Hour()

	switch {
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	fmt.Println("\n--- 5. The 'fallthrough' keyword ---")

	// fallthrough forces execution of the NEXT case,
	// IGNORING the condition of that next case.
	
	score := 50
	fmt.Println("Score evaluation:")

	switch {
	case score >= 50:
		fmt.Print("You passed. ")
		fallthrough // Continues to the next case immediately
	case score >= 100: // Logical paradox: 50 is not >= 100, but it prints anyway!
		fmt.Println("Wait, fallthrough executed this line anyway!")
	default:
		fmt.Println("You failed.")
	}
}

// ---------------------------------------------------------
// ⚠️ COMMON PITFALLS
// ---------------------------------------------------------
//
// 1. Variable Scope in "If with Init":
//    Variables declared in the if statement (like 'num' above) 
//    die immediately after the else block closes.
//    Attempting to access 'num' later in the code will fail.
//
// 2. 'fallthrough' logic is dangerous:
//    When using 'fallthrough', Go executes the next case explicitly
//    WITHOUT checking if the next case matches the condition.
//    Use it very rarely!
//
// 3. Opening Brace Placement:
//    Just like functions, the '{' for if/switch must be on the same line.
//    WRONG:
//    if x > 0
//    { ... }