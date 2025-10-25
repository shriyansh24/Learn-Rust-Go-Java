# Lesson 2: Variables and Constants

## 🎯 Learning Objectives

By the end of this lesson, you'll understand:
- How to declare variables in Go (three different ways!)
- The difference between variables and constants
- Go's type system and type inference
- Zero values (default values)
- Variable scope and naming conventions

## 🤔 What Are Variables?

**Real-world analogy**: Variables are like labeled boxes that hold values.

In a program, you need to remember information:
- User's name
- Score in a game
- Temperature reading
- Items in a shopping cart

Variables give names to these values so you can use them later.

## 📦 Declaring Variables: Three Ways

Go gives you **three different ways** to declare variables. Each has its use case.

### Method 1: Full Declaration with `var`

```go
var name string = "Alice"
var age int = 25
var height float64 = 5.6
var isStudent bool = true
```

**Syntax**: `var variableName type = value`

**When to use**:
- When you want to be explicit about the type
- In package-level declarations (outside functions)
- When readability is more important than brevity

### Method 2: Type Inference with `var`

```go
var name = "Alice"     // Go infers type: string
var age = 25           // Go infers type: int
var height = 5.6       // Go infers type: float64
var isStudent = true   // Go infers type: bool
```

**Syntax**: `var variableName = value` (type is inferred)

**How inference works**:
- Go looks at the value
- Determines the appropriate type
- Uses that type

**Inference rules**:
```go
var x = 42        // int (default for whole numbers)
var y = 42.0      // float64 (default for decimals)
var z = "hello"   // string
var w = true      // bool
```

### Method 3: Short Declaration with `:=`

```go
name := "Alice"
age := 25
height := 5.6
isStudent := true
```

**Syntax**: `variableName := value`

**This is the most common way in Go!**

