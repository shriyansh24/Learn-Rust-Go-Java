# Chapter 4: Complex Data Types: Arrays, Slices, and Maps

Welcome to Chapter 4! We've learned about basic data types. Now, let's explore how to work with collections of data using arrays, slices, and maps.

## 🎯 The Why: Why Do We Need Collections?

Imagine building a to-do list application. You need a way to store a list of tasks, which could be 5, 10, or 100 items long. Using a separate variable for each task would be impossible.

This is where collection types come in. They allow you to store and manage multiple values in a single variable.

**Think of it this way**: If a basic data type like an `int` is a single box, a collection is a shelf or a filing cabinet designed to hold many boxes in an organized way.

## 📦 The What: Arrays, Slices, and Maps

Go provides three primary collection types. Understanding the difference is crucial.

### Arrays: Fixed-Size Lists

An array is a numbered sequence of elements of a specific length.

- **Fixed Size**: Once you declare an array of a certain size, it can never grow or shrink.
- **Same Type**: All elements in an array must be of the same type.

**Syntax**: `[n]T`, where `n` is the number of elements and `T` is the type.

```go
var numbers [5]int // An array of 5 integers
```

### Slices: The Powerhouse

A slice is a more flexible view into the elements of an array. Slices are far more common in Go than arrays.

- **Dynamic Size**: Slices can grow and shrink.
- **Reference Type**: A slice doesn't store any data itself; it just describes a section of an underlying array. If you pass a slice to a function, changes made inside the function will be visible outside.

**Syntax**: `[]T`, where `T` is the type.

```go
var scores []int // A slice of integers
scores = []int{10, 20, 30} // A slice literal
```

### Maps: Key-Value Storage

A map is a collection of key-value pairs. It's an unordered group of elements of one type, called the element type, indexed by a set of unique keys of another type, called the key type.

- **Unordered**: The order of items in a map is not guaranteed.
- **Unique Keys**: Each key in a map must be unique.

**Syntax**: `map[KeyType]ValueType`.

```go
// A map with string keys and int values
var userAges map[string]int
```

## 🛠️ The How: Using Collections in Code

Let's see them in action.

```go
package main

import "fmt"

func main() {
    // --- ARRAYS ---
    // The size is part of the type. [5]int and [4]int are different types.
    var highScores [5]int
    highScores[0] = 98
    highScores[1] = 95
    fmt.Println("Array of high scores:", highScores)
    fmt.Println("Length of array:", len(highScores))

    // --- SLICES ---
    // A slice literal is like an array literal without the length.
    tasks := []string{"learn go", "build a project", "profit"}
    fmt.Println("Slice of tasks:", tasks)

    // Slices can be appended to. This is a very common operation.
    tasks = append(tasks, "read a book")
    fmt.Println("Appended slice:", tasks)
    fmt.Println("Length of slice:", len(tasks))

    // --- MAPS ---
    // Use `make` to create an empty map
    userEmails := make(map[string]string)
    userEmails["Alice"] = "alice@example.com"
    userEmails["Bob"] = "bob@example.com"
    fmt.Println("Map of user emails:", userEmails)

    // Get a value from a map
    emailOfAlice := userEmails["Alice"]
    fmt.Println("Email of Alice:", emailOfAlice)

    // Delete a key-value pair
    delete(userEmails, "Bob")
    fmt.Println("Map after deleting Bob:", userEmails)

    // Checking if a key exists
    email, ok := userEmails["Charlie"]
    if ok {
        fmt.Println("Charlie's email is:", email)
    } else {
        fmt.Println("Charlie's email not found.")
    }
}
```

**Run this code!**

## 🤔 The Intuition: When to Use Which?

- **Use an Array** when you know the exact, fixed number of elements you need. This is rare. A common example is representing an RGB color with `[3]int{red, green, blue}`.
- **Use a Slice** almost every other time you need a list of things. It's the most common collection type in Go. If you think "I need a list/array," your first thought should be a slice.
- **Use a Map** when you need to associate values with keys, like looking up a user's ID by their username or storing configuration settings.

## ⚠️ Common Pitfalls

1.  **Array vs. Slice Declaration**: `[5]int` is an array. `[]int` is a slice. Forgetting the length creates a slice. They are not interchangeable.
2.  **Zero Value of a Map**: The zero value of a map is `nil`. You can't add keys to a `nil` map. You must initialize it with `make()` first.
    ```go
    var m map[string]int
    // m["key"] = 1 // This will cause a runtime panic!
    ```
3.  **Appending to a Slice**: The `append` function returns a *new* slice. You must assign the result back to the original slice variable: `mySlice = append(mySlice, newValue)`.

## 🚀 Practice: Exercises

1.  **To-Do List**: Create a slice of strings to represent a to-do list. Add a few tasks, print the list, then remove the *second* task from the list and print it again. (Hint: `append` can be used to remove an element by slicing *around* it: `s = append(s[:i], s[i+1:]...)`).
2.  **Word Count**: Create a program that counts the occurrences of each word in a given string. Use a map to store the word counts.

## ✅ Solutions

<details>
<summary>Exercise 1 Solution</summary>

```go
package main

import "fmt"

func main() {
    todoList := []string{"buy milk", "walk the dog", "write code"}
    fmt.Println("Initial list:", todoList)

    // To remove the second item (at index 1)
    todoList = append(todoList[:1], todoList[2:]...)
    fmt.Println("Updated list:", todoList)
}
```
</details>

<details>
<summary>Exercise 2 Solution</summary>

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "go is awesome go is fun go go go"
	words := strings.Fields(text) // splits the string into a slice of words

	wordCounts := make(map[string]int)

	for _, word := range words {
		wordCounts[word]++
	}

	fmt.Println("Word counts:", wordCounts)
}
```
</details>

---

Fantastic! You now have the tools to manage collections of data, which is a huge step forward. Next up, we'll see how to better organize our code by creating our own functions and packages.
