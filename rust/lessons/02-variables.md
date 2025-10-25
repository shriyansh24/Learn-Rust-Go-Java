# Lesson 2: Variables and Mutability

## 🎯 Learning Objectives

By the end of this lesson, you'll understand:
- What variables are and why we need them
- Rust's unique approach: immutable by default
- How to make variables mutable
- Variable naming conventions
- The concept of shadowing

## 🤔 What Are Variables?

**Real-world analogy**: Variables are like labeled boxes that hold values.

```
┌─────────────────┐
│  age            │
│  [25]           │
└─────────────────┘
```

You can:
- Put a value in the box (assignment)
- Look at what's in the box (reading)
- Replace the value (mutation, if allowed)

**In programming**: Variables give names to values so you can:
- Remember values for later use
- Make code readable (better than magic numbers)
- Reuse values throughout your program

## 📦 Creating Variables

### Basic Syntax

```rust
fn main() {
    let x = 5;
    println!("The value of x is: {}", x);
}
```

**Breaking it down**:
- `let`: Keyword that declares a variable
- `x`: The variable name (the label on the box)
- `=`: Assignment operator (puts value in the box)
- `5`: The value being stored
- `;`: Ends the statement

### More Examples

```rust
fn main() {
    let age = 25;
    let name = "Alice";
    let is_student = true;
    let pi = 3.14;
    
    println!("Name: {}", name);
    println!("Age: {}", age);
    println!("Is student: {}", is_student);
    println!("Pi: {}", pi);
}
```

## 🔒 Immutability: Rust's Safety Feature

### The Surprise: Variables Are Immutable by Default

Try this:

```rust
fn main() {
    let x = 5;
    println!("x is: {}", x);
    
    x = 6;  // This will cause an error!
    println!("x is now: {}", x);
}
```

**Error**:
```
cannot assign twice to immutable variable `x`
```

**Wait, what?** In most languages, variables can be changed. Why does Rust prevent this?

### The Reasoning Behind Immutability

**Problem in traditional languages**:
```rust
// Imagine a large program
let config_value = 100;

// ... hundreds of lines of code ...

// Someone changes it
config_value = 200;  // In other languages, this might be allowed

// ... more code that assumes it's still 100 ...
```

**Issues**:
1. Hard to track where values change
2. Bugs from unexpected mutations
3. Race conditions in concurrent code
4. Harder to reason about code

**Rust's Solution**: Make variables immutable by default
- **Forces you to be explicit** about what can change
- **Prevents accidental modifications**
- **Makes code easier to understand** (if immutable, value never changes)
- **Enables compiler optimizations**
- **Safer concurrent programming**

### The Mental Model

Think of immutable variables as **constants with a name computed at runtime**:

```rust
let x = expensive_calculation();  // Result never changes after this
```

Once `x` is set, you (and the compiler) know it will never change. This is powerful for reasoning about your code.

## 🔓 Mutability: When You Need to Change

### Making Variables Mutable

Use the `mut` keyword:

```rust
fn main() {
    let mut x = 5;
    println!("x is: {}", x);
    
    x = 6;  // Now this works!
    println!("x is now: {}", x);
}
```

**Output**:
```
x is: 5
x is now: 6
```

### When to Use `mut`?

**Use mutable variables when**:
- Building up a value incrementally
- Accumulating results in a loop
- Updating state based on conditions

**Examples**:

```rust
fn main() {
    // Accumulator
    let mut sum = 0;
    sum += 1;
    sum += 2;
    sum += 3;
    println!("Sum: {}", sum);  // Sum: 6
    
    // Counter
    let mut count = 0;
    count += 1;
    count += 1;
    println!("Count: {}", count);  // Count: 2
    
    // Status tracker
    let mut is_active = true;
    is_active = false;
    println!("Active: {}", is_active);  // Active: false
}
```

### Best Practice

**Start with immutable, add `mut` only when needed**:

```rust
// Good: Clear what changes
let max_size = 100;          // Won't change
let mut current_size = 0;    // Will change

// Avoid: Everything mutable "just in case"
let mut max_size = 100;      // Unnecessarily mutable
let mut current_size = 0;
```

## 🎭 Shadowing: A Unique Feature

### What is Shadowing?

You can declare a new variable with the same name:

```rust
fn main() {
    let x = 5;           // First x
    let x = x + 1;       // Second x (shadows the first)
    let x = x * 2;       // Third x (shadows the second)
    
    println!("x is: {}", x);  // x is: 12
}
```

**What happened?**
1. First `x` = 5
2. Second `x` = 5 + 1 = 6 (first `x` is shadowed)
3. Third `x` = 6 * 2 = 12 (second `x` is shadowed)

### Shadowing vs. Mutation

These are **different**:

```rust
// Shadowing: Creates a new variable
let x = 5;
let x = x + 1;  // New variable, same name

// Mutation: Changes existing variable
let mut x = 5;
x = x + 1;      // Same variable, new value
```

**Key difference**: Shadowing creates a new variable, so you can even change the type!

```rust
fn main() {
    let spaces = "   ";           // Type: &str (string)
    let spaces = spaces.len();    // Type: usize (number)
    println!("Number of spaces: {}", spaces);  // Works!
}
```