**When to use**:
- Inside functions (can't use outside)
- Most of the time in everyday code
- When you want concise, readable code

**Restrictions**:
- Only works inside functions
- Can't be used for package-level variables
- At least one variable on the left must be new

### Comparing the Three Methods

```go
package main

import "fmt"

// Package-level: Must use 'var'
var globalMessage = "I'm global"

func main() {
    // Method 1: Explicit type
    var greeting string = "Hello"
    
    // Method 2: Type inference with var
    var name = "Alice"
    
    // Method 3: Short declaration (most common)
    age := 25
    
    fmt.Println(globalMessage)
    fmt.Println(greeting, name, "Age:", age)
}
```

## 🎨 Multiple Variable Declaration

### Declare Multiple Variables at Once

```go
// Method 1: Same line with var
var x, y, z int = 1, 2, 3

// Method 2: Same line with inference
var a, b, c = 1, "hello", true

// Method 3: Short declaration
name, age := "Alice", 25

// Method 4: Block declaration
var (
    firstName string = "John"
    lastName  string = "Doe"
    userAge   int    = 30
    isActive  bool   = true
)
```

**Example in action**:

```go
package main

import "fmt"

func main() {
    // Multiple variables at once
    x, y, z := 10, 20, 30
    fmt.Printf("x=%d, y=%d, z=%d\n", x, y, z)
    
    // Different types
    name, age, height := "Bob", 25, 5.9
    fmt.Printf("%s is %d years old and %.1f feet tall\n", 
               name, age, height)
    
    // Block declaration for readability
    var (
        username = "alice123"
        email    = "alice@example.com"
        verified = true
    )
    fmt.Printf("User: %s, Email: %s, Verified: %t\n", 
               username, email, verified)
}
```

## 🔄 Reassigning Variables

Variables can be changed after declaration:

```go
package main

import "fmt"

func main() {
    // Initial value
    score := 0
    fmt.Println("Score:", score)
    
    // Change value
    score = 10
    fmt.Println("Score:", score)
    
    // Change again
    score = 20
    fmt.Println("Score:", score)
    
    // Can use in calculations
    score = score + 5
    fmt.Println("Score:", score)
    
    // Shorthand operators
    score += 10  // Same as: score = score + 10
    fmt.Println("Score:", score)
}
```

**Important**: You can change the value, but NOT the type:

```go
age := 25       // type: int
age = 30        // OK: still an int
age = "thirty"  // ERROR: cannot use string as int
```

## 🧊 Zero Values (Default Values)

In Go, uninitialized variables get **zero values** (not undefined or null):

```go
package main

import "fmt"

func main() {
    var i int        // 0
    var f float64    // 0.0
    var b bool       // false
    var s string     // "" (empty string)
    
    fmt.Printf("int: %d\n", i)
    fmt.Printf("float: %f\n", f)
    fmt.Printf("bool: %t\n", b)
    fmt.Printf("string: '%s'\n", s)
}
```

**Output**:
```
int: 0
float: 0.000000
bool: false
string: ''
```

**Why zero values?**
- No "undefined" or "null pointer" errors
- Variables always have a valid value
- Makes code safer and more predictable

**Think of it like**: When you get a new notebook, it's blank (zero value), not "undefined".

## 🔒 Constants: Values That Never Change

Constants are like variables, but they can't be changed after creation.

### Declaring Constants

```go
const pi = 3.14159
const greeting = "Hello"
const daysInWeek = 7
const isEnabled = true
```

**Syntax**: `const name = value`

**Key differences from variables**:
1. Use `const` instead of `var` or `:=`
2. Cannot be changed after declaration
3. Must be assigned a value at declaration
4. Value must be known at compile time

### Constants in Action

```go
package main

import "fmt"

const (
    // Mathematical constants
    Pi = 3.14159
    E  = 2.71828
    
    // Application constants
    MaxUsers = 100
    AppName  = "MyApp"
)

func main() {
    fmt.Println("Pi:", Pi)
    fmt.Println("Max users:", MaxUsers)
    
    // This would cause an error:
    // Pi = 3.14  // ERROR: cannot assign to constant
    
    // Constants can be used in calculations
    radius := 5.0
    area := Pi * radius * radius
    fmt.Printf("Area of circle: %.2f\n", area)
}
```

### When to Use Constants

**Use constants when**:
- Value will never change (mathematical constants, configs)
- You want to prevent accidental modification
- Giving a name to a "magic number"

**Examples**:
```go
const (
    SecondsPerMinute = 60
    MinutesPerHour   = 60
    HoursPerDay      = 24
)

// Use them
const SecondsPerDay = SecondsPerMinute * MinutesPerHour * HoursPerDay
```

## 📛 Naming Conventions

Go has strong opinions about naming:

### The Rules (Enforced)

1. **Start with a letter or underscore**
2. **Can contain letters, digits, underscores**
3. **Case-sensitive**: `name` and `Name` are different

### The Conventions (Style)

**camelCase for local variables**:
```go
firstName := "John"
lastName := "Doe"
userAge := 30
isLoggedIn := true
```

**PascalCase for exported (public) names**:
```go
const MaxConnections = 100  // Exported (visible outside package)
const minConnections = 10   // Not exported (private to package)
```

**Short names in small scopes**:
```go
// OK in small scopes
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// Prefer descriptive names in larger scopes
userAccountBalance := 1000.50
```

**Acronyms stay uppercase**:
```go
userID := 123        // Not userId
httpClient := ...    // Not HTTPClient (unless exported)
HTTPURL := ...       // Exported
```

## 🎯 Type System

Go is **statically typed**: Every variable has a type that can't change.

### Common Types

```go
// Integer types (different sizes)
var age int = 25
var smallNum int8 = 127         // -128 to 127
var mediumNum int16 = 32000     // -32,768 to 32,767
var regularNum int32 = 2000000  // ~2 billion range
var bigNum int64 = 9000000000   // ~9 quintillion range

// Unsigned integers (no negative)
var count uint = 100            // 0 and positive only

// Floating point
var price float32 = 19.99       // 32-bit decimal
var pi float64 = 3.14159        // 64-bit decimal (more precision)

// Boolean
var isReady bool = true         // true or false

// String
var message string = "Hello"    // Text

// Rune (Unicode character)
var letter rune = 'A'           // Single character (int32 alias)

// Byte (8-bit unsigned integer)
var b byte = 255                // 0 to 255 (uint8 alias)
```

### Type Conversion

**Must be explicit** (no automatic conversion):

```go
package main

import "fmt"

func main() {
    // Automatic conversion doesn't work
    var i int = 42
    var f float64 = i  // ERROR: cannot use i (int) as float64
    
    // Must convert explicitly
    var i2 int = 42
    var f2 float64 = float64(i2)  // OK: explicit conversion
    fmt.Println(f2)
    
    // Converting between numeric types
    var x int32 = 100
    var y int64 = int64(x)
    var z float64 = float64(y)
    fmt.Printf("x=%d (int32), y=%d (int64), z=%.1f (float64)\n", x, y, z)
}
```

## 🐛 Common Mistakes

### Mistake 1: Using `:=` Outside Functions

```go
package main

name := "Alice"  // ERROR: non-declaration statement outside function body

func main() {
    // Use var instead at package level
}
```

**Fix**: Use `var name = "Alice"` at package level

### Mistake 2: Re-declaring with `:=`

```go
func main() {
    name := "Alice"
    name := "Bob"  // ERROR: no new variables on left side of :=
}
```

**Fix**: Use `=` for reassignment:
```go
name := "Alice"
name = "Bob"  // OK
```

### Mistake 3: Unused Variables

```go
func main() {
    name := "Alice"  // ERROR: name declared and not used
}
```

**Go enforces**: No unused variables! Either use it or remove it.

**Temporary workaround** (for development):
```go
func main() {
    name := "Alice"
    _ = name  // Blank identifier to "use" it temporarily
}
```

### Mistake 4: Trying to Change Constants

```go
const pi = 3.14
pi = 3.14159  // ERROR: cannot assign to constant
```

**Fix**: Use a variable if you need to change it

## 🎓 Key Takeaways

1. **Three ways to declare variables**: `var` with type, `var` with inference, `:=` short form
2. **`:=` is most common** but only works inside functions
3. **Zero values**: Uninitialized variables get safe defaults
4. **Constants use `const`** and can't be changed
5. **Go is statically typed** - type can't change after declaration
6. **No unused variables allowed** - Go keeps code clean
7. **Use camelCase for variables**, PascalCase for exports

## 💪 Practice Exercises

### Exercise 1: Temperature Converter
```go
package main

import "fmt"

func main() {
    // Declare celsius as 25
    // Convert to fahrenheit: F = C * 9/5 + 32
    // Print both values
}
```

### Exercise 2: Multiple Variables
Create variables for a person:
- name (string)
- age (int)
- height (float64)
- isStudent (bool)

Print them in a sentence.

### Exercise 3: Constants
Define constants for:
- Speed of light: 299792458 m/s
- Days in a year: 365
- Your favorite number

Use them in calculations.

### Exercise 4: Type Conversion
```go
// Start with an int
// Convert to float64
// Multiply by 3.14
// Convert back to int
// Print the result
```

## 🎯 Challenge

Create a program that:
1. Declares variables for a product (name, price, quantity)
2. Calculates total cost
3. Applies a 10% discount
4. Prints itemized receipt with formatting

Make it look professional!

## 📚 What's Next?

In [Lesson 3: Data Types](./03-data-types.md), you'll dive deeper into Go's type system and learn about more complex types.

---

**Great job!** 🎉 You now know how to work with variables and constants in Go. The three declaration methods might seem redundant, but each has its place. You'll develop intuition for which to use with practice!
