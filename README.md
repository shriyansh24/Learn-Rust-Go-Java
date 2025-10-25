# Learn Rust, Go, and Java: A Comprehensive Learning Platform

![Rust CI](https://github.com/shriyansh24/Learn-Rust-Go-Java/workflows/Rust%20CI/badge.svg)
![Go CI](https://github.com/shriyansh24/Learn-Rust-Go-Java/workflows/Go%20CI/badge.svg)
![Java CI](https://github.com/shriyansh24/Learn-Rust-Go-Java/workflows/Java%20CI/badge.svg)
![License](https://img.shields.io/badge/license-MIT-blue.svg)

Welcome to your **production-grade learning platform** for mastering three powerful programming languages: **Rust**, **Go (Golang)**, and **Java**. This repository goes far beyond basic tutorials, offering deep technical insights, cross-language comparisons, and real-world projects.

## 🎯 What Makes This Different?

This is not just another tutorial repository. We provide:

### Deep Technical Content
- **Under-the-Hood Explanations**: Understand memory models, compilation pipelines, and runtime behavior
- **Cross-Language Comparisons**: See how each language solves the same problems differently
- **Performance Analysis**: Benchmark data, memory profiles, and optimization techniques
- **Production Patterns**: Industry-standard practices and design patterns

### Comprehensive Learning Resources
- **Progressive Lessons**: From "Hello World" to advanced concurrency patterns
- **Interactive Exercises**: Auto-graded coding challenges with instant feedback
- **Real-World Projects**: Build CLI tools, web servers, databases, and compilers
- **Visual Learning**: Mermaid diagrams, memory layouts, and execution flows

### Professional Development Tools
- **Debugging Guides**: Master GDB, Delve, and JDB with practical examples
- **Profiling Tutorials**: CPU, memory, and I/O profiling across all languages
- **Testing Strategies**: Unit tests, integration tests, and benchmarking
- **CI/CD Integration**: GitHub Actions workflows for automated testing

## 📚 Repository Structure

```
Learn-Rust-Go-Java/
├── common/                    # Cross-language concepts
│   ├── concepts/             # Deep dives into CS fundamentals
│   │   ├── memory-management.md
│   │   ├── concurrency-models.md
│   │   ├── type-systems.md
│   │   └── compilation-vs-interpretation.md
│   ├── projects/             # Language-comparison projects
│   │   ├── 01-task-tracker/
│   │   ├── 02-web-server/
│   │   ├── 03-database/
│   │   └── 04-compiler/
│   └── tools/                # Development tools guides
│       ├── debugging/
│       ├── profiling/
│       └── testing/
├── rust/
│   ├── lessons/              # Step-by-step tutorials
│   ├── exercises/            # Practice problems
│   ├── deep-dives/           # Advanced topics
│   ├── projects/             # Rust-specific projects
│   ├── examples/             # Code examples
│   └── progress.json         # Track your progress
├── go/
│   └── [same structure as rust/]
├── java/
│   └── [same structure as rust/]
├── ROADMAP.md                # Learning paths with Mermaid diagrams
└── GETTING_STARTED.md        # Setup instructions
```

## 🎓 Learning Paths

### Core Concepts (Start Here!)
Understand fundamental computer science concepts across all languages:

- **[Memory Management](./common/concepts/memory-management.md)**: Stack vs Heap, Garbage Collection, Ownership
- **[Concurrency Models](./common/concepts/concurrency-models.md)**: Threads, Goroutines, Async/Await
- **[Type Systems](./common/concepts/type-systems.md)**: Static vs Dynamic, Generics, Type Inference
- **[Compilation vs Interpretation](./common/concepts/compilation-vs-interpretation.md)**: AOT, JIT, Bytecode

### Language-Specific Paths

#### [Rust Tutorial](./rust/README.md)
**Systems Programming, Memory Safety, Zero-Cost Abstractions**
- Memory safety without garbage collection
- Ownership, borrowing, and lifetimes
- Fearless concurrency with compile-time guarantees
- Building high-performance systems
- [30 Lessons](./rust/progress.json) | [16 Exercises](./rust/exercises/) | [4 Projects](./common/projects/)

#### [Go Tutorial](./go/README.md)
**Simplicity, Concurrency, Cloud-Native Development**
- Clean, minimalist syntax
- Goroutines and channels for easy concurrency
- Fast compilation and execution
- Building scalable web services
- [29 Lessons](./go/progress.json) | [15 Exercises](./go/exercises/) | [4 Projects](./common/projects/)

#### [Java Tutorial](./java/README.md)
**Enterprise Development, Object-Oriented Programming, Platform Independence**
- Robust OOP with classes and interfaces
- Write Once, Run Anywhere (WORA)
- Extensive standard library and frameworks
- Android app development
- [37 Lessons](./java/progress.json) | [16 Exercises](./java/exercises/) | [4 Projects](./common/projects/)

## 🛠️ Real-World Projects

Build progressively complex projects that demonstrate practical application:

### [Project 1: Task Tracker](./common/projects/01-task-tracker/)
**Difficulty**: Beginner | **Time**: 4-8 hours
- CLI application with file persistence
- Learn: Data structures, file I/O, error handling
- Implement in all three languages and compare

### [Project 2: Web Server](./common/projects/02-web-server/)
**Difficulty**: Intermediate | **Time**: 8-12 hours
- Concurrent HTTP server with REST API
- Learn: Concurrency, networking, JSON handling
- Performance benchmark across languages

### [Project 3: Key-Value Database](./common/projects/03-database/)
**Difficulty**: Advanced | **Time**: 16-24 hours
- Storage engine with WAL and indexing
- Learn: B-trees, transactions, ACID properties
- Production-grade persistence layer

### [Project 4: Compiler](./common/projects/04-compiler/)
**Difficulty**: Expert | **Time**: 24-40 hours
- Lexer, parser, type checker, code generator
- Learn: Language theory, ASTs, optimization
- Compile to bytecode or native code

## 🚀 How to Use This Repository

### Quick Start
👉 **[Get Started Now](./GETTING_STARTED.md)** - Installation and setup for all languages
👉 **[View Roadmap](./ROADMAP.md)** - Visual learning paths with Mermaid diagrams

### Recommended Learning Sequence

**For Beginners:**
```
1. Read Getting Started Guide
2. Choose: Go → Java → Rust (easiest to hardest)
3. Complete lessons in order
4. Practice with exercises
5. Build projects
6. Read deep-dive concepts
```

**For Experienced Programmers:**
```
1. Read cross-language concepts first (common/)
2. Choose: Rust → Go → Java (most interesting to familiar)
3. Skip basics, focus on unique features
4. Build all 4 projects
5. Compare implementations
```

### Learning Strategy
1. **Code Every Day**: Consistency beats intensity (30 min/day minimum)
2. **Type, Don't Copy**: Muscle memory matters
3. **Read Error Messages**: They're your teachers
4. **Debug Actively**: Use debuggers, not just print statements
5. **Test Everything**: Write tests as you code
6. **Teach Others**: Best way to solidify understanding

## 📋 Prerequisites

- Basic computer literacy
- Ability to use a text editor
- Curiosity and willingness to learn
- No prior programming experience needed!

## 🛠️ Setup Instructions

Each language directory contains setup instructions for:
- Installing the necessary tools
- Setting up your development environment
- Running your first program
- Debugging common issues

## 🔧 Professional Development Tools

Master the tools that professionals use daily:

### [Debugging Guide](./common/tools/debugging/)
- **Rust**: GDB, LLDB, rust-gdb, VSCode integration
- **Go**: Delve debugger, pprof, race detector
- **Java**: JDB, IntelliJ debugger, VisualVM
- Core dumps, remote debugging, conditional breakpoints

### [Profiling Guide](./common/tools/profiling/)
- **CPU Profiling**: perf, flamegraphs, async-profiler
- **Memory Profiling**: Valgrind, heaptrack, Eclipse MAT
- **Benchmarking**: Criterion, go test -bench, JMH
- Performance optimization techniques

### [Testing Guide](./common/tools/testing/)
- **Unit Testing**: cargo test, go test, JUnit 5
- **Integration Testing**: Test strategies for each language
- **Mocking**: mockall, testify, Mockito
- **Coverage**: tarpaulin, go cover, JaCoCo

## 🤔 Which Language Should I Learn First?

### Based on Your Goals

| Goal | Recommended Language | Why? |
|------|---------------------|------|
| **First Programming Language** | **Go** | Simplest syntax, gentle curve, modern design |
| **Systems Programming** | **Rust** | Memory control, performance, safety guarantees |
| **Web Development** | **Go** or **Java** | Go for microservices, Java for enterprise |
| **Android Development** | **Java** | Official Android language (with Kotlin) |
| **Game Development** | **Rust** | Performance critical, memory control |
| **Cloud/DevOps** | **Go** | Fast compilation, static binaries, popular in cloud |
| **Data Structures & Algorithms** | **Java** | Extensive collections, OOP patterns |
| **Embedded Systems** | **Rust** | No runtime overhead, safety without GC |

### Learning Difficulty (1-10)

- **Go**: 4/10 - Simplest, fastest to productivity
- **Java**: 6/10 - OOP concepts take time
- **Rust**: 9/10 - Steep but rewarding

## 📊 Comparison Matrix

| Aspect | Rust | Go | Java |
|--------|------|-----|------|
| **Syntax Complexity** | High | Low | Medium |
| **Compilation Speed** | Slow | Fast | Medium (to bytecode) |
| **Runtime Performance** | Excellent | Great | Good (after JIT warmup) |
| **Memory Safety** | Compile-time | Runtime | Runtime (with GC) |
| **Concurrency Model** | Threads + Async | Goroutines | Threads + Virtual Threads |
| **Learning Curve** | Steep | Gentle | Moderate |
| **Ecosystem Maturity** | Growing | Mature | Very Mature |
| **Job Market** | Growing | Strong | Largest |
| **Best For** | Systems, Performance | Cloud, Web Services | Enterprise, Android |

## 🤝 Contributing

Found a typo? Have a better explanation? Want to add more content? Contributions are welcome!

Please read our **[Contributing Guide](./CONTRIBUTING.md)** to get started.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.

## 🙏 Acknowledgments

Special thanks to all contributors who help make programming education accessible to everyone!

## ⭐ Star This Repository

If you find this helpful, please star the repository to help others discover it!

---

**Remember**: Programming is a skill that develops with practice. Don't get discouraged if concepts don't click immediately. Keep coding, keep experimenting, and most importantly - have fun! 🎉