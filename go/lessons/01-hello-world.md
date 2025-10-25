# Lesson 1: Hello World - Your First Go Program

## 🎯 Learning Objectives

By the end of this lesson, you'll understand:
- The structure of a Go program
- How to run Go code
- What packages are and why we need them
- The role of the `main` function and package
- How Go's compilation works

## 🤔 Why Start with Hello World?

"Hello World" is programming tradition because:
1. **Validates setup**: Confirms Go is installed correctly
2. **Introduces syntax**: Learn basic structure without complexity
3. **Builds confidence**: See immediate results
4. **Universal starting point**: Every programmer has been here

## 📝 The Code

Create a file named `hello.go`:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

That's it! Just 7 lines (including blank lines). Let's understand every single part.

## 🔍 Line-by-Line Breakdown

### Line 1: `package main`

**`package`**: Keyword that declares which package this file belongs to
- Every Go file must start with a package declaration
- Packages are Go's way of organizing code

**`main`**: The package name
- Special name for executable programs
- Libraries use other names (like `fmt`, `math`, `http`)

**Think of it like**: An apartment building. Each file is an apartment, and the package is the building it belongs to. The `main` building is special - it's where your program starts.

**Why do we need packages?**
- **Organization**: Group related code together
- **Reusability**: Use code from other packages
- **Namespace**: Avoid name conflicts (two packages can have same function name)

### Line 2: (Blank line)

**Why blank lines?** 
- Readability
- Go convention: blank line after package declaration
- `go fmt` (formatter) adds this automatically

### Line 3: `import "fmt"`

**`import`**: Keyword to bring in external packages
- Like borrowing tools from a toolbox
- Makes functions from other packages available

**`"fmt"`**: The format package from Go's standard library
- **fmt** stands for "format"
- Contains functions for formatted I/O (input/output)
- `Println` comes from this package

**The quotes**: Package names in quotes are from:
- Standard library (built into Go)
- External packages (downloaded separately)

**Without import**:
```go
package main

func main() {
    fmt.Println("Hello!")  // ERROR: undefined: fmt
}
```

### Line 4: (Blank line)

**Convention**: Blank line between imports and code

### Line 5: `func main() {`

**`func`**: Keyword that declares a function
- Functions are reusable blocks of code
- Like a recipe you can follow multiple times

**`main`**: The function name
- Special name: program execution starts here
- Every executable Go program must have a `main` function in the `main` package

