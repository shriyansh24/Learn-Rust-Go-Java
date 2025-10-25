# Memory Management: A Deep Dive Across Rust, Go, and Java

## Table of Contents
1. [Introduction](#introduction)
2. [The Stack vs The Heap](#the-stack-vs-the-heap)
3. [Memory Management Strategies](#memory-management-strategies)
4. [Rust: Ownership & Borrowing](#rust-ownership--borrowing)
5. [Go: Garbage Collection](#go-garbage-collection)
6. [Java: Managed Memory with GC](#java-managed-memory-with-gc)
7. [Performance Implications](#performance-implications)
8. [Memory Visualization](#memory-visualization)
9. [Best Practices](#best-practices)

---

## Introduction

Memory management is one of the most critical aspects of systems programming. Understanding how memory works is essential for:
- Writing performant code
- Avoiding memory leaks
- Preventing security vulnerabilities
- Making informed architectural decisions

This guide explores how Rust, Go, and Java handle memory differently, and why each approach matters.

---

## The Stack vs The Heap

Before diving into language-specific implementations, let's understand the fundamental memory regions:

### The Stack

**What is it?**
The stack is a region of memory that stores:
- Local variables
- Function parameters
- Return addresses
- CPU register states

**Key Characteristics:**
- **LIFO (Last In, First Out)**: Data is pushed and popped in strict order
- **Fixed size**: Determined at compile time
- **Fast allocation**: Just move the stack pointer
- **Automatic cleanup**: Deallocated when function returns
- **Limited size**: Typically 1-8 MB (can cause stack overflow)

**Visual Representation:**
```
┌─────────────────┐  ← Stack grows downward (lower addresses)
│  Local Var 3    │
├─────────────────┤
│  Local Var 2    │
├─────────────────┤
│  Local Var 1    │  ← Stack Pointer (SP)
├─────────────────┤
│  Return Address │
├─────────────────┤
│  Previous Frame │
└─────────────────┘  ← Stack base (higher addresses)
```

### The Heap

**What is it?**
The heap is a region of memory for dynamic allocation:
- Objects with unknown size at compile time
- Data that outlives function scope
- Large data structures

**Key Characteristics:**
- **Random access**: Allocate/deallocate in any order
- **Dynamic size**: Can grow as needed (up to system limits)
- **Slower allocation**: Requires memory management algorithms
- **Manual/automatic cleanup**: Depends on language
- **Larger size**: Limited by available RAM

**Visual Representation:**
```
┌─────────────────┐  ← Heap grows upward (higher addresses)
│      Free       │
├─────────────────┤
│   Object C      │
├─────────────────┤
│      Free       │
├─────────────────┤
│   Object B      │
├─────────────────┤
│   Object A      │
└─────────────────┘  ← Heap base (lower addresses)
```

### Why This Distinction Matters

```
Stack Allocation:
┌─────────────────────┐
│ int x = 42;         │  ← Allocated on stack
│                     │  ← Automatically freed when out of scope
└─────────────────────┘

Heap Allocation:
┌─────────────────────┐
│ Box::new(42)        │  ← Creates pointer on stack
│         │           │
│         ▼           │
│    [42] on heap     │  ← Needs explicit management
└─────────────────────┘
```

---

## Memory Management Strategies

### 1. Manual Memory Management (C/C++)

**How it works:**
- Programmer explicitly allocates (`malloc`/`new`)
- Programmer explicitly deallocates (`free`/`delete`)

**Pros:**
- Full control
- Zero overhead
- Predictable performance

**Cons:**
- Memory leaks (forgot to free)
- Double-free errors
- Use-after-free bugs
- Dangling pointers

### 2. Garbage Collection (Go, Java)

**How it works:**
- Runtime automatically tracks object references
- Periodically identifies unreachable objects
- Reclaims memory automatically

**Pros:**
- No manual deallocation
- Prevents many memory bugs
- Easier to write correct code

**Cons:**
- Non-deterministic pauses (GC cycles)
- Memory overhead (tracking metadata)
- Potential latency spikes

### 3. Ownership System (Rust)

**How it works:**
- Compiler enforces ownership rules at compile time
- Memory freed when owner goes out of scope
- Borrowing rules prevent data races

**Pros:**
- Memory safety without GC
- Zero runtime overhead
- Predictable performance

**Cons:**
- Steeper learning curve
- Borrow checker can be restrictive
- More upfront thinking required

---

## Rust: Ownership & Borrowing

### Core Ownership Rules

1. Each value has exactly one owner
2. When the owner goes out of scope, the value is dropped
3. Values can be moved or borrowed, but not both simultaneously

### Example: Ownership Transfer

```rust
fn main() {
    let s1 = String::from("hello");  // s1 owns the String
    let s2 = s1;                     // Ownership moved to s2

    // println!("{}", s1);           // ❌ Error! s1 no longer valid
    println!("{}", s2);              // ✅ s2 is the owner
}
```

**Memory Layout:**
```
Before move:
Stack:                  Heap:
┌──────────┐           ┌───────────┐
│ s1       │ ────────► │  "hello"  │
│ ptr: 0x..│           └───────────┘
│ len: 5   │
│ cap: 5   │
└──────────┘

After move (s2 = s1):
Stack:                  Heap:
┌──────────┐           ┌───────────┐
│ s1 ❌    │           │  "hello"  │
└──────────┘           └───────────┘
┌──────────┐              ▲
│ s2       │ ─────────────┘
│ ptr: 0x..│
│ len: 5   │
│ cap: 5   │
└──────────┘
```

### Borrowing: References

```rust
fn main() {
    let s1 = String::from("hello");

    let len = calculate_length(&s1);  // Borrow s1 (immutable reference)

    println!("Length of '{}' is {}", s1, len);  // ✅ Still valid
}

fn calculate_length(s: &String) -> usize {
    s.len()
}  // s goes out of scope, but doesn't own data, so nothing is dropped
```

**Borrowing Rules:**
- Any number of immutable references (`&T`)
- **OR** exactly one mutable reference (`&mut T`)
- References must always be valid (no dangling references)

### Lifetimes: Compiler's Safety Net

```rust
fn longest<'a>(x: &'a str, y: &'a str) -> &'a str {
    if x.len() > y.len() {
        x
    } else {
        y
    }
}

fn main() {
    let string1 = String::from("long string");
    let string2 = String::from("short");

    let result = longest(&string1, &string2);
    println!("Longest: {}", result);
}
```

**Lifetime Annotation `'a`:**
- Tells compiler that returned reference lives as long as the shorter of the input references
- Prevents dangling references at compile time

### Under the Hood: Drop Trait

```rust
struct MyStruct {
    data: Vec<i32>,
}

impl Drop for MyStruct {
    fn drop(&mut self) {
        println!("Dropping MyStruct!");
        // Cleanup happens automatically, but you can add custom logic
    }
}

fn main() {
    {
        let _s = MyStruct { data: vec![1, 2, 3] };
        // Use _s...
    }  // ← `drop` called here automatically

    println!("After scope");
}
```

---

## Go: Garbage Collection

### Concurrent Mark-Sweep Collector

Go uses a **tricolor mark-and-sweep** garbage collector that runs concurrently with application code.

### How It Works

1. **Allocation:** Objects allocated on heap
2. **Mark Phase:** GC identifies reachable objects
3. **Sweep Phase:** Unreachable objects are freed

**Tricolor Marking:**
```
White: Not yet scanned (potential garbage)
Gray:  Scanned, but references not yet scanned
Black: Scanned, all references scanned (definitely alive)
```

### Example: Escape Analysis

Go's compiler performs **escape analysis** to determine if a variable can be stack-allocated:

```go
package main

func stackAllocation() int {
    x := 42  // Allocated on stack (doesn't escape)
    return x
}

func heapAllocation() *int {
    x := 42  // Allocated on heap (escapes via return)
    return &x
}

func main() {
    a := stackAllocation()      // Stack allocation
    b := heapAllocation()       // Heap allocation
    println(a, *b)
}
```

**Check escape analysis:**
```bash
go build -gcflags="-m" main.go
# Output:
# ./main.go:8:2: moved to heap: x
```

### GC Tuning with GOGC

Control GC frequency with `GOGC` environment variable:

```bash
# Default: GOGC=100 (GC when heap doubles)
GOGC=50  ./myapp  # More frequent GC, less memory
GOGC=200 ./myapp  # Less frequent GC, more memory
GOGC=off ./myapp  # Disable GC (not recommended)
```

### Memory Profiling

```go
package main

import (
    "runtime"
    "runtime/pprof"
    "os"
)

func main() {
    f, _ := os.Create("mem.prof")
    defer f.Close()

    // Allocate some memory
    data := make([]byte, 10<<20)  // 10 MB
    _ = data

    runtime.GC()  // Force GC
    pprof.WriteHeapProfile(f)
}
```

**Analyze with:**
```bash
go tool pprof mem.prof
# (pprof) top
# (pprof) list main.main
```

---

## Java: Managed Memory with GC

### Generational Garbage Collection

Java divides heap into generations based on object age:

```
┌─────────────────────────────────────────────┐
│              Java Heap                      │
├─────────────────────────┬───────────────────┤
│     Young Generation    │ Old Generation    │
├────────┬────────┬───────┤                   │
│  Eden  │   S0   │   S1  │   (Tenured)       │
└────────┴────────┴───────┴───────────────────┘
```

**Regions:**
- **Eden:** New objects allocated here
- **Survivor 0/1:** Objects that survived 1+ GC cycles
- **Old/Tenured:** Long-lived objects promoted from young gen

### GC Types

#### 1. Minor GC (Young Generation)
```
Before Minor GC:
┌──────────────────┬─────┬─────┬──────────┐
│  Eden (full)     │ S0  │ S1  │   Old    │
└──────────────────┴─────┴─────┴──────────┘

After Minor GC:
┌──────────────────┬─────┬─────┬──────────┐
│  Eden (empty)    │     │ S0  │   Old    │
└──────────────────┴─────┴─────┴──────────┘
       ↓                    ↑       ↑
   Collected          Survived   Promoted
```

#### 2. Major/Full GC (Entire Heap)
- Slower, more comprehensive
- Triggered when old generation fills up

### Modern GC Algorithms

**G1GC (Garbage-First):**
```bash
java -XX:+UseG1GC -XX:MaxGCPauseMillis=200 MyApp
```
- Divides heap into regions
- Prioritizes regions with most garbage
- Predictable pause times

**ZGC (Low-Latency):**
```bash
java -XX:+UseZGC -Xmx16g MyApp
```
- Sub-millisecond pause times
- Concurrent marking and compaction
- Handles multi-TB heaps

**Shenandoah:**
```bash
java -XX:+UseShenandoahGC MyApp
```
- Ultra-low pause times
- Concurrent evacuation

### Example: Reference Types

```java
import java.lang.ref.*;

public class ReferenceTypes {
    public static void main(String[] args) {
        // Strong reference (normal)
        Object strong = new Object();

        // Soft reference (cleared when memory low)
        SoftReference<Object> soft = new SoftReference<>(new Object());

        // Weak reference (cleared at next GC)
        WeakReference<Object> weak = new WeakReference<>(new Object());

        // Phantom reference (for cleanup actions)
        ReferenceQueue<Object> queue = new ReferenceQueue<>();
        PhantomReference<Object> phantom = new PhantomReference<>(new Object(), queue);

        System.gc();  // Suggest GC (not guaranteed)

        System.out.println("Soft: " + soft.get());     // Might be null
        System.out.println("Weak: " + weak.get());     // Likely null
        System.out.println("Phantom: " + phantom.get()); // Always null
    }
}
```

### Memory Monitoring

```java
import java.lang.management.*;

public class MemoryMonitor {
    public static void main(String[] args) {
        MemoryMXBean memoryBean = ManagementFactory.getMemoryMXBean();

        MemoryUsage heapUsage = memoryBean.getHeapMemoryUsage();

        System.out.println("Init: " + heapUsage.getInit() / (1024*1024) + " MB");
        System.out.println("Used: " + heapUsage.getUsed() / (1024*1024) + " MB");
        System.out.println("Committed: " + heapUsage.getCommitted() / (1024*1024) + " MB");
        System.out.println("Max: " + heapUsage.getMax() / (1024*1024) + " MB");
    }
}
```

---

## Performance Implications

### Benchmark Comparison

**Memory Allocation Speed (Approximate):**

| Operation | Rust | Go | Java |
|-----------|------|-----|------|
| Stack allocation | ~1 ns | ~1 ns | ~1 ns |
| Small heap allocation | ~10-50 ns | ~20-100 ns | ~50-200 ns |
| Large allocation (1MB) | ~50 µs | ~100 µs | ~100 µs |
| Deallocation | Immediate | GC cycle | GC cycle |

**Latency:**

| Language | Max Pause | P99 Latency |
|----------|-----------|-------------|
| Rust | None (deterministic) | Microseconds |
| Go | 1-10 ms | Milliseconds |
| Java (G1) | 10-200 ms | Milliseconds |
| Java (ZGC) | <1 ms | Sub-millisecond |

### Use Case Recommendations

**Choose Rust when:**
- Real-time systems (no GC pauses)
- Embedded systems (limited memory)
- High-performance computing
- Systems programming

**Choose Go when:**
- Web services (acceptable latency)
- Microservices (fast development)
- Concurrent applications
- Cloud-native tools

**Choose Java when:**
- Enterprise applications
- Large-scale systems (mature ecosystem)
- Android development
- When GC pause times are acceptable

---

## Memory Visualization

### Tool Recommendations

**Rust:**
```bash
# Valgrind for memory profiling
valgrind --leak-check=full --show-leak-kinds=all ./target/debug/myapp

# Heaptrack
heaptrack ./target/release/myapp
heaptrack_gui heaptrack.myapp.*.gz
```

**Go:**
```bash
# Built-in profiling
go tool pprof -http=:8080 mem.prof

# Trace viewer
go tool trace trace.out
```

**Java:**
```bash
# VisualVM
jvisualvm

# Java Flight Recorder
java -XX:+FlightRecorder -XX:StartFlightRecording=filename=app.jfr MyApp
jfr print --events GarbageCollection app.jfr
```

---

## Best Practices

### Rust

1. **Prefer stack allocation**
   ```rust
   let x = [0; 1000];  // Stack (if size known)
   let y = vec![0; 1000];  // Heap (dynamic)
   ```

2. **Use `Rc` for shared ownership**
   ```rust
   use std::rc::Rc;
   let shared = Rc::new(vec![1, 2, 3]);
   let clone1 = Rc::clone(&shared);
   let clone2 = Rc::clone(&shared);
   ```

3. **Leverage `Cow` for copy-on-write**
   ```rust
   use std::borrow::Cow;
   fn process(s: Cow<str>) {
       // Borrows if possible, clones only when modified
   }
   ```

### Go

1. **Reduce allocations**
   ```go
   // Bad: allocates on every call
   func bad() []int {
       return make([]int, 100)
   }

   // Good: reuse slice
   var pool = make([]int, 100)
   func good() []int {
       return pool[:0]  // Reset length, reuse capacity
   }
   ```

2. **Use sync.Pool for temporary objects**
   ```go
   var bufferPool = sync.Pool{
       New: func() interface{} {
           return new(bytes.Buffer)
       },
   }

   func process() {
       buf := bufferPool.Get().(*bytes.Buffer)
       defer bufferPool.Put(buf)
       buf.Reset()
       // Use buf...
   }
   ```

### Java

1. **Right-size collections**
   ```java
   // Bad: will resize multiple times
   List<String> list = new ArrayList<>();

   // Good: preallocate if size known
   List<String> list = new ArrayList<>(1000);
   ```

2. **Use primitive arrays over boxed types**
   ```java
   // Prefer:
   int[] numbers = new int[1000];

   // Over:
   Integer[] numbers = new Integer[1000];  // More memory overhead
   ```

3. **Avoid finalizers, use try-with-resources**
   ```java
   // Bad: finalizers are unpredictable
   @Override
   protected void finalize() { /* cleanup */ }

   // Good: explicit resource management
   try (FileInputStream fis = new FileInputStream("file.txt")) {
       // Use fis...
   }  // Automatically closed
   ```

---

## Summary Comparison

| Aspect | Rust | Go | Java |
|--------|------|-----|------|
| **Strategy** | Ownership | Garbage Collection | Garbage Collection |
| **Safety** | Compile-time | Runtime | Runtime |
| **Performance** | Highest | High | High |
| **Predictability** | Deterministic | Low-latency GC | Tunable GC |
| **Learning Curve** | Steep | Gentle | Moderate |
| **Use Cases** | Systems, Real-time | Web services, Tools | Enterprise, Android |

---

## Further Reading

**Rust:**
- [The Rustonomicon](https://doc.rust-lang.org/nomicon/) - Unsafe Rust and memory details
- [Rust Performance Book](https://nnethercote.github.io/perf-book/)

**Go:**
- [Go Memory Model](https://go.dev/ref/mem)
- [Go GC Guide](https://tip.golang.org/doc/gc-guide)

**Java:**
- [Java Memory Management White Paper](https://www.oracle.com/technetwork/java/javase/memorymanagement-whitepaper-150215.pdf)
- [JVM Internals](https://blog.jamesdbloom.com/JVMInternals.html)

---

**Next:** [Concurrency Models](./concurrency-models.md)
