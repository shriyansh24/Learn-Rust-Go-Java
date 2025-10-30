# Lesson 3: Control Flow - Making Decisions and Repeating Actions

## 🎯 Learning Objectives

By the end of this lesson, you will be able to:
-   **Understand the "why"**: Grasp why control flow is the backbone of any useful program.
-   **Make decisions**: Use `if`, `else if`, and `else` to control which code runs.
-   **Handle multiple conditions**: Use `switch` for cleaner, more readable logic.
-   **Repeat actions**: Master Go's powerful and unified `for` loop.
-   **Build a complete program**: Apply these concepts to build a fun, interactive Number Guessing Game.

## 🤔 The "Why": Beyond Straight-Line Code

So far, our programs have been like a straight road—they execute one line after another, from top to bottom. But what makes programming powerful is the ability to create roads that fork and loop.

-   **Why do we need forks?** (Conditional Logic)
    Imagine a login system. It needs to check if the password is correct. If it is, it should take the user to the home screen. If not, it should show an error. The program must **make a decision** based on input. That’s a fork in the road.

-   **Why do we need loops?** (Repetition)
    Imagine you need to process 1,000 files, or send a daily report every 24 hours. You don’t want to write the same code 1,000 times. You need a way to **repeat an action** until a condition is met. That's a loop.

**Control flow** is the set of tools that lets us direct the "flow" of our program's execution, turning simple scripts into intelligent applications.

## 🚀 Core Project: The Number Guessing Game

To escape "tutorial hell," we won't just learn concepts in isolation. We're going to build a complete, working program: **a number guessing game**.

**The goal of the game:**
1.  The program will pick a secret random number between 1 and 100.
2.  The player will have 10 chances to guess the number.
3.  After each guess, the program will tell the player if their guess was **too high** or **too low**.
4.  If the player guesses correctly, they win! If they run out of chances, they lose.

This single project will require every concept we learn in this lesson. Let's get started.

---

## 🚦 Part 1: Making Decisions with `if`, `else if`, `else`

The `if` statement is the most fundamental tool for decision-making. It checks if a condition is true and, if so, executes a block of code.

### The Basic `if`

```go
package main

import "fmt"

func main() {
    age := 20

    // The condition is 'age >= 18'
    if age >= 18 {
        // This code only runs if the condition is true
        fmt.Println("You are old enough to vote.")
    }
}
```

### Adding `else`: The "Otherwise" Clause

What if the condition is false? `else` lets you run a different block of code.

```go
    age := 16
    if age >= 18 {
        fmt.Println("You can vote.")
    } else {
        // This runs if the condition is false
        fmt.Println("You cannot vote yet.")
    }
```

### Chaining Conditions with `else if`

You can check multiple conditions in a sequence.

```go
    score := 85
    if score >= 90 {
        fmt.Println("Grade: A")
    } else if score >= 80 {
        fmt.Println("Grade: B")
    } else if score >= 70 {
        fmt.Println("Grade: C")
    } else {
        fmt.Println("Grade: D or F")
    }
```

### 🎮 Applying `if` to Our Guessing Game

Let's use this to give the player feedback. Assume the secret number is `42` and the player guesses `50`.

```go
    secretNumber := 42
    guess := 50

    if guess < secretNumber {
        fmt.Println("Too low!")
    } else if guess > secretNumber {
        fmt.Println("Too high!")
    } else {
        fmt.Println("You got it!")
    }
```
This simple block of code is the core logic of our game!

### Idiomatic Go: The `if` with a Short Statement

Go allows you to declare a variable right before the condition. This is extremely useful for keeping variables scoped only to where they are needed.

**Syntax**: `if initialization_statement; condition { ... }`

```go
// 'num' is created and is only accessible inside the 'if/else' block.
if num := 9; num < 0 {
    fmt.Println(num, "is negative")
} else if num < 10 {
    fmt.Println(num, "has 1 digit")
} else {
    fmt.Println(num, "has multiple digits")
}

// fmt.Println(num) // This would cause an ERROR: undefined: num
```

---

## 🔀 Part 2: Handling Multiple Cases with `switch`

A `switch` statement is a cleaner way to write a long chain of `if-else if` statements, especially when you're checking the value of a single variable.

### Basic `switch`

```go
package main

import "fmt"

func main() {
    day := "Wednesday"

    switch day {
    case "Monday":
        fmt.Println("Time to start the week!")
    case "Friday":
        fmt.Println("Weekend is almost here!")
    case "Sunday", "Saturday": // You can check multiple values
        fmt.Println("It's the weekend!")
    default: // This runs if no other case matches
        fmt.Println("It's a regular day.")
    }
}
```
**Key Go Feature**: Unlike many other languages, Go's `switch` cases do not "fall through" by default. The `break` is automatic, which prevents common bugs.

### `switch` without an expression

You can also use `switch` as a cleaner `if-else` chain. This is very common in Go.

```go
    score := 88

    switch {
    case score >= 90:
        fmt.Println("Excellent!")
    case score >= 80:
        fmt.Println("Good job.")
    default:
        fmt.Println("Keep practicing.")
    }
```

