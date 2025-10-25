# Concurrency Models: Threading, Async, and Beyond

## Table of Contents
1. [Introduction to Concurrency](#introduction-to-concurrency)
2. [Concurrency vs Parallelism](#concurrency-vs-parallelism)
3. [Threading Models](#threading-models)
4. [Rust: Fearless Concurrency](#rust-fearless-concurrency)
5. [Go: Goroutines and Channels](#go-goroutines-and-channels)
6. [Java: Threads and Executors](#java-threads-and-executors)
7. [Synchronization Primitives](#synchronization-primitives)
8. [Async/Await Patterns](#asyncawait-patterns)
9. [Performance Analysis](#performance-analysis)
10. [Common Pitfalls](#common-pitfalls)

---

## Introduction to Concurrency

**Concurrency** is the composition of independently executing computations. It's about dealing with multiple things at once, whether they're actually running simultaneously (parallelism) or just managed together.

### Why Concurrency Matters

```
Without Concurrency:
Task 1 ████████████ (12s)
Task 2             ████████████ (12s)
Task 3                         ████████████ (12s)
Total: 36 seconds

With Concurrency:
Task 1 ████████████
Task 2 ████████████
Task 3 ████████████
Total: 12 seconds (3x faster)
```

### Key Benefits
- **Responsiveness**: UI remains responsive during long operations
- **Throughput**: Handle more requests per second
- **Resource Utilization**: Keep all CPU cores busy
- **Latency**: Multiple I/O operations in parallel

---

## Concurrency vs Parallelism

**Critical Distinction:**

```
Concurrency: Dealing with many things at once
Parallelism:  Doing many things at once
```

### Visual Representation

```
Concurrency (Single Core):
┌────────────────────────────┐
│ Task A ─ Task B ─ Task A ─│  ← Context switching
│ Task B ─ Task A ─ Task B ─│
└────────────────────────────┘
         Time →

Parallelism (Multi-Core):
Core 1: ████████████████████  Task A
Core 2: ████████████████████  Task B
         Time →
```

### Example Scenarios

**Concurrency without Parallelism:**
```
Single chef managing multiple dishes:
- Start pasta boiling
- While pasta cooks, chop vegetables
- While vegetables cook, prepare sauce
- Return to pasta, drain it
```

**Parallelism:**
```
Two chefs working simultaneously:
Chef 1: Makes pasta and sauce
Chef 2: Prepares salad and dessert
```

---

## Threading Models

### 1. OS Threads (Native Threads)

**Characteristics:**
- Managed by operating system
- 1:1 mapping with kernel threads
- Heavy (1-2 MB stack per thread)
- Expensive context switching

```
┌─────────────────────────────┐
│   Application Process       │
├─────────────────────────────┤
│ Thread 1 │ Thread 2 │ ...  │
└────┬──────────┬──────────────┘
     │          │
     ▼          ▼
┌─────────────────────────────┐
│   Kernel Space              │
│ Thread 1 │ Thread 2 │ ...  │
└─────────────────────────────┘
```

**Used by:** Java (default), Rust (std::thread)

### 2. Green Threads (User-Space Threads)

**Characteristics:**
- Managed by language runtime
- M:N mapping (many user threads to N kernel threads)
- Lightweight (few KB per thread)
- Fast context switching

```
┌─────────────────────────────┐
│   Application Process       │
│ G1 │ G2 │ G3 │ G4 │ ...    │  ← Many green threads
└─────────────────────────────┘
         │
         ▼
┌─────────────────────────────┐
│   Runtime Scheduler         │
└─────────────────────────────┘
         │
         ▼
┌─────────────────────────────┐
│   Kernel Space              │
│   T1   │   T2   │           │  ← Few OS threads
└─────────────────────────────┘
```

**Used by:** Go (goroutines), Erlang, Ruby fibers

### 3. Async/Await (Event Loop)

**Characteristics:**
- Cooperative multitasking
- Single-threaded or multi-threaded runtime
- Efficient for I/O-bound tasks
- Requires explicit yielding

```
┌─────────────────────────────┐
│       Event Loop            │
├─────────────────────────────┤
│  Task Queue                 │
│  ┌───┐ ┌───┐ ┌───┐         │
│  │ A │→│ B │→│ C │ ...     │
│  └───┘ └───┘ └───┘         │
└─────────────────────────────┘
         ▲
         │ await points
```

**Used by:** Rust (tokio, async-std), JavaScript, Python (asyncio)

---

## Rust: Fearless Concurrency

Rust's ownership system prevents data races at compile time!

### Thread-Based Concurrency

```rust
use std::thread;
use std::time::Duration;

fn main() {
    // Spawn a thread
    let handle = thread::spawn(|| {
        for i in 1..10 {
            println!("Thread: {}", i);
            thread::sleep(Duration::from_millis(1));
        }
    });

    // Main thread continues
    for i in 1..5 {
        println!("Main: {}", i);
        thread::sleep(Duration::from_millis(1));
    }

    // Wait for spawned thread to finish
    handle.join().unwrap();
}
```

**Output (interleaved):**
```
Main: 1
Thread: 1
Main: 2
Thread: 2
...
```

### Ownership Across Threads

```rust
use std::thread;

fn main() {
    let v = vec![1, 2, 3];

    // ❌ This won't compile - can't share mutable state
    // let handle = thread::spawn(|| {
    //     println!("Vector: {:?}", v);
    // });

    // ✅ Use `move` to transfer ownership
    let handle = thread::spawn(move || {
        println!("Vector: {:?}", v);
    });

    // println!("{:?}", v);  // ❌ Error: v was moved

    handle.join().unwrap();
}
```

### Message Passing with Channels

```rust
use std::sync::mpsc;
use std::thread;

fn main() {
    let (tx, rx) = mpsc::channel();

    // Spawn producer thread
    thread::spawn(move || {
        let vals = vec![
            String::from("hello"),
            String::from("from"),
            String::from("thread"),
        ];

        for val in vals {
            tx.send(val).unwrap();
            thread::sleep(Duration::from_millis(100));
        }
    });

    // Receive in main thread
    for received in rx {
        println!("Received: {}", received);
    }
}
```

**Channel Types:**
```
mpsc: Multiple Producer, Single Consumer
┌────────┐
│ Prod 1 │─┐
└────────┘ │
┌────────┐ │    ┌────────┐    ┌──────────┐
│ Prod 2 │─┼───▶│Channel │───▶│ Consumer │
└────────┘ │    └────────┘    └──────────┘
┌────────┐ │
│ Prod 3 │─┘
└────────┘
```

### Shared State with Arc and Mutex

```rust
use std::sync::{Arc, Mutex};
use std::thread;

fn main() {
    // Arc: Atomic Reference Counting (thread-safe Rc)
    // Mutex: Mutual Exclusion (one thread at a time)
    let counter = Arc::new(Mutex::new(0));
    let mut handles = vec![];

    for _ in 0..10 {
        let counter = Arc::clone(&counter);
        let handle = thread::spawn(move || {
            let mut num = counter.lock().unwrap();
            *num += 1;
        });
        handles.push(handle);
    }

    for handle in handles {
        handle.join().unwrap();
    }

    println!("Result: {}", *counter.lock().unwrap());  // 10
}
```

**Memory Layout:**
```
Arc<Mutex<i32>>
┌─────────────┐
│ Arc         │
│ ref_count: 3│
└──────┬──────┘
       │
       ▼
┌─────────────┐
│ Mutex       │
│ lock: false │
│ data: 42    │
└─────────────┘
```

### Async/Await in Rust

```rust
use tokio::time::{sleep, Duration};

#[tokio::main]
async fn main() {
    let task1 = async {
        sleep(Duration::from_millis(100)).await;
        println!("Task 1 done");
    };

    let task2 = async {
        sleep(Duration::from_millis(50)).await;
        println!("Task 2 done");
    };

    // Run concurrently
    tokio::join!(task1, task2);
}
```

**Under the Hood:**
```rust
// `async fn` desugars to:
fn my_async() -> impl Future<Output = ()> {
    async {
        // body
    }
}

// Polling mechanism:
loop {
    match future.poll(cx) {
        Poll::Ready(value) => return value,
        Poll::Pending => { /* yield control */ }
    }
}
```

---

## Go: Goroutines and Channels

Go makes concurrency a first-class citizen with goroutines and channels.

### Goroutines: Lightweight Threads

```go
package main

import (
    "fmt"
    "time"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

func main() {
    // Start a goroutine
    go say("world")

    // Main goroutine
    say("hello")
}
```

**Goroutine Scheduling:**
```
┌──────────────────────────────────┐
│   Go Runtime Scheduler           │
├──────────────────────────────────┤
│ Work-Stealing Scheduler          │
│  G1  G2  G3  G4  G5 ... (10,000+)│  ← Goroutines
└────┬─────┬─────┬─────────────────┘
     │     │     │
     ▼     ▼     ▼
  ┌────┬────┬────┐
  │ P1 │ P2 │ P3 │  ← Processors (logical cores)
  └────┴────┴────┘
     │     │     │
     ▼     ▼     ▼
  ┌────┬────┬────┐
  │ M1 │ M2 │ M3 │  ← OS Threads
  └────┴────┴────┘
```

**Key:** G = Goroutine, M = Machine (OS thread), P = Processor (logical CPU)

### Channels: Communication

```go
package main

import "fmt"

func sum(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum  // Send sum to channel
}

func main() {
    s := []int{7, 2, 8, -9, 4, 0}

    c := make(chan int)
    go sum(s[:len(s)/2], c)
    go sum(s[len(s)/2:], c)

    x, y := <-c, <-c  // Receive from channel

    fmt.Println(x, y, x+y)  // -5 17 12
}
```

### Buffered Channels

```go
func main() {
    // Unbuffered: blocks until receiver ready
    ch1 := make(chan int)

    // Buffered: blocks only when full
    ch2 := make(chan int, 3)

    ch2 <- 1  // Non-blocking
    ch2 <- 2  // Non-blocking
    ch2 <- 3  // Non-blocking
    // ch2 <- 4  // Would block (buffer full)

    fmt.Println(<-ch2)  // 1
    fmt.Println(<-ch2)  // 2
    fmt.Println(<-ch2)  // 3
}
```

**Visual:**
```
Unbuffered:
Sender ─────── Receiver
    (blocks until handoff)

Buffered (size 3):
Sender ─►[1][2][3]─► Receiver
    (blocks when full)
```

### Select: Multiplexing Channels

```go
func fibonacci(c, quit chan int) {
    x, y := 0, 1
    for {
        select {
        case c <- x:
            x, y = y, x+y
        case <-quit:
            fmt.Println("quit")
            return
        }
    }
}

func main() {
    c := make(chan int)
    quit := make(chan int)

    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(<-c)
        }
        quit <- 0
    }()

    fibonacci(c, quit)
}
```

### Worker Pool Pattern

```go
package main

import (
    "fmt"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, j)
        time.Sleep(time.Second)
        results <- j * 2
    }
}

func main() {
    const numJobs = 5
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    // Start 3 workers
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // Send jobs
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs)

    // Collect results
    for a := 1; a <= numJobs; a++ {
        <-results
    }
}
```

**Visualization:**
```
        ┌──────────┐
Jobs ───┤  Worker 1│─┐
     └──┤  Worker 2│─┼──► Results
        ┤  Worker 3│─┘
        └──────────┘
```

### sync Package Primitives

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var mu sync.Mutex
    var wg sync.WaitGroup
    counter := 0

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            mu.Lock()
            counter++
            mu.Unlock()
        }()
    }

    wg.Wait()
    fmt.Println("Counter:", counter)  // 1000
}
```

---

## Java: Threads and Executors

Java provides comprehensive threading support with multiple abstraction levels.

### Basic Thread Creation

```java
// Method 1: Extend Thread class
class MyThread extends Thread {
    @Override
    public void run() {
        for (int i = 0; i < 5; i++) {
            System.out.println("Thread: " + i);
            try {
                Thread.sleep(100);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }
}

// Method 2: Implement Runnable
class MyRunnable implements Runnable {
    @Override
    public void run() {
        for (int i = 0; i < 5; i++) {
            System.out.println("Runnable: " + i);
        }
    }
}

public class ThreadExample {
    public static void main(String[] args) {
        MyThread t1 = new MyThread();
        t1.start();

        Thread t2 = new Thread(new MyRunnable());
        t2.start();

        // Method 3: Lambda (Java 8+)
        Thread t3 = new Thread(() -> {
            System.out.println("Lambda thread");
        });
        t3.start();
    }
}
```

### Executor Framework

```java
import java.util.concurrent.*;

public class ExecutorExample {
    public static void main(String[] args) {
        // Fixed thread pool
        ExecutorService executor = Executors.newFixedThreadPool(3);

        for (int i = 0; i < 10; i++) {
            final int taskId = i;
            executor.submit(() -> {
                System.out.println("Task " + taskId + " on " +
                    Thread.currentThread().getName());
                try {
                    Thread.sleep(1000);
                } catch (InterruptedException e) {
                    e.printStackTrace();
                }
            });
        }

        executor.shutdown();
        try {
            executor.awaitTermination(1, TimeUnit.MINUTES);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }
    }
}
```

**Executor Types:**
```
newFixedThreadPool(n)
┌───┬───┬───┐
│ T1│ T2│ T3│  (n threads)
└───┴───┴───┘

newCachedThreadPool()
┌───┬───┬───┬───┐
│ T1│ T2│ T3│...│  (grows as needed)
└───┴───┴───┴───┘

newSingleThreadExecutor()
┌───┐
│ T1│  (1 thread)
└───┘

newScheduledThreadPool(n)
┌───┬───┐
│ T1│ T2│  (scheduled tasks)
└───┴───┘
```

### Future and Callable

```java
import java.util.concurrent.*;

public class FutureExample {
    public static void main(String[] args) throws Exception {
        ExecutorService executor = Executors.newFixedThreadPool(2);

        // Callable returns a result
        Callable<Integer> task = () -> {
            Thread.sleep(2000);
            return 42;
        };

        Future<Integer> future = executor.submit(task);

        System.out.println("Doing other work...");

        // Blocks until result is ready
        Integer result = future.get();
        System.out.println("Result: " + result);

        executor.shutdown();
    }
}
```

### CompletableFuture (Java 8+)

```java
import java.util.concurrent.CompletableFuture;

public class CompletableFutureExample {
    public static void main(String[] args) {
        // Async computation
        CompletableFuture<String> future = CompletableFuture.supplyAsync(() -> {
            sleep(1000);
            return "Hello";
        });

        // Chain operations
        future
            .thenApply(s -> s + " World")
            .thenApply(String::toUpperCase)
            .thenAccept(System.out::println);  // HELLO WORLD

        // Combine multiple futures
        CompletableFuture<String> f1 = CompletableFuture.supplyAsync(() -> "Hello");
        CompletableFuture<String> f2 = CompletableFuture.supplyAsync(() -> "World");

        CompletableFuture<String> combined = f1.thenCombine(f2, (s1, s2) -> s1 + " " + s2);
        System.out.println(combined.join());  // Hello World
    }

    static void sleep(int ms) {
        try { Thread.sleep(ms); } catch (Exception e) {}
    }
}
```

### Synchronized and Locks

```java
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

public class SynchronizationExample {
    // Method 1: synchronized keyword
    private int counter = 0;

    public synchronized void increment() {
        counter++;
    }

    // Method 2: synchronized block
    public void incrementBlock() {
        synchronized(this) {
            counter++;
        }
    }

    // Method 3: Explicit Lock
    private Lock lock = new ReentrantLock();

    public void incrementLock() {
        lock.lock();
        try {
            counter++;
        } finally {
            lock.unlock();
        }
    }
}
```

### Virtual Threads (Java 21+)

```java
// Traditional thread (heavy)
Thread.ofPlatform().start(() -> {
    System.out.println("Platform thread");
});

// Virtual thread (lightweight)
Thread.ofVirtual().start(() -> {
    System.out.println("Virtual thread");
});

// Executor with virtual threads
try (var executor = Executors.newVirtualThreadPerTaskExecutor()) {
    for (int i = 0; i < 10_000; i++) {
        executor.submit(() -> {
            Thread.sleep(Duration.ofSeconds(1));
            return i;
        });
    }
}  // Blocks until all tasks complete
```

**Virtual Threads vs Platform Threads:**
```
Platform Threads:
┌────────────────────────┐
│ App Thread Pool (100)  │
└────────┬───────────────┘
         │ 1:1
         ▼
┌────────────────────────┐
│ OS Threads (100)       │
└────────────────────────┘

Virtual Threads:
┌────────────────────────┐
│ Virtual Threads (1M+)  │
└────────┬───────────────┘
         │ M:N
         ▼
┌────────────────────────┐
│ Carrier Threads (10)   │
└────────┬───────────────┘
         │ 1:1
         ▼
┌────────────────────────┐
│ OS Threads (10)        │
└────────────────────────┘
```

---

## Synchronization Primitives

### Comparison Table

| Primitive | Rust | Go | Java |
|-----------|------|-----|------|
| **Mutex** | `Mutex<T>` | `sync.Mutex` | `synchronized`, `ReentrantLock` |
| **Read-Write Lock** | `RwLock<T>` | `sync.RWMutex` | `ReentrantReadWriteLock` |
| **Atomic Operations** | `AtomicI32`, etc. | `sync/atomic` | `AtomicInteger`, etc. |
| **Condition Variable** | `Condvar` | `sync.Cond` | `Condition` |
| **Once** | `Once` | `sync.Once` | N/A (use `static` initializer) |
| **Barrier** | `Barrier` | `sync.WaitGroup` | `CyclicBarrier`, `CountDownLatch` |

### Example: Atomic Operations

**Rust:**
```rust
use std::sync::atomic::{AtomicU32, Ordering};
use std::sync::Arc;
use std::thread;

let counter = Arc::new(AtomicU32::new(0));
let mut handles = vec![];

for _ in 0..10 {
    let counter = Arc::clone(&counter);
    let handle = thread::spawn(move || {
        counter.fetch_add(1, Ordering::SeqCst);
    });
    handles.push(handle);
}

for handle in handles {
    handle.join().unwrap();
}

println!("{}", counter.load(Ordering::SeqCst));  // 10
```

**Go:**
```go
import "sync/atomic"

var counter uint64 = 0

for i := 0; i < 10; i++ {
    go func() {
        atomic.AddUint64(&counter, 1)
    }()
}

time.Sleep(time.Second)
fmt.Println(atomic.LoadUint64(&counter))  // 10
```

**Java:**
```java
import java.util.concurrent.atomic.AtomicInteger;

AtomicInteger counter = new AtomicInteger(0);

for (int i = 0; i < 10; i++) {
    new Thread(() -> counter.incrementAndGet()).start();
}

Thread.sleep(1000);
System.out.println(counter.get());  // 10
```

---

## Async/Await Patterns

### Comparison

| Feature | Rust (tokio) | Java (CompletableFuture) |
|---------|--------------|--------------------------|
| **Syntax** | `async`/`await` | Callback chains |
| **Runtime** | Explicit (tokio, async-std) | Built-in |
| **Cancellation** | Drop future | `cancel()` method |
| **Error Handling** | `Result<T, E>` | Exceptions |

### Error Handling in Async Rust

```rust
use tokio::fs::File;
use tokio::io::AsyncReadExt;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let mut file = File::open("data.txt").await?;
    let mut contents = String::new();
    file.read_to_string(&mut contents).await?;
    println!("{}", contents);
    Ok(())
}
```

---

## Performance Analysis

### Context Switch Overhead

| Type | Latency |
|------|---------|
| OS Thread switch | ~1-10 µs |
| Goroutine switch | ~200 ns |
| Async task switch | ~50-100 ns |

### Memory Overhead

| Type | Stack Size |
|------|------------|
| OS Thread (Java/Rust) | 1-2 MB |
| Goroutine | 2-4 KB (growable) |
| Async task (Rust) | Minimal (state machine) |

### Scalability

**Goroutines:**
```
10,000 goroutines = ~40 MB
100,000 goroutines = ~400 MB
```

**OS Threads:**
```
10,000 threads = ~10 GB (not practical)
```

**Async Tasks:**
```
10,000 tasks = ~10 MB (state machines)
```

---

## Common Pitfalls

### 1. Data Races

**Problem:**
```rust
// ❌ Won't compile in Rust
let mut v = vec![1, 2, 3];
thread::spawn(|| v.push(4));
v.push(5);  // Race condition!
```

**Solution:**
```rust
// ✅ Use Mutex
let v = Arc::new(Mutex::new(vec![1, 2, 3]));
```

### 2. Deadlocks

**Problem:**
```go
// ❌ Deadlock: both goroutines wait forever
ch := make(chan int)
ch <- 42  // Blocks forever (no receiver)
```

**Solution:**
```go
// ✅ Use buffered channel or spawn goroutine
ch := make(chan int, 1)
ch <- 42
```

### 3. Goroutine Leaks

**Problem:**
```go
// ❌ Goroutine never terminates
go func() {
    for {
        // No exit condition
        doWork()
    }
}()
```

**Solution:**
```go
// ✅ Use context for cancellation
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

go func() {
    for {
        select {
        case <-ctx.Done():
            return
        default:
            doWork()
        }
    }
}()
```

---

## Best Practices

### Rust
1. Prefer message passing (channels) over shared state
2. Use `Arc` only when necessary
3. Consider `async` for I/O-bound tasks
4. Use `rayon` for data parallelism

### Go
1. "Don't communicate by sharing memory; share memory by communicating"
2. Use channels to orchestrate goroutines
3. Always handle goroutine termination
4. Limit goroutine creation (use worker pools for intensive tasks)

### Java
1. Use `ExecutorService` instead of raw threads
2. Prefer `CompletableFuture` for async operations
3. Consider virtual threads (Java 21+) for massive concurrency
4. Use `synchronized` sparingly, prefer `java.util.concurrent` primitives

---

## Summary

| Aspect | Rust | Go | Java |
|--------|------|-----|------|
| **Model** | Threads + Async | Goroutines | Threads + Virtual Threads |
| **Safety** | Compile-time | Runtime | Runtime |
| **Overhead** | Low (async) / Medium (threads) | Low | Medium / Low (virtual) |
| **Ease of Use** | Moderate | High | Moderate |
| **Best For** | Systems, Real-time | Web services, I/O | Enterprise, General |

---

## Further Reading

**Rust:**
- [The Rust Book: Fearless Concurrency](https://doc.rust-lang.org/book/ch16-00-concurrency.html)
- [Tokio Tutorial](https://tokio.rs/tokio/tutorial)

**Go:**
- [Go Concurrency Patterns](https://go.dev/blog/pipelines)
- [Effective Go: Concurrency](https://go.dev/doc/effective_go#concurrency)

**Java:**
- [Java Concurrency in Practice](https://jcip.net/)
- [JEP 425: Virtual Threads](https://openjdk.org/jeps/425)

---

**Next:** [Type Systems](./type-systems.md)
