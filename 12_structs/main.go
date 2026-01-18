package main

import (
	"encoding/json"
	"fmt"
)

// ---------------------------------------------------------
// TOPIC: Structs, Embedding, and Tags
// ---------------------------------------------------------

// 1. BASIC STRUCT DECLARATION
// A struct is a collection of fields. It's the core building block of data in Go.
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// 2. EMBEDDING (COMPOSITION)
// Go does not have class inheritance. Instead, it uses composition via embedding.
// By listing a type directly without a field name, fields are "promoted".
type Employee struct {
	Person   // Embedded struct (Anonymous field)
	JobTitle string
	Salary   int
}

// 3. STRUCT TAGS
// Tags are metadata attached to fields. They are accessed via reflection.
// Commonly used for JSON, database mapping (ORM), or validation.
type Product struct {
	ID          int     `json:"product_id"` // Rename field in JSON
	Name        string  `json:"name"`
	Description string  `json:"desc,omitempty"` // "omitempty": hide if empty string
	Price       float64 `json:"-"`              // "-": always ignore this field in JSON
	IsAvailable bool    // No tag: uses field name as is (IsAvailable)
}

func main() {
	fmt.Println("--- 1. Basic Structs ---")

	// Way A: explicit field names (Recommended)
	p1 := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}
	fmt.Println("Person 1:", p1)
	fmt.Println("Last Name:", p1.LastName)

	// Way B: implicit order (Not recommended for complex structs)
	p2 := Person{"Jane", "Smith", 25}
	fmt.Printf("Person 2: %+v\n", p2) // %+v prints field names too

	// -----------------------------------------------------

	fmt.Println("\n--- 2. Embedding & Promotion ---")

	emp := Employee{
		Person:   Person{FirstName: "Alice", LastName: "Wonder", Age: 40},
		JobTitle: "Engineer",
		Salary:   100000,
	}

	// You can access embedded fields directly (Promotion)
	// emp.Person.FirstName works, but emp.FirstName is shorter.
	fmt.Println("Employee Name:", emp.FirstName)
	fmt.Println("Job:", emp.JobTitle)

	// -----------------------------------------------------

	fmt.Println("\n--- 3. Anonymous Structs ---")

	// Useful for one-time data containers (e.g., inside a function or test)
	config := struct {
		Env  string
		Port int
	}{
		Env:  "Production",
		Port: 8080,
	}
	fmt.Printf("Config: %+v\n", config)

	// -----------------------------------------------------

	fmt.Println("\n--- 4. Struct Tags (JSON) ---")

	prod := Product{
		ID:          101,
		Name:        "Coffee Mug",
		Description: "",    // Empty, so it will be omitted in JSON
		Price:       12.99, // Ignored in JSON
		IsAvailable: true,
	}

	// Marshal converts the struct to JSON bytes
	jsonData, err := json.MarshalIndent(prod, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Convert bytes to string to print
	fmt.Println(string(jsonData))
}

// ---------------------------------------------------------
// ⚠️ COMMON PITFALLS
// ---------------------------------------------------------
//
// 1. Exported vs Unexported Fields (Visibility):
//    If a field name starts with a Lowercase letter (e.g., "age"),
//    it is PRIVATE to the package.
//
//    CRITICAL: The "encoding/json" package CANNOT see private fields.
//
//    type User struct {
//        password string // JSON marshal will result in empty/missing field!
//    }
//
// 2. The Ambiguity Problem in Embedding:
//    If you embed two structs that both have a field named "ID",
//    you cannot access ".ID" directly. You must be specific.
//
//    type A struct { ID int }
//    type B struct { ID int }
//    type C struct { A; B }
//
//    c := C{}
//    c.ID = 1 // COMPILER ERROR: ambiguous selector
//    c.A.ID = 1 // Correct
//
// 3. Modifying Structs in Functions:
//    Structs are value types. If you pass a struct to a function,
//    it is copied. To modify it, you MUST pass a pointer (*Struct).
