# Type Systems: From Static to Generic

## Table of Contents
1. [Introduction to Type Systems](#introduction-to-type-systems)
2. [Static vs Dynamic Typing](#static-vs-dynamic-typing)
3. [Type Inference](#type-inference)
4. [Rust's Type System](#rusts-type-system)
5. [Go's Type System](#gos-type-system)
6. [Java's Type System](#javas-type-system)
7. [Generics and Parametric Polymorphism](#generics-and-parametric-polymorphism)
8. [Type Safety and Soundness](#type-safety-and-soundness)
9. [Performance Implications](#performance-implications)
10. [Best Practices](#best-practices)

---

## Introduction to Type Systems

A **type system** is a set of rules that assigns a property called "type" to program constructs (variables, functions, expressions) and enforces constraints on how they can be used.

### Why Type Systems Matter

```
Without Types:
x = 5
x = "hello"  ← What operations can we perform on x?
x + 10       ← What does this mean?

With Types:
int x = 5;
x + 10       ✅ Clear: integer addition = 15
x + "hello"  ❌ Error: can't add string to int
```

### Benefits of Strong Type Systems

1. **Early Error Detection**: Catch bugs at compile time
2. **Documentation**: Types serve as executable documentation
3. **IDE Support**: Better autocomplete and refactoring
4. **Performance**: Compiler optimizations based on type info
5. **Maintenance**: Easier to understand and modify code

---

## Static vs Dynamic Typing

### Static Typing

**Definition:** Types checked at compile time

**Languages:** Rust, Go, Java, C++, TypeScript

**Example (Rust):**
```rust
let x: i32 = 5;
// x = "hello";  ❌ Compile error: expected i32, found &str
```

**Pros:**
- Catch type errors before runtime
- Better performance (no runtime type checks)
- Superior tooling and IDE support
- Self-documenting code

**Cons:**
- More verbose
- Less flexible for prototyping
- Steeper learning curve

### Dynamic Typing

**Definition:** Types checked at runtime

**Languages:** Python, JavaScript, Ruby, PHP

**Example (Python):**
```python
x = 5          # x is int
x = "hello"    # x is now str (allowed)
x + 10         # ❌ Runtime error: can't add str and int
```

**Pros:**
- Rapid prototyping
- More flexible
- Less boilerplate
- Easier for beginners

**Cons:**
- Type errors discovered at runtime
- Slower performance (runtime checks)
- Requires more testing
- Refactoring is riskier

### Comparison Table

| Aspect | Static Typing | Dynamic Typing |
|--------|---------------|----------------|
| **Error Detection** | Compile time | Runtime |
| **Performance** | Faster (optimized) | Slower (checks at runtime) |
| **Flexibility** | Lower | Higher |
| **Tooling** | Excellent | Good |
| **Verbosity** | Higher | Lower |
| **Learning Curve** | Steeper | Gentler |

---

## Type Inference

**Type inference** allows the compiler to deduce types automatically without explicit annotations.

### Rust Type Inference

```rust
// Explicit type annotation
let x: i32 = 5;

// Type inference
let x = 5;        // Compiler infers i32
let y = 5.0;      // Compiler infers f64
let s = "hello";  // Compiler infers &str

// Inference from usage
let mut v = Vec::new();  // Type unknown
v.push(1);               // Now compiler knows: Vec<i32>

// Multiple possibilities - requires annotation
let guess: u32 = "42".parse().expect("Not a number!");
```

**How it works:**
```rust
fn add(a: i32, b: i32) -> i32 {
    a + b
}

let result = add(5, 10);
// Compiler flow:
// 1. `add` returns i32
// 2. `result` stores return value
// 3. Therefore, `result` is i32
```

### Go Type Inference

```go
// Explicit type
var x int = 5

// Type inference with :=
x := 5           // Inferred as int
y := 5.0         // Inferred as float64
s := "hello"     // Inferred as string

// Multiple assignment
a, b := 10, "test"  // a: int, b: string

// Inference in function return
func getValue() int {
    return 42
}
result := getValue()  // result is int
```

### Java Type Inference

```java
// Before Java 10 - explicit type
List<String> list = new ArrayList<String>();

// Java 7+ - diamond operator
List<String> list = new ArrayList<>();

// Java 10+ - var keyword (local variables only)
var list = new ArrayList<String>();  // Inferred as ArrayList<String>
var x = 5;                           // Inferred as int
var s = "hello";                     // Inferred as String

// Java 11+ - var in lambda parameters
BiFunction<Integer, Integer, Integer> add = (var a, var b) -> a + b;
```

**Limitations:**
```java
// ❌ Can't use var without initializer
// var x;

// ❌ Can't use var for fields
class Example {
    // var field = 5;  // Not allowed
}

// ❌ Can't use var for method parameters/return types
// public var getValue() { return 42; }
```

---

## Rust's Type System

### Core Principles

1. **Type Safety**: No null pointer dereferences (use `Option<T>`)
2. **Zero-Cost Abstractions**: High-level abstractions compile to efficient code
3. **Ownership Integration**: Types enforce memory safety rules

### Primitive Types

```rust
// Integer types
let a: i8 = 127;          // -128 to 127
let b: i16 = 32767;       // -32,768 to 32,767
let c: i32 = 2147483647;  // Default integer type
let d: i64 = 9223372036854775807;
let e: i128 = 1;          // 128-bit signed

let ua: u8 = 255;         // 0 to 255
let ub: u16 = 65535;
let uc: u32 = 4294967295;
let ud: u64 = 1;
let ue: u128 = 1;

// Platform-dependent
let ptr_sized: isize = 100;  // Pointer-sized signed
let ptr_usized: usize = 100; // Pointer-sized unsigned

// Floating point
let x: f32 = 3.14;
let y: f64 = 2.718;  // Default float type

// Boolean
let t: bool = true;
let f: bool = false;

// Character (4 bytes - Unicode scalar)
let c: char = 'z';
let z: char = '😎';
```

### Compound Types

```rust
// Tuple
let tup: (i32, f64, u8) = (500, 6.4, 1);
let (x, y, z) = tup;  // Destructuring
let first = tup.0;

// Array (fixed size)
let a: [i32; 5] = [1, 2, 3, 4, 5];
let first = a[0];

// Slice (view into array)
let slice: &[i32] = &a[1..3];  // [2, 3]

// String types
let s1: &str = "hello";           // String slice (immutable)
let s2: String = String::from("hello");  // Owned string

// Vector (dynamic array)
let v: Vec<i32> = vec![1, 2, 3];
```

### Option and Result

```rust
// Option<T> - handles absence of value (no null!)
enum Option<T> {
    Some(T),
    None,
}

fn divide(a: f64, b: f64) -> Option<f64> {
    if b == 0.0 {
        None
    } else {
        Some(a / b)
    }
}

match divide(10.0, 2.0) {
    Some(result) => println!("Result: {}", result),
    None => println!("Cannot divide by zero"),
}

// Result<T, E> - handles errors
enum Result<T, E> {
    Ok(T),
    Err(E),
}

use std::fs::File;

fn open_file(path: &str) -> Result<File, std::io::Error> {
    File::open(path)
}

match open_file("data.txt") {
    Ok(file) => println!("File opened"),
    Err(e) => println!("Error: {}", e),
}
```

### Traits: Rust's Interfaces

```rust
// Define a trait
trait Summary {
    fn summarize(&self) -> String;

    // Default implementation
    fn default_summary(&self) -> String {
        String::from("(Read more...)")
    }
}

// Implement trait for a type
struct Article {
    title: String,
    content: String,
}

impl Summary for Article {
    fn summarize(&self) -> String {
        format!("{}: {}", self.title, &self.content[..50])
    }
}

// Trait bounds
fn notify<T: Summary>(item: &T) {
    println!("Breaking news! {}", item.summarize());
}

// Multiple trait bounds
fn process<T: Summary + Display>(item: &T) {
    // ...
}

// Where clause (more readable)
fn complex<T, U>(t: &T, u: &U) -> String
where
    T: Summary + Clone,
    U: Display + Debug,
{
    // ...
}
```

### Phantom Types

```rust
use std::marker::PhantomData;

struct Meter;
struct Foot;

struct Distance<Unit> {
    value: f64,
    _marker: PhantomData<Unit>,
}

impl<Unit> Distance<Unit> {
    fn new(value: f64) -> Self {
        Distance {
            value,
            _marker: PhantomData,
        }
    }
}

let meters = Distance::<Meter>::new(100.0);
let feet = Distance::<Foot>::new(328.0);

// ❌ Type mismatch:
// let sum = meters.value + feet.value;  // Different types!
```

---

## Go's Type System

### Design Philosophy

- **Simplicity**: Minimal type system complexity
- **Clarity**: Explicit over implicit
- **Composition**: Interfaces over inheritance

### Basic Types

```go
// Boolean
var b bool = true

// Integer types
var i8 int8 = 127             // -128 to 127
var i16 int16 = 32767
var i32 int32 = 2147483647    // Also known as 'rune'
var i64 int64 = 9223372036854775807

var u8 uint8 = 255            // Also known as 'byte'
var u16 uint16 = 65535
var u32 uint32 = 4294967295
var u64 uint64 = 1

// Platform-dependent
var i int = 100               // 32 or 64 bit
var u uint = 100
var ptr uintptr = 0x1234      // Pointer-sized unsigned

// Floating point
var f32 float32 = 3.14
var f64 float64 = 2.718       // Default

// Complex numbers
var c64 complex64 = 1 + 2i
var c128 complex128 = 3 + 4i

// String
var s string = "hello"
```

### Composite Types

```go
// Array (fixed size)
var a [5]int = [5]int{1, 2, 3, 4, 5}
b := [...]int{1, 2, 3}  // Size inferred

// Slice (dynamic view)
var s []int = []int{1, 2, 3}
s = append(s, 4)

// Map
var m map[string]int = make(map[string]int)
m["age"] = 30

// Struct
type Person struct {
    Name string
    Age  int
}

p := Person{Name: "Alice", Age: 30}
```

### Interfaces

```go
// Interface definition
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// Empty interface (any type)
var any interface{}
any = 42
any = "hello"
any = true

// Type assertion
var i interface{} = "hello"
s := i.(string)       // Type assertion
fmt.Println(s)        // hello

s, ok := i.(string)   // Safe type assertion
if ok {
    fmt.Println(s)
}

// Type switch
func describe(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("Integer: %d\n", v)
    case string:
        fmt.Printf("String: %s\n", v)
    default:
        fmt.Printf("Unknown type\n")
    }
}
```

### Structural Typing

```go
// Interface implementation is implicit
type Animal interface {
    Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
    return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
    return "Meow!"
}

// Both Dog and Cat implement Animal automatically!
func makeSound(a Animal) {
    fmt.Println(a.Speak())
}

makeSound(Dog{})  // Woof!
makeSound(Cat{})  // Meow!
```

### Generics (Go 1.18+)

```go
// Generic function
func Min[T constraints.Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}

fmt.Println(Min(3, 5))        // 3
fmt.Println(Min(3.5, 2.1))    // 2.1
fmt.Println(Min("b", "a"))    // a

// Generic type
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item
}

// Usage
intStack := Stack[int]{}
intStack.Push(1)
intStack.Push(2)
fmt.Println(intStack.Pop())  // 2

stringStack := Stack[string]{}
stringStack.Push("hello")
```

---

## Java's Type System

### Class-Based Object-Oriented

Java's type system is built around classes and interfaces with nominal typing.

### Primitive vs Reference Types

```java
// Primitive types (value types)
byte b = 127;              // 8-bit
short s = 32767;           // 16-bit
int i = 2147483647;        // 32-bit (default)
long l = 9223372036854775807L;  // 64-bit

float f = 3.14f;           // 32-bit
double d = 2.718;          // 64-bit (default)

boolean bool = true;
char c = 'A';              // 16-bit Unicode

// Reference types (heap-allocated)
Integer boxedInt = 42;     // Wrapper class
String str = "hello";
int[] array = {1, 2, 3};
List<Integer> list = new ArrayList<>();
```

**Autoboxing/Unboxing:**
```java
// Autoboxing: primitive → wrapper
Integer wrapped = 42;      // int → Integer

// Unboxing: wrapper → primitive
int primitive = wrapped;   // Integer → int

// Performance consideration
List<Integer> list = new ArrayList<>();
for (int i = 0; i < 1000; i++) {
    list.add(i);  // Autoboxing on every iteration (overhead)
}
```

### Generics

```java
// Generic class
public class Box<T> {
    private T value;

    public void set(T value) {
        this.value = value;
    }

    public T get() {
        return value;
    }
}

Box<Integer> intBox = new Box<>();
intBox.set(42);
Integer value = intBox.get();

Box<String> strBox = new Box<>();
strBox.set("hello");
String str = strBox.get();
```

**Bounded Type Parameters:**
```java
// Upper bound
public class NumberBox<T extends Number> {
    private T value;

    public double doubleValue() {
        return value.doubleValue();  // Can call Number methods
    }
}

NumberBox<Integer> intBox = new NumberBox<>();    // ✅
NumberBox<Double> doubleBox = new NumberBox<>();  // ✅
// NumberBox<String> strBox = new NumberBox<>();  // ❌ String is not a Number

// Multiple bounds
public class MultiBox<T extends Number & Comparable<T>> {
    // T must be both a Number and Comparable
}
```

**Wildcards:**
```java
// Unbounded wildcard
public void printList(List<?> list) {
    for (Object obj : list) {
        System.out.println(obj);
    }
}

// Upper bounded wildcard (covariance)
public double sum(List<? extends Number> numbers) {
    double sum = 0;
    for (Number num : numbers) {
        sum += num.doubleValue();
    }
    return sum;
}

List<Integer> ints = Arrays.asList(1, 2, 3);
List<Double> doubles = Arrays.asList(1.1, 2.2, 3.3);
sum(ints);     // ✅
sum(doubles);  // ✅

// Lower bounded wildcard (contravariance)
public void addIntegers(List<? super Integer> list) {
    list.add(42);      // ✅ Can add Integer
    list.add(100);     // ✅
    // Integer x = list.get(0);  // ❌ Can't guarantee return type
}

List<Number> numbers = new ArrayList<>();
addIntegers(numbers);  // ✅
```

### Type Erasure

**Problem:** Generic type information is erased at runtime

```java
// At compile time:
List<String> strings = new ArrayList<String>();

// At runtime (after type erasure):
List strings = new ArrayList();  // Raw type

// Consequence: Can't check generic types at runtime
if (list instanceof List<String>) {  // ❌ Compiler error
    // ...
}

if (list instanceof List) {  // ✅ OK (raw type)
    // ...
}

// Can't create generic arrays
// T[] array = new T[10];  // ❌ Won't compile
```

### Interfaces and Abstract Classes

```java
// Interface (contract)
public interface Drawable {
    void draw();

    default void display() {  // Java 8+ default methods
        System.out.println("Displaying...");
        draw();
    }

    static void info() {  // Java 8+ static methods
        System.out.println("Drawable interface");
    }
}

// Abstract class (partial implementation)
public abstract class Shape {
    private String color;

    public Shape(String color) {
        this.color = color;
    }

    public abstract double area();  // Abstract method

    public String getColor() {      // Concrete method
        return color;
    }
}

// Concrete class
public class Circle extends Shape implements Drawable {
    private double radius;

    public Circle(String color, double radius) {
        super(color);
        this.radius = radius;
    }

    @Override
    public double area() {
        return Math.PI * radius * radius;
    }

    @Override
    public void draw() {
        System.out.println("Drawing circle");
    }
}
```

### Sealed Classes (Java 17+)

```java
// Only specified classes can extend
public sealed class Shape
    permits Circle, Rectangle, Triangle {
}

public final class Circle extends Shape {
    // Circle cannot be extended
}

public non-sealed class Rectangle extends Shape {
    // Rectangle can be extended by anyone
}

public sealed class Triangle extends Shape
    permits EquilateralTriangle {
    // Triangle can only be extended by EquilateralTriangle
}
```

---

## Generics and Parametric Polymorphism

### Comparison Across Languages

| Feature | Rust | Go | Java |
|---------|------|-----|------|
| **Syntax** | `<T>` | `[T any]` | `<T>` |
| **Constraints** | Trait bounds | Interface constraints | Bounded types |
| **Monomorphization** | Yes (compile-time) | Partial | No (type erasure) |
| **Runtime Cost** | Zero | Low | Low (boxing overhead) |

### Monomorphization vs Type Erasure

**Rust (Monomorphization):**
```rust
fn print<T: Display>(value: T) {
    println!("{}", value);
}

print(42);      // Generates: print_i32(42)
print("hello"); // Generates: print_str("hello")

// Pros: Zero runtime cost
// Cons: Larger binary size (code duplication)
```

**Java (Type Erasure):**
```java
public <T> void print(T value) {
    System.out.println(value);
}

print(42);      // Runtime: print(Object)
print("hello"); // Runtime: print(Object)

// Pros: Smaller binary size
// Cons: Can't access type info at runtime, some overhead
```

---

## Type Safety and Soundness

### Null Safety

**Rust:** No null! Use `Option<T>`
```rust
let x: Option<i32> = Some(5);
let y: Option<i32> = None;

match x {
    Some(val) => println!("Value: {}", val),
    None => println!("No value"),
}
```

**Go:** Has nil, but explicit pointer types
```go
var ptr *int = nil

if ptr != nil {
    fmt.Println(*ptr)
}
```

**Java:** Has null (billion-dollar mistake)
```java
String s = null;
// s.length();  // NullPointerException at runtime!

// Java 8+ Optional
Optional<String> opt = Optional.ofNullable(s);
opt.ifPresent(str -> System.out.println(str.length()));
```

### Type Safety Guarantees

**Rust:**
- No null pointer dereferences (compile-time)
- No data races (compile-time)
- No use-after-free (compile-time)

**Go:**
- Nil checks at runtime
- Race detector (runtime, opt-in)
- Garbage collected (prevents use-after-free)

**Java:**
- NullPointerException at runtime
- Thread-safe classes available
- Garbage collected

---

## Performance Implications

### Type System Overhead

| Language | Compile Time | Runtime Overhead | Binary Size |
|----------|--------------|------------------|-------------|
| **Rust** | Slow (monomorphization) | Zero | Large |
| **Go** | Fast | Minimal | Small-Medium |
| **Java** | Medium | GC + boxing | Medium (JIT helps) |

### Example: Generic Sorting

**Rust:**
```rust
// Compile-time specialization for each type
fn sort<T: Ord>(arr: &mut [T]) {
    arr.sort();
}

// Generated code:
// sort_i32(&mut [1, 2, 3])
// sort_String(&mut ["a", "b", "c"])
```

**Java:**
```java
// Single implementation with Object
public static <T extends Comparable<T>> void sort(List<T> list) {
    Collections.sort(list);
}

// Runtime: All T treated as Object, some boxing overhead
```

---

## Best Practices

### Rust

1. **Embrace the type system**
   ```rust
   // Good: Explicit, type-safe
   fn process(id: UserId) { }

   // Bad: Primitive obsession
   fn process(id: i32) { }
   ```

2. **Use newtype pattern**
   ```rust
   struct UserId(i32);
   struct OrderId(i32);
   // Can't accidentally mix up!
   ```

3. **Prefer `Option`/`Result` over panics**
   ```rust
   fn divide(a: f64, b: f64) -> Result<f64, String> {
       if b == 0.0 {
           Err("Division by zero".to_string())
       } else {
           Ok(a / b)
       }
   }
   ```

### Go

1. **Accept interfaces, return structs**
   ```go
   // Good
   func Process(r io.Reader) *Result {
       // ...
   }

   // Bad
   func Process(r *os.File) io.Reader {
       // ...
   }
   ```

2. **Keep interfaces small**
   ```go
   type Reader interface {
       Read(p []byte) (n int, err error)
   }
   ```

3. **Use type assertions carefully**
   ```go
   if val, ok := i.(string); ok {
       // Use val
   }
   ```

### Java

1. **Prefer generics over raw types**
   ```java
   // Good
   List<String> list = new ArrayList<>();

   // Bad
   List list = new ArrayList();  // Raw type
   ```

2. **Use bounded wildcards for API flexibility**
   ```java
   // PECS: Producer Extends, Consumer Super
   public void copy(List<? extends T> src, List<? super T> dest) {
       // ...
   }
   ```

3. **Avoid primitive arrays for generic collections**
   ```java
   // Prefer: List<Integer>
   // Over:   int[]
   ```

---

## Summary

| Aspect | Rust | Go | Java |
|--------|------|-----|------|
| **Typing** | Static, inferred | Static, inferred | Static, explicit |
| **Generics** | Full, monomorphized | Limited, runtime | Full, type erasure |
| **Null Safety** | No null (`Option`) | Nil pointers | Null (Optional in 8+) |
| **Type Inference** | Extensive | Moderate | Limited (`var`) |
| **Complexity** | High | Low | Medium |

---

## Further Reading

**Rust:**
- [The Rust Book: Types](https://doc.rust-lang.org/book/ch03-02-data-types.html)
- [Rust by Example: Generics](https://doc.rust-lang.org/rust-by-example/generics.html)

**Go:**
- [Effective Go: Interfaces](https://go.dev/doc/effective_go#interfaces)
- [Go Generics Tutorial](https://go.dev/doc/tutorial/generics)

**Java:**
- [Java Generics Tutorial](https://docs.oracle.com/javase/tutorial/java/generics/)
- [Effective Java: Item 26-33 (Generics)](https://www.oreilly.com/library/view/effective-java/9780134686097/)

---

**Next:** [Compilation vs Interpretation](./compilation-vs-interpretation.md)
