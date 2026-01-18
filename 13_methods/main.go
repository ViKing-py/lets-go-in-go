package main

import "fmt"

// ---------------------------------------------------------
// TOPIC: Methods (Value Receivers vs Pointer Receivers)
// ---------------------------------------------------------

// We define a simple struct to represent a User account.
type User struct {
	Username string
	Balance  int
}

// ---------------------------------------------------------
// 1. VALUE RECEIVER (func (u User))
// ---------------------------------------------------------
// - Creates a COPY of the struct.
// - Changes inside the method DO NOT affect the original struct.
// - Efficient for small structs (like points, dates) but bad for large ones (copying takes memory).

// ShowInfo uses a Value Receiver because it only READS data.
func (u User) ShowInfo() {
	fmt.Printf("User: %s | Balance: $%d\n", u.Username, u.Balance)
}

// TryToDeposit attempts to modify the balance using a Value Receiver.
// THIS WILL FAIL to update the original user.
func (u User) TryToDeposit(amount int) {
	u.Balance += amount
	fmt.Printf("  -> Inside TryToDeposit: Balance becomes $%d (This is a copy!)\n", u.Balance)
}

// ---------------------------------------------------------
// 2. POINTER RECEIVER (func (u *User))
// ---------------------------------------------------------
// - Passes the MEMORY ADDRESS of the struct.
// - Changes inside the method DO affect the original struct.
// - Efficient for large structs (no copying).

// Deposit uses a Pointer Receiver because it MODIFIES data.
func (u *User) Deposit(amount int) {
	u.Balance += amount
	fmt.Printf("  -> Inside Deposit: Balance becomes $%d (Original updated)\n", u.Balance)
}

// Rename also modifies the struct, so it needs a pointer.
func (u *User) Rename(newName string) {
	u.Username = newName
}

func main() {
	// Create a user instance
	// Note: We don't necessarily need to create it as a pointer (&User) initially.
	// Go handles the conversion automatically when calling methods.
	myUser := User{
		Username: "gopher123",
		Balance:  100,
	}

	fmt.Println("--- Initial State ---")
	myUser.ShowInfo()

	fmt.Println("\n--- Attempting update with VALUE RECEIVER ---")
	// Calling TryToDeposit (Value Receiver)
	// Go copies 'myUser' into 'u' inside the function.
	myUser.TryToDeposit(50)

	// Check if it changed
	fmt.Print("Result in main: ")
	myUser.ShowInfo() // Spoiler: It is still $100

	fmt.Println("\n--- Attempting update with POINTER RECEIVER ---")
	// Calling Deposit (Pointer Receiver)
	// Go passes the address of 'myUser'.
	myUser.Deposit(50)

	// Check if it changed
	fmt.Print("Result in main: ")
	myUser.ShowInfo() // Success: It is now $150

	fmt.Println("\n--- Renaming with POINTER RECEIVER ---")
	myUser.Rename("super_gopher")
	myUser.ShowInfo()
}

// ---------------------------------------------------------
// ⚠️ COMMON PITFALLS & BEST PRACTICES
// ---------------------------------------------------------
//
// 1. When to use Pointer Receivers (*T):
//    - If the method needs to modify the receiver (state mutation).
//    - If the struct is very large (e.g., contains a big array or image).
//      Copying a large struct is expensive; passing a pointer is cheap.
//    - If you want consistency: If some methods of the struct are pointers,
//      it's usually best practice to make ALL methods pointers for that struct.
//
// 2. When to use Value Receivers (T):
//    - If the struct is small (e.g., time.Time, simple Point{x, y}).
//    - If the struct is immutable (you never change it, only read it).
//    - If the type is a map, function, or channel (these are reference types
//      by definition, so they don't need *T to be modified internally,
//      though maps are rarely used as method receivers).
//
// 3. The "Nil" Pointer Trap:
//    - You CAN call a method on a nil pointer!
//    - Inside the method, you must check if the receiver is nil to avoid a panic.
//
//    Example:
//    func (u *User) IsRich() bool {
//        if u == nil { return false } // Safety check
//        return u.Balance > 1000
//    }
