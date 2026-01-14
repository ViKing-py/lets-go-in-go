package main

import "fmt"

// ---------------------------------------------------------
// TOPIC: Maps (Hash Tables / Dictionaries)
// ---------------------------------------------------------

func main() {
	// 1. CREATING MAPS
	// A map maps keys to values.
	// Syntax: map[KeyType]ValueType

	// Method A: Using make() (Best for empty maps)
	// We must initialize the map before writing to it.
	userRoles := make(map[string]string)

	// Method B: Map Literal (Best if you have initial data)
	currencies := map[string]string{
		"USD": "US Dollar",
		"EUR": "Euro",
		"UAH": "Ukrainian Hryvnia", // Note the trailing comma!
	}

	fmt.Println("Initial currencies:", currencies)

	// 2. ADDING & UPDATING KEYS
	// If the key doesn't exist, it is added.
	// If the key exists, the value is overwritten.
	userRoles["admin"] = "Super User"
	userRoles["editor"] = "Content Manager"

	fmt.Println("User Roles:", userRoles)

	// 3. RETRIEVING VALUES & CHECKING EXISTENCE
	// This is specific to Go.

	// If we ask for a key that DOES NOT exist, Go returns the "zero value"
	// for that type (e.g., "" for string, 0 for int).
	role := userRoles["guest"]
	fmt.Printf("Role for guest: '%s' (This is empty string, not nil)\n", role)

	// The "Comma Ok" Idiom
	// To know if a key truly exists or if it's just a zero value, we use a second return variable.
	// val, ok := map[key]

	val, ok := userRoles["viewer"]
	if ok {
		fmt.Printf("Viewer exists: %s\n", val)
	} else {
		fmt.Println("Key 'viewer' does not exist in the map.")
	}

	// 4. DELETING KEYS
	// We use the built-in delete() function.
	delete(currencies, "USD")
	fmt.Println("Currencies after deletion:", currencies)

	// If you delete a key that doesn't exist, nothing happens (no error).
	delete(currencies, "NOT_EXISTING") // Safe operation
}

// ---------------------------------------------------------
// ⚠️ COMMON PITFALLS
// ---------------------------------------------------------
//
// 1. The NIL Map Panic (CRITICAL):
//    Declaring a map without initializing it creates a "nil" map.
//    You can read from a nil map, but writing to it causes a runtime PANIC.
//
//    WRONG:
//    var m map[string]int
//    m["key"] = 1 // PANIC! "assignment to entry in nil map"
//
//    CORRECT:
//    m := make(map[string]int)
//    m["key"] = 1
//
// 2. Random Iteration Order:
//    When you loop over a map using "range", the order is NOT guaranteed.
//    It is randomized intentionally by Go to prevent developers from relying on order.
//
// 3. Maps are Reference Types:
//    If you pass a map to a function and modify it inside that function,
//    the changes persist in the original map.
