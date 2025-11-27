# Chapter 5: Functions & Packages: Organizing Your Code

Welcome to Chapter 5! We've learned about functions, but how do we organize them as our projects grow? The answer is **packages**. This chapter covers both.

## 🎯 The Why: Why Use Functions?

Imagine you're writing a program, and you find yourself writing the same piece of code over and over again. Not only is this tedious, but if you need to change that code, you have to find and update every single copy.

Functions solve this problem. They allow you to group a set of statements together and give them a name. You can then "call" that function whenever you need to run that code.

**Think of it this way**: A function is like a recipe. You write it down once, and then you can use it as many times as you want without having to rewrite the instructions every time.

## 📦 The What: Anatomy of a Go Function

A function in Go is defined with the `func` keyword. Here's a basic example:

```go
func sayHello() {
    // This is the function body
    fmt.Println("Hello, Go!")
}
```

This function is named `sayHello`, it takes no arguments, and it doesn't return any values.

### Functions with Parameters

You can pass data to a function through parameters.

```go
func greet(name string) {
    fmt.Println("Hello,", name)
}
```

This function, `greet`, takes one argument: a `string` named `name`.

### Functions with Return Values

Functions can also return values.

```go
func add(a int, b int) int {
    return a + b
}
```

This function, `add`, takes two `int` arguments and returns an `int`.

### Multiple Return Values

A unique feature of Go is that functions can return multiple values. This is often used to return a result and an error.

```go
func divide(a int, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}
```

## 🛠️ The How: Using Functions in Your Code

Let's put it all together in a program.

```go
package main

import "fmt"

// A simple function with no parameters or return values
func sayHello() {
    fmt.Println("Hello, Go!")
}

// A function that takes a string parameter
func greet(name string) {
    fmt.Println("Hello,", name)
}

// A function that takes two integers and returns an integer
func add(a int, b int) int {
    return a + b
}

// A function that returns multiple values
func divide(a int, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    sayHello()

    greet("Alice")
    greet("Bob")

    sum := add(5, 3)
    fmt.Println("5 + 3 =", sum)

    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("10 / 2 =", result)
    }

    result, err = divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("10 / 0 =", result)
    }
}
```

**Run this code!**

## 🤔 The Intuition: Functions as Building Blocks

Think of functions as the Lego bricks of your program. You can create small, simple functions that do one thing well, and then you can combine them to build larger, more complex programs.

This makes your code:
- **Easier to read**: Small functions are easier to understand than one giant block of code.
- **Easier to test**: You can test each function individually to make sure it's working correctly.
- **Easier to reuse**: You can use the same function in many different places in your program.

## ⚠️ Common Pitfalls

1. **Unused Return Values**: If a function returns a value, you must use it or explicitly ignore it with the blank identifier `_`.
   ```go
   // This will cause a compile error
   // add(5, 3)

   // This is okay
   _ = add(5, 3)
   ```

2. **Parameter Mismatch**: You must call a function with the correct number and types of arguments.
   ```go
   // This will cause a compile error
   // greet(123)
   ```

## 🚀 Practice: Exercises

1. **Greeting Function**: Write a function that takes a name and an age and returns a formatted string, like "Hello, my name is Alice and I am 30 years old."
2. **Calculator**: Write functions for addition, subtraction, multiplication, and division. Then, use them to build a simple calculator.

## ✅ Solutions

<details>
<summary>Exercise 1 Solution</summary>

```go
package main

import "fmt"

func formatGreeting(name string, age int) string {
    return fmt.Sprintf("Hello, my name is %s and I am %d years old.", name, age)
}

func main() {
    greeting := formatGreeting("Alice", 30)
    fmt.Println(greeting)
}
```
</details>

<details>
<summary>Exercise 2 Solution</summary>

```go
package main

import "fmt"

func add(a, b float64) float64 {
    return a + b
}

func subtract(a, b float64) float64 {
    return a - b
}

func multiply(a, b float64) float64 {
    return a * b
}

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    fmt.Println("10 + 5 =", add(10, 5))
    fmt.Println("10 - 5 =", subtract(10, 5))
    fmt.Println("10 * 5 =", multiply(10, 5))

    result, err := divide(10, 5)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("10 / 5 =", result)
    }
}
```
</details>

---

## 📦 Part 2: Packages - Organizing Our Codebase

As your project grows, putting all your functions in a single `main.go` file becomes messy. Packages are Go's way of organizing and reusing code. A package is simply a directory containing one or more Go files.

### The Why: Why Use Packages?

- **Organization**: Group related code together (e.g., all database logic in a `database` package).
- **Reusability**: Use the same package in multiple different projects.
- **Encapsulation**: Hide implementation details and only expose a clean public API.

### The How: Creating and Using a Custom Package

Let's refactor our calculator into a `calculator` package.

**1. Project Structure:**

First, we need a Go module. A module is a collection of related Go packages that are versioned together as a single unit.

Create a new directory for our project, say `myproject`, and inside it, run:
```bash
go mod init myproject
```

This creates a `go.mod` file. Now, create the following file structure:

```
myproject/
├── go.mod
├── main.go
└── calculator/
    └── calculator.go
```

**2. Create the Package:**

Place the following code in `myproject/calculator/calculator.go`:

```go
// The package clause defines the package this file belongs to.
package calculator

// Add is a function. Because it starts with a capital letter,
// it is an "exported" function and can be used by other packages.
func Add(a, b float64) float64 {
    return a + b
}

// subtract starts with a lowercase letter, so it is "unexported".
// It can only be used by other functions inside the `calculator` package.
func subtract(a, b float64) float64 {
    return a - b
}
```

**Key Concept: Exported vs. Unexported**
In Go, if a name (like a function or variable) starts with a capital letter, it is **exported**—visible and usable from other packages. If it starts with a lowercase letter, it is **unexported**—private to its own package.

**3. Use the Package:**

Now, let's use our new `calculator` package in `myproject/main.go`:

```go
package main

import (
	"fmt"
	"myproject/calculator" // Import our custom package
)

func main() {
    result := calculator.Add(10, 5)
    fmt.Println("10 + 5 =", result)

    // The following line would cause a compile error because subtract is not exported:
    // result = calculator.subtract(10, 5)
}
```

To run this, navigate to the `myproject` directory in your terminal and run `go run .`. Go will automatically find and build the files.

### 🤔 The Intuition: Packages as Toolboxes

Think of packages as toolboxes. The `fmt` package is a toolbox for formatting text. Your `calculator` package is a toolbox for doing math. Exported functions are the tools you can see and use. Unexported functions are the internal machinery that makes the tools work but which you don't need to worry about.

### 🚀 Practice: A New Package

1.  **Greeter Package**: Create a new package named `greeter`.
2.  Inside the `greeter` package, create an exported function `SayHello` that takes a `name` (string) and returns a greeting string.
3.  Use your `greeter` package in `main.go` to print a greeting.

You're doing great! You've learned how to create and use functions and how to organize them into packages. This is a fundamental skill for building any real Go application. Next, we'll learn about pointers and how Go manages memory.
