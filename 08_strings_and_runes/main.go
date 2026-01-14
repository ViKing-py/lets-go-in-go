package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// ---------------------------------------------------------
// TOPIC: Strings, Bytes, and Runes
// ---------------------------------------------------------

func main() {
	// 1. STRINGS ARE BYTE SLICES (READ-ONLY)
	// In Go, a string is a read-only slice of bytes.
	// It is NOT a slice of characters like in some other languages.

	// 'A' is a standard ASCII character (1 byte)
	// '世' is a Kanji character (3 bytes in UTF-8)
	s := "Hello, 世界" 

	fmt.Println("--- 1. Bytes vs Characters ---")
	fmt.Printf("String: %s\n", s)
	
	// len() returns the number of BYTES, not characters!
	// "Hello, " is 7 bytes. "世界" is 6 bytes (3 each). Total: 13.
	fmt.Printf("Length (bytes): %d\n", len(s))

	// If we iterate with a standard counter, we get individual bytes (uint8).
	fmt.Printf("Byte at index 7 (start of 世): %v\n", s[7]) 
	fmt.Println()

	// 2. WHAT IS A RUNE?
	// A 'rune' is an alias for 'int32'.
	// It represents a Unicode Code Point (a single character, regardless of how many bytes it takes).

	fmt.Println("--- 2. Iterating correctly (Range) ---")
	
	// The 'range' loop specifically handles UTF-8 decoding for strings.
	// It iterates over Runes, not Bytes.
	for index, char := range s {
		// %c prints the character, %d prints the byte position, %T prints the type
		fmt.Printf("%d: %c (Type: %T)\n", index, char, char)
	}
	
	// Note how the index jumps from 7 to 10. That's because '世' took bytes 7, 8, and 9.
	fmt.Println()

	// 3. COUNTING CHARACTERS
	// To count actual human-readable characters, do not use len().
	// Use the unicode/utf8 package.
	charCount := utf8.RuneCountInString(s)
	fmt.Printf("Actual character count: %d\n", charCount)
	fmt.Println()

	// 4. THE 'STRINGS' PACKAGE
	// This standard library package contains useful utilities.
	fmt.Println("--- 3. 'strings' Package Helpers ---")

	sample := "  Go Language  "

	// Trimming spaces
	trimmed := strings.TrimSpace(sample)
	fmt.Printf("Trimmed: '%s'\n", trimmed)

	// ToLower / ToUpper
	fmt.Println("Upper:", strings.ToUpper(trimmed))

	// Checking contents
	fmt.Println("Contains 'Go':", strings.Contains(trimmed, "Go"))
	
	// Replacing
	// -1 means replace ALL occurrences. 1 would replace only the first.
	replaced := strings.Replace(trimmed, "Language", "Gopher", -1)
	fmt.Println("Replaced:", replaced)

	// Splitting and Joining
	sentence := "a,b,c,d"
	parts := strings.Split(sentence, ",") // Returns a slice []string
	fmt.Printf("Split: %v\n", parts)

	joined := strings.Join(parts, "-")
	fmt.Println("Joined:", joined)
}

// ---------------------------------------------------------
// ⚠️ COMMON PITFALLS
// ---------------------------------------------------------
//
// 1. Strings are Immutable
//    You cannot change a specific character in a string by index.
//    
//    s := "Hello"
//    s[0] = 'h' // COMPILER ERROR: cannot assign to s[0]
//
//    FIX: You must convert it to a []rune or []byte, change it, and cast back,
//    or create a new string using string concatenation/replacement.
//
// 2. Using len() for text validation
//    If you are checking if a username is max 10 characters:
//    if len(username) > 10 { ... } 
//    This is risky if the user inputs emojis or non-English characters. 
//    "🇺🇦" (Ukrainian flag emoji) is 8 bytes long but looks like 1 character.
//
// 3. Single quotes vs Double quotes
//    "A" -> String (slice of bytes)
//    'A' -> Rune (int32)
//    They are not interchangeable types.