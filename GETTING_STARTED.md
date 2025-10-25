# Quick Start Guide

This guide helps you get started quickly with Rust, Go, and Java. Choose your language and follow the steps!

## 📋 Prerequisites

- A computer running Windows, macOS, or Linux
- Internet connection for downloading tools
- A text editor (VS Code recommended)
- Basic command-line knowledge

## 🦀 Rust Quick Start

### 1. Install Rust

**All Platforms**: Visit [rustup.rs](https://rustup.rs/) and follow instructions, or:

```bash
# macOS/Linux
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh

# Windows: Download and run rustup-init.exe from rustup.rs
```

Verify installation:
```bash
rustc --version
cargo --version
```

### 2. Create Your First Project

```bash
cargo new hello_rust
cd hello_rust
```

### 3. Examine the Code

Look at `src/main.rs`:
```rust
fn main() {
    println!("Hello, world!");
}
```

### 4. Run Your Program

```bash
cargo run
```

You should see: `Hello, world!`

### 5. Next Steps

- Open [rust/README.md](./rust/README.md) for the full tutorial
- Start with [Lesson 1: Hello World](./rust/lessons/01-hello-world.md)
- Try modifying the message and running again

### Recommended Tools

- **VS Code** with rust-analyzer extension
- **IntelliJ IDEA** with Rust plugin
- Online: [Rust Playground](https://play.rust-lang.org/)

---

## 🐹 Go Quick Start

### 1. Install Go

**macOS**:
```bash
brew install go
```

**Linux (Ubuntu/Debian)**:
```bash
sudo apt update
sudo apt install golang-go
```

**Windows**: Download installer from [go.dev/dl](https://go.dev/dl/)

Verify installation:
```bash
go version
```

### 2. Create Your First Program

```bash
mkdir hello_go
cd hello_go
```

Create `hello.go`:
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

### 3. Run Your Program

```bash
go run hello.go
```

You should see: `Hello, Go!`

### 4. Build an Executable

```bash
go build hello.go
./hello  # or hello.exe on Windows
```

### 5. Next Steps

- Open [go/README.md](./go/README.md) for the full tutorial
- Start with [Lesson 1: Hello World](./go/lessons/01-hello-world.md)
- Explore [Go by Example](https://gobyexample.com/)

### Recommended Tools

- **VS Code** with Go extension
- **GoLand** by JetBrains
- Online: [Go Playground](https://go.dev/play/)

---

## ☕ Java Quick Start

### 1. Install Java JDK

**macOS**:
```bash
brew install openjdk@17
```

**Linux (Ubuntu/Debian)**:
```bash
sudo apt update
sudo apt install default-jdk
```

**Windows**: Download JDK from [Oracle](https://www.oracle.com/java/technologies/downloads/) or [AdoptOpenJDK](https://adoptopenjdk.net/)

Verify installation:
```bash
java -version
javac -version
```

### 2. Create Your First Program

```bash
mkdir hello_java
cd hello_java
```

Create `HelloWorld.java`:
```java
public class HelloWorld {
    public static void main(String[] args) {
        System.out.println("Hello, Java!");
    }
}
```

**Important**: Filename must match class name!

### 3. Compile Your Program

```bash
javac HelloWorld.java
```

This creates `HelloWorld.class`

### 4. Run Your Program

```bash
java HelloWorld
```

You should see: `Hello, Java!`

### 5. Next Steps

- Open [java/README.md](./java/README.md) for the full tutorial
- Start with [Lesson 1: Hello World](./java/lessons/01-hello-world.md)
- Try [CodingBat](https://codingbat.com/java) for practice

### Recommended Tools

- **IntelliJ IDEA Community Edition** (best for beginners)
- **Eclipse IDE**
- **VS Code** with Java Extension Pack
- Online: [JDoodle](https://www.jdoodle.com/online-java-compiler)

---

## 🎯 Choosing Your First Language

### Start with Go if:
- ✅ You're completely new to programming
- ✅ You want the simplest syntax
- ✅ You're interested in web services or CLI tools
- ✅ You want fast compilation and execution

### Start with Rust if:
- ✅ You want to understand low-level programming
- ✅ You're interested in systems programming
- ✅ You want to learn about memory management
- ✅ You enjoy understanding "why" things work

### Start with Java if:
- ✅ You want to learn Object-Oriented Programming deeply
- ✅ You're interested in Android development
- ✅ You want the largest job market
- ✅ You prefer mature, stable ecosystems

**Can't decide?** Start with Go - it's the easiest to pick up, and you can always learn the others later!

---

## 🐛 Troubleshooting Common Issues

### Rust

**Problem**: `cargo: command not found`
- **Solution**: Restart your terminal after installation, or manually add to PATH

**Problem**: `linking with cc failed`
- **Solution**: Install C compiler (Xcode on macOS, build-essential on Linux, Visual Studio on Windows)

### Go

**Problem**: `go: command not found`
- **Solution**: Check that Go is in your PATH. Run `export PATH=$PATH:/usr/local/go/bin`

**Problem**: `cannot find package`
- **Solution**: Run `go mod init myproject` to initialize a module

### Java

**Problem**: `javac: command not found`
- **Solution**: Install JDK (not just JRE) and ensure it's in PATH

**Problem**: `error: class HelloWorld is public, should be declared in a file named HelloWorld.java`
- **Solution**: Filename must exactly match the public class name (case-sensitive)

**Problem**: `Error: Could not find or load main class`
- **Solution**: Run with `java HelloWorld` (no .class extension)

---

## 📚 Learning Path Recommendations

### Week 1: Choose and Install
- Pick one language
- Install tools
- Complete Hello World
- Set up your development environment

### Week 2-3: Basics
- Variables and data types
- Control flow (if/else, loops)
- Functions
- Basic input/output

### Week 4-5: Intermediate Concepts
- Data structures (arrays, collections)
- Error handling
- Code organization (modules/packages)

### Week 6-8: Language-Specific Features
- **Rust**: Ownership, borrowing, lifetimes
- **Go**: Goroutines, channels, interfaces
- **Java**: OOP, classes, inheritance, polymorphism

### Ongoing: Build Projects
- Start small (calculator, to-do list)
- Gradually increase complexity
- Share your code for feedback

---

## 💡 Study Tips

1. **Code Every Day**: Even 30 minutes is valuable
2. **Type, Don't Copy**: Type out examples yourself
3. **Break Things**: Experiment and see what happens
4. **Read Error Messages**: They're teaching you
5. **Build Something**: Apply what you learn
6. **Join Communities**: Ask questions, help others
7. **Be Patient**: Programming takes time to learn

---

## 🤝 Getting Help

### Rust
- [Rust Users Forum](https://users.rust-lang.org/)
- [r/rust](https://reddit.com/r/rust)
- [Rust Discord](https://discord.gg/rust-lang)

### Go
- [Go Forum](https://forum.golangbridge.org/)
- [r/golang](https://reddit.com/r/golang)
- [Gophers Slack](https://gophers.slack.com/)

### Java
- [Stack Overflow - Java](https://stackoverflow.com/questions/tagged/java)
- [r/javahelp](https://reddit.com/r/javahelp)
- [r/java](https://reddit.com/r/java)

### All Languages
- [Stack Overflow](https://stackoverflow.com/)
- [GitHub Discussions](https://github.com/)
- This repository's Issues section

---

## 🎉 Ready to Start?

Pick your language and dive into the full tutorial:

- **[Rust Tutorial](./rust/README.md)** → Systems programming with safety
- **[Go Tutorial](./go/README.md)** → Simple, fast, and productive
- **[Java Tutorial](./java/README.md)** → Enterprise-ready OOP

**Remember**: Every expert programmer was once a beginner. Take it one step at a time, practice regularly, and don't be afraid to make mistakes. Happy coding! 🚀
