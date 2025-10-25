# Lesson 1: Hello World - Your First Java Program

## 🎯 Learning Objectives

By the end of this lesson, you'll understand:
- The structure of a Java program
- How to compile and run Java code
- What classes are and why Java uses them
- The role of the `main` method
- How the JVM (Java Virtual Machine) works

## 🤔 Why Start with Hello World?

"Hello World" is the traditional first program because:
1. **Simple**: No complex logic, focus on structure
2. **Complete**: A working program that does something
3. **Validates Setup**: Confirms Java is installed correctly
4. **Universal**: Millions of programmers started here

## 📝 The Code

Create a file named `HelloWorld.java`:

```java
public class HelloWorld {
    public static void main(String[] args) {
        System.out.println("Hello, World!");
    }
}
```

**Important**: The filename MUST match the class name: `HelloWorld.java`

Let's understand every single part.

## 🔍 Line-by-Line Breakdown

### Line 1: `public class HelloWorld {`

**`public`**: Access modifier
- Means this class can be accessed from anywhere
- Makes the class visible to the Java runtime

**`class`**: Keyword that declares a class
- Everything in Java is inside a class
- Classes are blueprints for objects
- Think: a class is like a cookie cutter, objects are the cookies

**`HelloWorld`**: The class name
- Must match the filename (case-sensitive!)
- Follows PascalCase convention (capitalize each word)
- Should be descriptive

**`{`**: Opening brace
- Marks the beginning of the class body
- Everything between `{` and `}` is part of the class

**Why do we need a class for everything?**

Java is object-oriented by design:
- Real-world things are modeled as objects
- Even simple programs need structure
- Classes provide organization and encapsulation

**Analogy**: Think of a class like a house blueprint. Even if you just want to print "Hello World", you still need a structure (house) to do it in.

### Line 2: `public static void main(String[] args) {`

This is the **most important line** in Java! Let's break it into pieces:

**`public`**: Access modifier
- The main method must be public
- Java runtime needs to access it from outside

**`static`**: Means this method belongs to the class itself
- Not to any particular object (instance)
- Can be called without creating an object
- Required for `main` because program starts before objects exist

**Think about it**: When your program starts, no objects exist yet. Java needs a starting point that doesn't require an object. That's why `main` is static.

**`void`**: Return type
- This method doesn't return any value
- `void` means "returns nothing"
- Other methods might return int, String, etc.

**`main`**: The method name
- Special name recognized by Java runtime
- Entry point of every Java application
- Like the front door to your program

**`String[] args`**: Parameter
- `String[]` is an array of strings
- `args` is the parameter name (short for "arguments")
- Allows passing command-line arguments to your program
- We'll learn more about arrays later

**`{`**: Opening brace
- Marks the beginning of the method body

**Complete signature breakdown**:
```
public  →  Accessible from anywhere
static  →  Belongs to class, not objects
void    →  Returns nothing
main    →  Entry point (special name)
String[] args  →  Command-line arguments
```

**Why so complicated for "Hello World"?**
- Java prioritizes structure and explicitness
- This structure enables powerful features as programs grow
- Same pattern scales from tiny to huge applications

### Line 3: `System.out.println("Hello, World!");`

Let's break this into parts:

**`System`**: A class from Java standard library
- Contains useful system-related functionality
- Always available (part of `java.lang` package)

**`.out`**: A static field in System class
- Represents the standard output stream
- Usually connected to your console/terminal

**`.println`**: A method of the `out` object
- **print** = output text
- **ln** = line (adds newline after printing)

**`("Hello, World!")`**: Method argument
- The text to print
- Enclosed in double quotes (String literal)

**`;`**: Semicolon
- Ends the statement
- Required after every statement in Java
- Like a period ending a sentence

**Full chain**:
```
System    →  System class
  .out    →  Standard output stream
    .println()  →  Print with newline
```

**Analogy**: It's like saying "In the System building, go to the 'out' office, and use the 'println' machine to print this message."

### Lines 4 & 5: `} }`

**First `}`**: Closes the `main` method
**Second `}`**: Closes the `HelloWorld` class

Every opening `{` must have a matching closing `}`.

## 🏗️ The Complete Structure

```java
public class HelloWorld {              // Class declaration
    public static void main(String[] args) {  // Main method
        System.out.println("Hello, World!");  // Statement
    }                                  // End method
}                                      // End class
```

**Indentation**: Notice how each level is indented
- Makes code readable
- Shows nesting structure
- Most IDEs do this automatically

## 🚀 Compiling and Running

### Step 1: Compile (Transform to Bytecode)

```bash
javac HelloWorld.java
```

**What happened?**
- `javac` = Java Compiler
- Reads `HelloWorld.java` (source code)
- Creates `HelloWorld.class` (bytecode)

Check what was created:
```bash
ls
# You'll see: HelloWorld.java and HelloWorld.class
```

**What is bytecode?**
- Intermediate representation between source code and machine code
- Not human-readable
- Not CPU-specific
- Runs on any platform with JVM

### Step 2: Run (Execute on JVM)

```bash
java HelloWorld
```

**Note**: Use the class name, NOT the filename (no `.class` extension)

**Output**:
```
Hello, World!
```

**What happened?**
1. `java` command starts the JVM
2. JVM loads `HelloWorld.class`
3. JVM finds the `main` method
4. Executes the code inside `main`
5. Prints "Hello, World!"
6. Program ends

