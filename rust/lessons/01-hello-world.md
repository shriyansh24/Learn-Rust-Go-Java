# Lesson 1: Hello World - Your First Rust Program

## 🎯 Learning Objectives

By the end of this lesson, you'll understand:
- How to create a Rust project
- The structure of a Rust program
- How to compile and run Rust code
- What's happening "under the hood" with Hello World

## 🤔 Why Start with Hello World?

Every programming journey starts with "Hello World" because:
1. **It's Simple**: Focus on the process, not complex logic
2. **It's Complete**: A full program that does something
3. **It's Traditional**: Join millions of programmers in this ritual
4. **It Validates Setup**: Confirms everything is installed correctly

## 🏗️ Creating Your First Project

### Method 1: Using Cargo (Recommended)

```bash
cargo new hello_world
cd hello_world
```

**What just happened?**
- `cargo new` creates a new Rust project
- It generates a directory structure
- It initializes a git repository
- It creates a basic "Hello World" program for you

Let's explore what was created:

```
hello_world/
├── Cargo.toml          # Project configuration file
├── src/
│   └── main.rs         # Your source code
└── .git/               # Git repository (hidden)
```

### Understanding `Cargo.toml`

Open `Cargo.toml` and you'll see:

```toml
[package]
name = "hello_world"
version = "0.1.0"
edition = "2021"

[dependencies]
```

**What does this mean?**
- `[package]`: Metadata about your project
- `name`: The name of your program
- `version`: Follows semantic versioning (major.minor.patch)
- `edition`: Rust edition (like a language version, but backward compatible)
- `[dependencies]`: External libraries you'll use (empty for now)

**Think of it like**: Cargo.toml is like a recipe card - it tells Rust how to build your project.

## 📝 The Code

Open `src/main.rs`:

```rust
fn main() {
    println!("Hello, world!");
}
```

Only 3 lines! Let's understand each part:

### Line 1: `fn main() {`

**`fn`**: Keyword that declares a function
- Short for "function"
- Functions are reusable blocks of code

**`main`**: The name of this function
- Special name: `main` is where your program starts
- Every executable Rust program must have a `main` function
- Think of it as the "entrance door" to your program

**`()`**: Empty parentheses for parameters
- Functions can take inputs (called parameters)
- This function takes no inputs
- We'll learn more about parameters in Lesson 4

**`{`**: Opening curly brace
- Marks the start of the function's body
- Everything between `{` and `}` is inside the function

### Line 2: `println!("Hello, world!");`

**`println!`**: A macro that prints to the screen
- The `!` indicates it's a **macro**, not a regular function
- **Why a macro?** Macros can do things regular functions can't
- `ln` means "line" - it adds a newline after printing

**`("Hello, world!")`**: The message to print
- Text in quotes is a **string literal**
- The parentheses `()` pass the string to `println!`

**`;`**: Semicolon ends the statement
- Most lines in Rust end with semicolons
- It's like a period at the end of a sentence

### Line 3: `}`

**`}`**: Closing curly brace
- Marks the end of the function's body
- Must match the opening `{`

## 🚀 Running Your Program

### Three Ways to Run

**1. Cargo Run (Easiest)**
```bash
cargo run
```

Output:
```
   Compiling hello_world v0.1.0
    Finished dev [unoptimized + debuginfo] target(s) in 0.50s
     Running `target/debug/hello_world`
Hello, world!
```

**What happened?**
1. Cargo compiled your code
2. It created an executable
3. It ran the executable
4. Your message was printed

**2. Build Then Run (Two Steps)**
```bash
cargo build
./target/debug/hello_world
```

**Why two methods?**
- `cargo run` is convenient for development
- `cargo build` + run separately is useful when you want to run multiple times without rebuilding

**3. Using rustc Directly (Manual)**
```bash
rustc src/main.rs
./main
```

**When to use this?**
- Quick single-file programs
- Learning how compilation works
- Usually, use `cargo` instead

## 🔍 Deep Dive: What's Really Happening?

### Compilation Process

```
Source Code (main.rs)
        ↓
    Rust Compiler (rustc)
        ↓
    Machine Code (executable)
        ↓
    Your Computer Runs It
```

**Key Insight**: Rust is a **compiled language**
- Your code is translated to machine code BEFORE running
- Errors are caught during compilation, not at runtime
- The resulting program is fast (no interpreter overhead)

**Compare to Python** (interpreted):
```
Source Code → Python Interpreter → Runs line by line
```

### The Executable

After building, check `target/debug/`:
```bash
ls -lh target/debug/hello_world
```

This is a native executable:
- Contains machine code specific to your CPU
- Runs directly on your operating system
- Can be distributed and run without Rust installed
- Fast startup and execution

## 🎨 Experimenting

### Experiment 1: Change the Message

```rust
fn main() {
    println!("Welcome to Rust programming!");
}
```

Run it: `cargo run`

### Experiment 2: Multiple Print Statements

```rust
fn main() {
    println!("Hello, world!");
    println!("Welcome to Rust!");
    println!("This is fun!");
}
```

**Question**: How many lines of output? Answer: 3 (each `println!` adds a line)

### Experiment 3: Print Without Newline

```rust
fn main() {
    print!("Hello, ");
    print!("world!");
    println!();  // Just a newline
}
```

**Output**: `Hello, world!` (all on one line)

### Experiment 4: Formatting

```rust
fn main() {
    println!("I am learning {}", "Rust");
    println!("I am {} years old", 25);
    println!("{} + {} = {}", 2, 3, 5);
}
```

**`{}`** is a placeholder for values
- Like fill-in-the-blanks
- We'll learn more about formatting in later lessons

## 🐛 Common Mistakes

### Mistake 1: Forgetting the Semicolon

```rust
fn main() {
    println!("Hello, world!")  // Missing semicolon
}
```

**Error**: 
```
expected `;`, found `}`
```

**Fix**: Add `;` at the end

### Mistake 2: Wrong Function Name

```rust
fn start() {  // Wrong name!
    println!("Hello, world!");
}
```

**Error**: `main` function not found

**Fix**: Rename to `main`

### Mistake 3: Missing Exclamation Mark

```rust
fn main() {
    println("Hello, world!");  // Should be println!
}
```

**Error**: Cannot find function `println`

**Fix**: Use `println!` (with `!`)

### Mistake 4: Mismatched Quotes

```rust
fn main() {
    println!("Hello, world!');  // Mixed quotes
}
```

**Fix**: Use matching quotes (both `"` or both `'` for chars)

## 🎓 Key Takeaways

1. **Every Rust program starts at `main()`**
2. **`println!` is a macro for printing** (note the `!`)
3. **Semicolons end most statements**
4. **Rust is compiled, not interpreted**
5. **Cargo is your project manager** - use it!
6. **Error messages are helpful** - read them carefully

## 💪 Practice Exercises

### Exercise 1: Personalized Greeting
Modify your program to print your name:
```
Hello, [Your Name]!
Welcome to Rust programming.
```

### Exercise 2: ASCII Art
Print a simple ASCII art (like a smiley face):
```
  ^_^
```

### Exercise 3: Quote Printer
Print your favorite quote using multiple `println!` statements.

### Exercise 4: Math Announcement
Print a math fact using placeholders:
```
The square of 5 is 25
```

## 🎯 Challenge

Create a program that prints a short story or poem (at least 5 lines). Make it creative!

## 📚 What's Next?

In [Lesson 2: Variables and Mutability](./02-variables.md), you'll learn how to store and manipulate data.

---

**Congratulations!** 🎉 You've written and run your first Rust program. You understand the structure of a Rust program and how compilation works. You're now a Rust programmer!
