# Profiling Guide: Performance Analysis

## Overview

Learn to identify performance bottlenecks using profiling tools for CPU, memory, and I/O analysis.

---

## Profiling Principles

1. **Measure before optimizing**: "Premature optimization is the root of all evil"
2. **Profile in production-like environment**: Dev != Production
3. **Focus on hot paths**: 80/20 rule applies
4. **Consider the whole system**: CPU, memory, I/O, network

---

## Rust Profiling

### CPU Profiling

**Using perf (Linux):**
```bash
# Compile with debug info
cargo build --release

# Record
perf record --call-graph=dwarf ./target/release/myapp

# View report
perf report

# Generate flame graph
git clone https://github.com/brendangregg/FlameGraph
perf script | ./FlameGraph/stackcollapse-perf.pl | ./FlameGraph/flamegraph.pl > flame.svg
```

**Using cargo-flamegraph:**
```bash
cargo install flamegraph

cargo flamegraph
# Opens flame.svg in browser
```

### Memory Profiling

**Using Valgrind:**
```bash
cargo build
valgrind --leak-check=full ./target/debug/myapp
```

**Using heaptrack:**
```bash
heaptrack ./target/release/myapp
heaptrack_gui heaptrack.myapp.*.gz
```

### Benchmarking

```rust
// Criterion.rs
use criterion::{black_box, criterion_group, criterion_main, Criterion};

fn fibonacci(n: u64) -> u64 {
    match n {
        0 => 0,
        1 => 1,
        n => fibonacci(n - 1) + fibonacci(n - 2),
    }
}

fn criterion_benchmark(c: &mut Criterion) {
    c.bench_function("fib 20", |b| b.iter(|| fibonacci(black_box(20))));
}

criterion_group!(benches, criterion_benchmark);
criterion_main!(benches);
```

```toml
[dev-dependencies]
criterion = "0.5"

[[bench]]
name = "my_benchmark"
harness = false
```

Run: `cargo bench`

---

## Go Profiling

### CPU Profiling

**Built-in pprof:**
```go
import (
    "os"
    "runtime/pprof"
)

func main() {
    f, _ := os.Create("cpu.prof")
    defer f.Close()

    pprof.StartCPUProfile(f)
    defer pprof.StopCPUProfile()

    // Your code here
}
```

Analyze:
```bash
go tool pprof cpu.prof
# Interactive mode
(pprof) top10
(pprof) list funcName
(pprof) web  # Graphical view (requires graphviz)

# Or web UI
go tool pprof -http=:8080 cpu.prof
```

**HTTP Endpoint:**
```go
import _ "net/http/pprof"

go func() {
    log.Println(http.ListenAndServe("localhost:6060", nil))
}()

// Access http://localhost:6060/debug/pprof/
```

### Memory Profiling

```go
import "runtime/pprof"

func main() {
    f, _ := os.Create("mem.prof")
    defer f.Close()

    // Your code here

    runtime.GC()  // Force GC before profiling
    pprof.WriteHeapProfile(f)
}
```

Analyze:
```bash
go tool pprof mem.prof
(pprof) top10
(pprof) list funcName
```

### Benchmarking

```go
// main_test.go
func BenchmarkFibonacci(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fibonacci(20)
    }
}

func BenchmarkFibonacciParallel(b *testing.B) {
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            fibonacci(20)
        }
    })
}
```

Run:
```bash
go test -bench=. -benchmem
go test -bench=. -cpuprofile=cpu.prof
go test -bench=. -memprofile=mem.prof
```

---

## Java Profiling

### JVM Profiling Tools

**1. JProfiler / YourKit (Commercial)**

**2. VisualVM (Free)**
```bash
jvisualvm
# Attach to running Java process
```

**3. Java Flight Recorder (JFR)**
```bash
# Start recording
java -XX:+FlightRecorder -XX:StartFlightRecording=duration=60s,filename=recording.jfr MyApp

# Or trigger during runtime
jcmd <pid> JFR.start duration=60s filename=recording.jfr

# View with Mission Control
jmc
```

### Async-profiler (Low-overhead)

```bash
# Download from https://github.com/jvm-profiling-tools/async-profiler

# Attach to running process
./profiler.sh -d 30 -f flamegraph.svg <pid>
```

### Memory Analysis

**Heap Dump:**
```bash
# Manual heap dump
jmap -dump:format=b,file=heap.hprof <pid>

# Or on OOM
java -XX:+HeapDumpOnOutOfMemoryError -XX:HeapDumpPath=/tmp/heap.hprof MyApp
```

**Analyze with Eclipse MAT:**
```bash
# Download from https://www.eclipse.org/mat/
# Open heap.hprof
# Run "Leak Suspects" report
```

### JMH Benchmarking

```java
import org.openjdk.jmh.annotations.*;

@State(Scope.Thread)
public class MyBenchmark {

    @Benchmark
    public int testFibonacci() {
        return fibonacci(20);
    }

    private int fibonacci(int n) {
        if (n <= 1) return n;
        return fibonacci(n - 1) + fibonacci(n - 2);
    }
}
```

```xml
<dependency>
    <groupId>org.openjdk.jmh</groupId>
    <artifactId>jmh-core</artifactId>
    <version>1.36</version>
</dependency>
```

Run:
```bash
mvn clean install
java -jar target/benchmarks.jar
```

---

## Profiling Comparison

| Aspect | Rust | Go | Java |
|--------|------|-----|------|
| **CPU Profiling** | perf, flamegraph | pprof | JFR, async-profiler |
| **Memory Profiling** | Valgrind, heaptrack | pprof | VisualVM, MAT |
| **Benchmarking** | Criterion | testing | JMH |
| **Overhead** | Very Low | Low | Medium |
| **Ease of Use** | Medium | High | High |

---

## Optimization Tips

### Rust

1. **Use release mode**: `cargo build --release`
2. **Profile-guided optimization**:
   ```toml
   [profile.release]
   lto = true
   codegen-units = 1
   ```
3. **Avoid unnecessary clones**: Use references
4. **Use appropriate data structures**: Vec vs HashMap

### Go

1. **Reduce allocations**: Use sync.Pool for temporary objects
2. **Avoid interface boxing**: Use concrete types when possible
3. **Buffer I/O**: Use bufio for file operations
4. **Preallocate slices**: `make([]int, 0, expectedSize)`

### Java

1. **JVM tuning**:
   ```bash
   -Xms2g -Xmx2g  # Set initial and max heap
   -XX:+UseG1GC   # Use G1 garbage collector
   ```
2. **Avoid autoboxing**: Use primitive arrays
3. **StringBuilder for string concat**: Avoid `+` in loops
4. **Use appropriate collection sizes**: Preallocate capacity

---

**Next:** [Testing](../testing/)