Try with `mut`:
```rust
fn main() {
    let mut spaces = "   ";
    spaces = spaces.len();  // ERROR: Can't change type!
}
```

**Error**: expected `&str`, found `usize`

### When to Use Shadowing?

**Use shadowing when**:
1. Transforming a value through multiple steps
2. Converting between types
3. Using the same logical name for different representations

**Example: Data transformation**:

```rust
fn main() {
    // Read user input as string
    let input = "42";
    
    // Parse to number
    let input = input.parse::<i32>().unwrap();
    
    // Calculate
    let input = input * 2;
    
    println!("Result: {}", input);  // Result: 84
}
```

Each `input` is a logical step in processing, but a different variable.

## 📛 Naming Conventions

### The Rules (Enforced by Compiler)

1. **Must start with** a letter or underscore
2. **Can contain** letters, digits, and underscores
3. **Cannot be** a keyword (like `let`, `fn`, `if`)

```rust
// Valid
let age = 25;
let first_name = "Alice";
let _internal_value = 10;
let x1 = 5;

// Invalid
let 1st = 5;        // Can't start with digit
let first-name = "Alice";  // Can't use hyphen
let let = 5;        // 'let' is a keyword
```

### The Conventions (Style Guide)

**snake_case for variable names**:
```rust
let user_name = "Bob";        // Good
let userName = "Bob";         // Bad (not Rust style)
let UserName = "Bob";         // Bad (this is for types)
```

**Why snake_case?**
- Rust convention (all Rust code follows this)
- Easy to read
- Consistent with standard library
- Clippy (linter) will warn you

**Descriptive names**:
```rust
// Good: Clear meaning
let user_age = 25;
let max_connections = 100;
let is_authenticated = false;

// Bad: Unclear
let x = 25;
let n = 100;
let flag = false;
```

**Exception**: Short names in small scopes:
```rust
// OK for loop indices
for i in 0..10 {
    println!("{}", i);
}

// OK for math operations
let x = 5;
let y = 10;
let distance = (x * x + y * y).sqrt();
```

## 🎯 Practical Example

```rust
fn main() {
    // User profile (immutable data)
    let username = "alice42";
    let email = "alice@example.com";
    let account_created = "2024-01-15";
    
    // Game state (mutable data)
    let mut score = 0;
    let mut level = 1;
    let mut lives = 3;
    
    println!("Welcome, {}!", username);
    println!("Email: {}", email);
    
    // Play game
    score += 100;
    println!("Score: {}", score);
    
    score += 50;
    println!("Score: {}", score);
    
    level += 1;
    println!("Level up! Now at level: {}", level);
    
    // Data transformation using shadowing
    let score_display = format!("Score: {}", score);
    let score_display = score_display.to_uppercase();
    println!("{}", score_display);
}
```

## 🐛 Common Mistakes

### Mistake 1: Forgetting `mut`

```rust
fn main() {
    let counter = 0;
    counter += 1;  // ERROR
}
```

**Fix**: `let mut counter = 0;`

### Mistake 2: Unused Variables

```rust
fn main() {
    let x = 5;  // Warning: unused variable
}
```

**Fix**: Either use it or prefix with `_`:
```rust
let _x = 5;  // OK: explicitly unused
```

### Mistake 3: Shadowing When Mutation Intended

```rust
fn main() {
    let mut x = 5;
    let x = 10;  // Shadows, doesn't mutate!
    // Previous x (mutable) is now inaccessible
}
```

### Mistake 4: Trying to Change Type with `mut`

```rust
fn main() {
    let mut value = "text";
    value = 42;  // ERROR: can't change type
}
```

**Fix**: Use shadowing or separate variables

## 🎓 Key Takeaways

1. **Variables are immutable by default** - Rust's safety feature
2. **Use `mut` explicitly when you need to change values**
3. **Shadowing creates new variables with the same name**
4. **Shadowing allows type changes, mutation doesn't**
5. **Use snake_case for variable names**
6. **Descriptive names make code readable**

## 💪 Practice Exercises

### Exercise 1: Temperature Converter
```rust
fn main() {
    let celsius = 25;
    // Convert to Fahrenheit
    // Formula: F = C * 9/5 + 32
    // Use shadowing to transform the value
}
```

### Exercise 2: Score Tracker
Create a program that:
- Starts with score = 0
- Adds points for completing tasks
- Prints score after each update

### Exercise 3: Shadowing Practice
```rust
fn main() {
    let value = "123";
    // Shadow to parse to number
    // Shadow to multiply by 2
    // Shadow to convert back to string
    // Print result
}
```

### Exercise 4: Fix the Errors
```rust
fn main() {
    let x = 5;
    x = 6;
    
    let mut y = "hello";
    y = 10;
    
    let 2nd_value = 20;
}
```

## 🎯 Challenge

Create a simple counter program:
1. Start with a counter at 0
2. Increment it 5 times
3. Multiply final value by 10 using shadowing
4. Print the result

## 📚 What's Next?

In [Lesson 3: Data Types](./03-data-types.md), you'll learn about the different types of values Rust can work with.

---

**Great job!** 🎉 You now understand Rust's unique approach to variables. The immutable-by-default design might feel strange at first, but it's one of Rust's most powerful features for writing safe, correct programs.
