// Example 2: Variables and Constants
// Demonstrates Go's variable declaration styles

package main

import "fmt"

func main() {
    // Method 1: var keyword with type
    var name string = "Alice"
    fmt.Println("Name:", name)
    
    // Method 2: var keyword with type inference
    var age = 25
    fmt.Println("Age:", age)
    
    // Method 3: Short declaration (most common)
    city := "Seattle"
    fmt.Println("City:", city)
    
    // Multiple variables at once
    var x, y, z int = 1, 2, 3
    fmt.Println("x, y, z:", x, y, z)
    
    // Constants (cannot be changed)
    const pi = 3.14159
    const greeting = "Hello"
    fmt.Println("Pi:", pi)
    fmt.Println(greeting, "from constants!")
    
    // Variables can be reassigned (if not const)
    name = "Bob"
    fmt.Println("New name:", name)
    
    // Zero values (default values when not initialized)
    var defaultInt int         // 0
    var defaultString string   // ""
    var defaultBool bool       // false
    fmt.Printf("Defaults: %d, '%s', %t\n", defaultInt, defaultString, defaultBool)
}

/*
 * Key Points:
 * - Go has multiple ways to declare variables
 * - := is the most common (short declaration)
 * - Constants use 'const' keyword
 * - Uninitialized variables get zero values
 * - Go is statically typed but has type inference
 */
