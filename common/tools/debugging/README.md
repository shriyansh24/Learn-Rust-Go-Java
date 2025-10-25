# Debugging Guide: Rust, Go, and Java

## Overview

Comprehensive guide to debugging techniques, tools, and best practices across all three languages.

---

## General Debugging Principles

1. **Reproduce consistently**: Ensure you can trigger the bug reliably
2. **Isolate the problem**: Narrow down to smallest reproducible case
3. **Form hypotheses**: What could cause this behavior?
4. **Test hypotheses**: Add logging, breakpoints, or assertions
5. **Fix and verify**: Ensure fix doesn't break other functionality

---

## Rust Debugging

### Print Debugging

```rust
// Debug trait
#[derive(Debug)]
struct User {
    name: String,
    age: u32,
}

let user = User { name: "Alice".to_string(), age: 30 };
println!("{:?}", user);        // Debug: User { name: "Alice", age: 30 }
println!("{:#?}", user);       // Pretty: User {
                               //    name: "Alice",
                               //    age: 30,
                               // }

// dbg! macro (prints and returns value)
let sum = dbg!(x + y);         // [src/main.rs:10] x + y = 42
```

### GDB/LLDB Debugger

```bash
# Compile with debug symbols
cargo build

# Run with debugger
rust-gdb target/debug/myapp
# or
rust-lldb target/debug/myapp
```

**Common GDB Commands:**
```
break main             # Set breakpoint at main
run                    # Start program
next (n)              # Step over
step (s)              # Step into
print variable        # Print variable
backtrace (bt)        # Show call stack
continue (c)          # Continue execution
```

### VSCode Debugging

Install CodeLLDB extension, then add `.vscode/launch.json`:

```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "type": "lldb",
      "request": "launch",
      "name": "Debug",
      "cargo": {
        "args": ["build", "--bin=myapp"]
      },
      "args": [],
      "cwd": "${workspaceFolder}"
    }
  ]
}
```

### Common Issues

**1. Borrow Checker Errors**
```rust
// Error: cannot borrow as mutable
let s = String::from("hello");
let r1 = &s;
let r2 = &mut s;  // ❌

// Fix: No simultaneous immutable and mutable borrows
let mut s = String::from("hello");
{
    let r1 = &s;
    println!("{}", r1);
}  // r1 goes out of scope
let r2 = &mut s;  // ✅
```

**2. Lifetime Issues**
```rust
// Error: lifetime mismatch
fn longest(x: &str, y: &str) -> &str {  // ❌
    if x.len() > y.len() { x } else { y }
}

// Fix: Explicit lifetime annotation
fn longest<'a>(x: &'a str, y: &'a str) -> &'a str {  // ✅
    if x.len() > y.len() { x } else { y }
}
```

---

## Go Debugging

### Print Debugging

```go
import "fmt"

// Basic print
fmt.Println("Value:", x)

// Printf with format
fmt.Printf("x=%d, y=%s\n", x, y)

// Pretty print structs
fmt.Printf("%+v\n", user)  // {Name:Alice Age:30}
fmt.Printf("%#v\n", user)  // main.User{Name:"Alice", Age:30}
```

### Delve Debugger

```bash
# Install
go install github.com/go-delve/delve/cmd/dlv@latest

# Debug
dlv debug main.go

# Or attach to running process
dlv attach <pid>
```

**Delve Commands:**
```
break main.main       # Set breakpoint
continue (c)          # Continue
next (n)              # Step over
step (s)              # Step into
print x               # Print variable
locals                # Show local variables
goroutines            # List goroutines
goroutine 1           # Switch to goroutine 1
```

### VSCode Debugging

`.vscode/launch.json`:
```json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Launch Package",
      "type": "go",
      "request": "launch",
      "mode": "debug",
      "program": "${workspaceFolder}"
    }
  ]
}
```

### Common Issues

**1. Nil Pointer Dereference**
```go
var ptr *int
*ptr = 42  // panic: runtime error: invalid memory address

// Fix: Check before dereferencing
if ptr != nil {
    *ptr = 42
}
```

**2. Goroutine Leaks**
```go
// Problem: goroutine never exits
func leak() {
    ch := make(chan int)
    go func() {
        val := <-ch  // Blocks forever
        fmt.Println(val)
    }()
    // ch never receives, goroutine leaks
}

// Fix: Use context for cancellation
func noLeak(ctx context.Context) {
    ch := make(chan int)
    go func() {
        select {
        case val := <-ch:
            fmt.Println(val)
        case <-ctx.Done():
            return  // Exit on cancellation
        }
    }()
}
```

**3. Race Conditions**
```go
// Detect races
go run -race main.go

// Or during tests
go test -race
```

---

## Java Debugging

### Print Debugging