**`()`**: Empty parentheses for parameters
- Functions can take inputs (we'll learn this in Lesson 4)
- `main` takes no parameters

**`{`**: Opening brace
- Marks start of function body
- Everything between `{` and `}` is part of the function

**Key insight**: 
```
package main  →  This is an executable program
func main()   →  This is where execution begins
```

### Line 6: `fmt.Println("Hello, World!")`

Let's break this into pieces:

**`fmt`**: The package name
- Refers to the package we imported

**`.`**: Dot operator
- Accesses something inside the package
- Like opening a toolbox and reaching for a specific tool

**`Println`**: A function from the `fmt` package
- **P** = Print
- **ln** = line (adds newline after printing)
- Capital 'P' means it's **exported** (usable outside the package)

**`("Hello, World!")`**: Calling the function with an argument
- `()` = function call
- `"Hello, World!"` = the text to print (a string)

**Why the capital P in `Println`?**
In Go, capitalization matters:
- **Lowercase**: Private (only usable within the package)
- **Uppercase**: Public (usable from other packages)

Examples:
```go
fmt.Println()  // Works - exported (public)
fmt.println()  // ERROR - doesn't exist (would be private)
```

**`"Hello, World!"`**: A string literal
- Text enclosed in double quotes
- Can contain any text

**Indentation**: Notice the line is indented
- Not required by Go (unlike Python)
- Strongly recommended for readability
- `go fmt` does this automatically

### Line 7: `}`

**`}`**: Closing brace
- Marks end of function body
- Must match the opening `{` on line 5

## 🚀 Running Your Program

### Method 1: `go run` (Quick Testing)

```bash
go run hello.go
```

**Output**:
```
Hello, World!
```

**What happened?**
1. Go compiled your code to machine code
2. Saved the executable to a temporary location
3. Ran the executable
4. Cleaned up the temporary file

**When to use**: 
- Testing and development
- Quick scripts
- Learning and experimenting

### Method 2: `go build` (Create Executable)

```bash
go build hello.go
```

**No output?** That's good! It means compilation succeeded.

Check what was created:
```bash
ls
# You'll see: hello (or hello.exe on Windows)
```

Run it:
```bash
./hello        # Linux/macOS
hello.exe      # Windows
```

**Output**:
```
Hello, World!
```

**What happened?**
1. Go compiled your code
2. Created a standalone executable file
3. You ran that executable

**When to use**:
- Creating programs to distribute
- Performance testing (compiled is slightly faster)
- When you'll run the program multiple times

### Method 3: `go install` (Install to GOPATH)

```bash
go install hello.go
```

This puts the executable in your `$GOPATH/bin` directory (if it's in your PATH, you can run it from anywhere).

**When to use**: Creating command-line tools for personal use

## 🔍 Deep Dive: What's Really Happening?

### The Compilation Process

```
Source Code (hello.go)
        ↓
    Go Compiler
        ↓
    Machine Code
        ↓
    Executable File
```

**Key Insight**: Go is a **compiled language**
- Code is translated to machine code before running
- No interpreter needed at runtime
- Results in fast execution
- Standalone binaries (no runtime dependencies)

**Compare to Python** (interpreted):
```
Source Code → Python Interpreter → Executes line by line
```

**Benefits of compilation**:
1. **Catch errors early**: Many bugs found at compile time
2. **Fast execution**: Native machine code
3. **Easy deployment**: One file, no dependencies
4. **Cross-compilation**: Build for any OS from any OS

### The Executable

Check the size:
```bash
ls -lh hello
# macOS/Linux: ~2MB
# Windows: ~2MB
```

**Why so large for such simple code?**
- Go includes a small runtime (garbage collector, goroutine scheduler)
- All dependencies bundled in (no external dependencies)
- Worth it: the executable runs anywhere without installing Go

## 🎨 Experimenting

### Experiment 1: Change the Message

```go
package main

import "fmt"

func main() {
    fmt.Println("Welcome to Go programming!")
}
```

### Experiment 2: Multiple Print Statements

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
    fmt.Println("Welcome to Go!")
    fmt.Println("This is fun!")
}
```

**Output**:
```
Hello, World!
Welcome to Go!
This is fun!
```

Each `Println` prints on its own line.

### Experiment 3: Print Without Newline

```go
package main

import "fmt"

func main() {
    fmt.Print("Hello, ")
    fmt.Print("World!")
    fmt.Println()  // Just adds a newline
}
```

**Output**:
```
Hello, World!
```

**Difference**:
- `Println` = Print with newline
- `Print` = Print without newline

### Experiment 4: Formatted Printing

```go
package main

import "fmt"

func main() {
    fmt.Printf("Hello, %s!\n", "Alice")
    fmt.Printf("I am %d years old\n", 25)
    fmt.Printf("%d + %d = %d\n", 2, 3, 5)
}
```

**Output**:
```
Hello, Alice!
I am 25 years old
2 + 3 = 5
```

**Format specifiers**:
- `%s` = string
- `%d` = decimal integer
- `\n` = newline

We'll learn more about formatting in later lessons!

## 🐛 Common Mistakes

### Mistake 1: Wrong Package Name

```go
package hello  // Wrong! Should be 'main' for executable

import "fmt"

func main() {
    fmt.Println("Hello!")
}
```

**Error**: `go run: cannot run non-main package`

**Fix**: Use `package main` for executable programs

### Mistake 2: Wrong Function Name

```go
package main

import "fmt"

func start() {  // Wrong! Should be 'main'
    fmt.Println("Hello!")
}
```

**Error**: `runtime.main_main·f: function main is undeclared in the main package`

**Fix**: Name the function `main`

### Mistake 3: Forgot to Import

```go
package main

func main() {
    fmt.Println("Hello!")  // ERROR: undefined: fmt
}
```

**Fix**: Add `import "fmt"`

### Mistake 4: Unused Import

```go
package main

import "fmt"
import "math"  // Imported but not used

func main() {
    fmt.Println("Hello!")
}
```

**Error**: `imported and not used: "math"`

**Why is this an error?** Go enforces clean code - no dead imports!

**Fix**: Remove unused imports or use them

### Mistake 5: Lowercase `Println`

```go
package main

import "fmt"

func main() {
    fmt.println("Hello!")  // Wrong: lowercase 'p'
}
```

**Error**: `undefined: fmt.println`

**Fix**: Use `fmt.Println` (capital P)

## 🎓 Key Takeaways

1. **Every Go file starts with** `package` declaration
2. **Executable programs** use `package main` and `func main()`
3. **Import packages** to use their functions
4. **Capital letters** mean exported (public)
5. **Go is compiled** - fast and creates standalone executables
6. **`go run`** for quick testing, **`go build`** for distribution
7. **Go enforces clean code** - no unused imports/variables

## 💪 Practice Exercises

### Exercise 1: Personalized Greeting
Modify the program to print:
```
Hello, [Your Name]!
Welcome to Go programming.
Have a great day!
```

### Exercise 2: ASCII Art
Print a simple ASCII art:
```
   ^_^
  /   \
```

### Exercise 3: Using Multiple Packages
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("Current time:")
    fmt.Println(time.Now())
}
```

Run this and see what happens!

### Exercise 4: Format Practice
Print your name, age, and city using `Printf`:
```
My name is Alice
I am 25 years old
I live in Seattle
```

## 🎯 Challenge

Create a program that:
1. Prints a welcome banner
2. Shows today's date (use `time` package)
3. Prints a motivational quote
4. Has at least 5 lines of output

Make it visually interesting!

## 📚 What's Next?

In [Lesson 2: Variables and Constants](./02-variables.md), you'll learn how to store and work with data.

---

**Congratulations!** 🎉 You've written and run your first Go program. You understand the basic structure and how Go programs work. You're officially a Go programmer now!