## 🔍 Deep Dive: How Java Works

### The Process

```
Source Code (.java)
        ↓
    Java Compiler (javac)
        ↓
    Bytecode (.class)
        ↓
Java Virtual Machine (JVM)
        ↓
    Machine Code
        ↓
    Execution
```

### Why This Design?

**Problem**: Different CPUs understand different machine code
- x86 (Intel/AMD)
- ARM (phones, tablets)
- PowerPC
- etc.

**Solution**: Bytecode
- Write once, compile once
- Bytecode runs on any JVM
- JVM translates to specific machine code

**Benefits**:
1. **Platform Independence**: Same bytecode everywhere
2. **Security**: JVM can check bytecode before running
3. **Optimization**: JVM can optimize during runtime (JIT compilation)

**Trade-off**: 
- Slightly slower startup (JVM initialization)
- But: JIT makes runtime performance excellent

## 🎨 Experimenting

### Experiment 1: Change the Message

```java
public class HelloWorld {
    public static void main(String[] args) {
        System.out.println("Welcome to Java!");
    }
}
```

Compile and run:
```bash
javac HelloWorld.java
java HelloWorld
```

### Experiment 2: Multiple Print Statements

```java
public class HelloWorld {
    public static void main(String[] args) {
        System.out.println("Hello, World!");
        System.out.println("Welcome to Java!");
        System.out.println("Let's learn programming!");
    }
}
```

**Output**: Three lines, each from a separate `println`

### Experiment 3: Print Without Newline

```java
public class HelloWorld {
    public static void main(String[] args) {
        System.out.print("Hello, ");
        System.out.print("World!");
        System.out.println();  // Just newline
    }
}
```

**Output**: `Hello, World!` on one line

**Difference**:
- `println()` = print with newline
- `print()` = print without newline

### Experiment 4: Command-Line Arguments

```java
public class HelloWorld {
    public static void main(String[] args) {
        System.out.println("Hello, " + args[0] + "!");
    }
}
```

Compile and run with an argument:
```bash
javac HelloWorld.java
java HelloWorld Alice
```

**Output**: `Hello, Alice!`

**What's `args[0]`?**
- `args` is the array of command-line arguments
- `args[0]` is the first argument
- Arrays are 0-indexed (first element is 0, not 1)

## 🐛 Common Mistakes

### Mistake 1: Filename Doesn't Match Class Name

```java
// File: hello.java (wrong - should be HelloWorld.java)
public class HelloWorld {
    public static void main(String[] args) {
        System.out.println("Hello!");
    }
}
```

**Error**: 
```
error: class HelloWorld is public, should be declared in a file named HelloWorld.java
```

**Fix**: Rename file to `HelloWorld.java`

### Mistake 2: Forgetting Semicolon

```java
public class HelloWorld {
    public static void main(String[] args) {
        System.out.println("Hello, World!")  // Missing semicolon
    }
}
```

**Error**: `';' expected`

**Fix**: Add `;` at the end of the statement

### Mistake 3: Wrong Case

```java
public class HelloWorld {
    public static void main(String[] args) {
        system.out.println("Hello!");  // Wrong: lowercase 's'
    }
}
```

**Error**: `cannot find symbol: class system`

**Fix**: Use `System` (capital S)

### Mistake 4: Wrong Main Method Signature

```java
public class HelloWorld {
    public void main(String[] args) {  // Missing 'static'
        System.out.println("Hello!");
    }
}
```

**Error**: `Main method not found in class HelloWorld`

**Fix**: Add `static` keyword

### Mistake 5: Running Wrong Command

```bash
java HelloWorld.java  # Wrong: includes .java extension
java HelloWorld.class # Wrong: includes .class extension
```

**Correct**: `java HelloWorld` (just the class name)

## 🎓 Key Takeaways

1. **Filename must match public class name** (case-sensitive)
2. **Every Java program needs a main method**: `public static void main(String[] args)`
3. **Java is compiled to bytecode**, then run on JVM
4. **Semicolons end statements** - don't forget them!
5. **Java is case-sensitive** - `System` ≠ `system`
6. **Use `javac` to compile**, `java` to run
7. **Everything is inside a class** - Java's OOP design

## 💪 Practice Exercises

### Exercise 1: Personalized Greeting
Modify the program to print:
```
Hello, [Your Name]!
Welcome to Java programming.
Today is a great day to code!
```

### Exercise 2: ASCII Art
Create a program that prints:
```
  *
 ***
*****
 ***
  *
```

### Exercise 3: Multiple Classes
Create a second file `Greeting.java`:
```java
public class Greeting {
    public static void main(String[] args) {
        System.out.println("Greetings from another class!");
    }
}
```

Compile and run both programs.

### Exercise 4: Command-Line Calculator
```java
public class Calculator {
    public static void main(String[] args) {
        // Print args[0] + args[1]
        // Hint: You'll learn about this in later lessons!
    }
}
```

## 🎯 Challenge

Create a program named `AboutMe.java` that prints:
- Your name
- Your age
- Your city
- Your favorite programming language (Java, of course!)
- A fun fact about yourself

Make it visually appealing with spacing and formatting!

## 📚 What's Next?

In [Lesson 2: Variables and Data Types](./02-variables.md), you'll learn how to store and manipulate data.

---

**Congratulations!** 🎉 You've written and run your first Java program. You understand the basic structure, compilation process, and how the JVM works. You're now on your way to becoming a Java developer!