```java
// System.out
System.out.println("Value: " + x);

// String formatting
System.out.printf("x=%d, y=%s\n", x, y);

// Logging (preferred)
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

private static final Logger log = LoggerFactory.getLogger(MyClass.class);

log.debug("Debug message");
log.info("Info message");
log.warn("Warning message");
log.error("Error message", exception);
```

### JDB (Java Debugger)

```bash
# Compile with debug info
javac -g MyClass.java

# Run with debugger
jdb MyClass
```

**JDB Commands:**
```
stop at MyClass:10    # Breakpoint at line 10
stop in MyClass.main  # Breakpoint in method
run                   # Start program
step                  # Step into
next                  # Step over
cont                  # Continue
print variable        # Print variable
where                 # Show stack trace
```

### IDE Debugging (IntelliJ IDEA / Eclipse)

**IntelliJ:**
- Click line number to set breakpoint
- Right-click → Debug
- Use debugger window for step/watch/evaluate

**Eclipse:**
- Double-click line number for breakpoint
- Debug As → Java Application
- Use Debug perspective

### Common Issues

**1. NullPointerException**
```java
// Problem
String s = null;
s.length();  // NullPointerException

// Fix 1: Null check
if (s != null) {
    s.length();
}

// Fix 2: Optional (Java 8+)
Optional<String> opt = Optional.ofNullable(s);
opt.ifPresent(str -> System.out.println(str.length()));
```

**2. ConcurrentModificationException**
```java
// Problem: Modifying collection during iteration
List<Integer> list = new ArrayList<>(Arrays.asList(1, 2, 3));
for (Integer i : list) {
    list.remove(i);  // ❌ ConcurrentModificationException
}

// Fix: Use Iterator
Iterator<Integer> it = list.iterator();
while (it.hasNext()) {
    it.next();
    it.remove();  // ✅
}
```

**3. Memory Leaks**
```java
// Problem: Static collection grows unbounded
public class Cache {
    private static Map<String, Object> cache = new HashMap<>();

    public static void put(String key, Object value) {
        cache.put(key, value);  // Never removed
    }
}

// Fix: Use weak references or cache eviction
import java.util.WeakHashMap;

private static Map<String, Object> cache = new WeakHashMap<>();
// Or use Caffeine/Guava cache with size/time limits
```

---

## Advanced Debugging Techniques

### Core Dumps (Post-Mortem Debugging)

**Rust/Go:**
```bash
# Enable core dumps
ulimit -c unlimited

# After crash, analyze
gdb ./myapp core
```

**Java:**
```bash
# Heap dump on OutOfMemoryError
java -XX:+HeapDumpOnOutOfMemoryError -XX:HeapDumpPath=/tmp/heap.hprof MyApp

# Analyze with Eclipse MAT or VisualVM
```

### Remote Debugging

**Go:**
```bash
# Start with Delve headless
dlv debug --headless --listen=:2345 --api-version=2

# Connect from another machine
dlv connect <host>:2345
```

**Java:**
```bash
# Start with debug port
java -agentlib:jdwp=transport=dt_socket,server=y,suspend=n,address=5005 MyApp

# Connect from IDE (IntelliJ/Eclipse) to localhost:5005
```

### Conditional Breakpoints

**GDB/LLDB:**
```
break main.rs:42 if x > 100
```

**IntelliJ:**
Right-click breakpoint → Condition: `x > 100`

### Logging Best Practices

1. **Use appropriate levels**: DEBUG < INFO < WARN < ERROR
2. **Include context**: Thread ID, request ID, user ID
3. **Structured logging**: JSON format for parsing
4. **Avoid logging sensitive data**: Passwords, credit cards
5. **Use correlation IDs**: Track requests across services

---

## Debugging Tools Comparison

| Tool | Rust | Go | Java |
|------|------|-----|------|
| **Native Debugger** | GDB, LLDB | Delve | JDB |
| **IDE** | VSCode, CLion | GoLand, VSCode | IntelliJ, Eclipse |
| **Print Debugging** | `println!`, `dbg!` | `fmt.Println` | `System.out`, Logging |
| **Memory Profiler** | Valgrind, heaptrack | pprof | VisualVM, YourKit |
| **Core Dumps** | Yes | Yes | Heap dumps |
| **Remote Debugging** | GDB server | Delve headless | JDWP |

---

## Resources

### Rust
- [Rust Debugging Guide](https://doc.rust-lang.org/book/appendix-04-useful-development-tools.html)
- [rust-gdb Tutorial](https://sourceware.org/gdb/current/onlinedocs/gdb/)

### Go
- [Delve Documentation](https://github.com/go-delve/delve/tree/master/Documentation)
- [Go Debugging with Delve](https://go.dev/doc/articles/wiki/)

### Java
- [Java Debugging Tutorial](https://www.baeldung.com/java-debugging)
- [IntelliJ Debugging](https://www.jetbrains.com/help/idea/debugging-code.html)

---

**Next:** [Profiling](../profiling/)