---

## 🔄 Part 3: Repeating Actions with the `for` Loop

### The "Why": Go's Unified `for` Loop

Many languages have multiple kinds of loops (`for`, `while`, `do-while`). Go's designers opted for simplicity: there is only one loop, the `for` loop, but it's incredibly flexible and can do everything the others can.

### Form 1: The Classic `for` Loop

This is the most common form, used when you know how many times you want to loop.

**Syntax**: `for initializer; condition; post-statement { ... }`

```go
// This loop will run 5 times
for i := 1; i <= 5; i++ {
    fmt.Println("Iteration number", i)
}
```

### Form 2: The "while" Loop

By omitting the initializer and post-statement, you get the equivalent of a `while` loop.

```go
    // This will loop as long as 'n' is less than 5
    n := 0
    for n < 5 {
        fmt.Println(n)
        n++
    }
```

### Form 3: The Infinite Loop

By omitting everything, you get an infinite loop. This is often used for servers that listen for requests forever, or for loops where the exit condition is complex and checked inside with a `break`.

```go
    for {
        fmt.Println("This will run forever!")
        // You need a way to get out!
        break // 'break' immediately exits the loop
    }
```
You can also use `continue` to skip the rest of the current iteration and move to the next one.

### 🎮 Applying `for` to Our Guessing Game

We need to give the player 10 chances. A classic `for` loop is perfect for this.

```go
// We'll give the player 10 guesses
for i := 1; i <= 10; i++ {
    fmt.Println("You have", 11-i, "guesses left.")
    // ... we'll put our guessing logic inside here ...
}
```

## 🧩 Putting It All Together: The Number Guessing Game

Now, let's combine everything we've learned to build the full game. Don't worry if this looks complex—read the comments carefully. The complete code is also available in `go/examples/03_guessing_game.go`.

```go
package main

import (
    "bufio"
    "fmt"
    "log"
    "math/rand"
    "os"
    "strconv"
    "strings"
    "time"
)

func main() {
    // 1. Generate a random secret number.
    // We need to seed the random number generator to make it different each time.
    rand.Seed(time.Now().UnixNano())
    secretNumber := rand.Intn(100) + 1

    fmt.Println("I've chosen a random number between 1 and 100.")
    fmt.Println("Can you guess it? You have 10 tries.")

    // 2. Set up a loop for 10 guesses.
    for guesses := 1; guesses <= 10; guesses++ {
        fmt.Printf("Guess #%d: ", guesses)

        // 3. Read the player's input from the console.
        reader := bufio.NewReader(os.Stdin)
        input, err := reader.ReadString('\n')
        if err != nil {
            log.Fatal(err) // Exit if there's an error reading input.
        }
        input = strings.TrimSpace(input)
        guess, err := strconv.Atoi(input)
        if err != nil {
            fmt.Println("Invalid input. Please enter a number.")
            continue // Skip the rest of this iteration.
        }

        // 4. Use `if/else` to compare the guess to the secret number.
        if guess < secretNumber {
            fmt.Println("Oops. Your guess was LOW.")
        } else if guess > secretNumber {
            fmt.Println("Oops. Your guess was HIGH.")
        } else {
            fmt.Println("Good job! You guessed it!")
            return // Exit the main function (and the program)
        }
    }

    // 5. If the loop finishes, the player didn't guess it.
    fmt.Println("Sorry, you didn't guess my number. It was:", secretNumber)
}
```

## 🎓 Key Takeaways

1.  **Control flow is essential**. It's what allows programs to make decisions and perform repetitive tasks.
2.  **Use `if/else` for decision-making**. The `if-short-statement` is a great way to keep variable scope tight.
3.  **Use `switch` for clean, multi-case logic**. It's often more readable than a long `if-else` chain.
4.  **Go has one loop: `for`**. It can be used as a classic `for` loop, a `while` loop, or an infinite loop. This simplicity is a core part of Go's design philosophy.
5.  **`break` exits a loop, `continue` skips to the next iteration**.

## 💪 Practice Exercises & Challenges

**Goal**: To truly learn this, you must write code. Try to solve these problems yourself before looking at the solutions in `go/examples/03_challenge_solutions.go`.

1.  **FizzBuzz**: Write a program that prints the numbers from 1 to 100. But for multiples of three, print "Fizz" instead of the number, and for the multiples of five, print "Buzz". For numbers which are multiples of both three and five, print "FizzBuzz".

2.  **Password Checker**: Write a program that checks if a user-entered password meets the following criteria:
    -   At least 8 characters long.
    -   Contains at least one uppercase letter.
    -   Contains at least one number.
    Provide feedback to the user on which criteria they are missing.

3.  **Sum of Numbers**: Write a program that calculates and prints the sum of all numbers from 1 to a given number `N`.

## 📚 What's Next?

You now know how to control the flow of your programs! But so far, we've only worked with simple data types like numbers and strings. In the next lesson, we'll learn how to work with collections of data using **Arrays, Slices, and Maps**.
