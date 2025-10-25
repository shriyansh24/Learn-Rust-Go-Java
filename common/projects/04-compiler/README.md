# Project 4: Simple Programming Language Compiler

## Overview

Build a compiler for a simple expression-based language that compiles to bytecode or native code. Learn about lexing, parsing, type checking, and code generation.

**Difficulty:** Expert
**Time Estimate:** 24-40 hours
**Concepts Covered:** Lexical analysis, parsing, AST, type systems, code generation, optimization

---

## Language Specification

### Sample Program

```
// Variables and expressions
let x = 10;
let y = 20;
let sum = x + y;

// Functions
fn add(a: int, b: int) -> int {
    return a + b;
}

// Control flow
if (sum > 25) {
    print("Large sum");
} else {
    print("Small sum");
}

// Loops
while (x < 100) {
    x = x + 1;
}
```

---

## Compilation Pipeline

```
Source Code
    │
    ▼
┌─────────────┐
│   Lexer     │  Characters → Tokens
└──────┬──────┘
       ▼
┌─────────────┐
│   Parser    │  Tokens → AST
└──────┬──────┘
       ▼
┌─────────────┐
│  Type Check │  Validate types
└──────┬──────┘
       ▼
┌─────────────┐
│  Optimizer  │  Constant folding, DCE
└──────┬──────┘
       ▼
┌─────────────┐
│  Code Gen   │  AST → Bytecode/Assembly
└──────┬──────┘
       ▼
   Executable
```

---

## Phase Breakdown

### Phase 1: Lexer (Tokenization)
```
Input:  "let x = 42;"
Output: [LET, IDENT("x"), EQUALS, NUMBER(42), SEMICOLON]
```

### Phase 2: Parser (AST Construction)
```
let x = 42;

AST:
LetStatement {
  name: "x",
  value: IntegerLiteral(42)
}
```

### Phase 3: Semantic Analysis
- Symbol table construction
- Type checking
- Scope resolution

### Phase 4: Code Generation
```
PUSH 42
STORE x
```

---

## Enhancement Ideas

- Arrays and structs
- Higher-order functions
- Garbage collection
- LLVM backend
- Standard library
- Package system

---

## Resources

- [Crafting Interpreters](https://craftinginterpreters.com/)
- [Writing An Interpreter In Go](https://interpreterbook.com/)
- [LLVM Tutorial](https://llvm.org/docs/tutorial/)

---

**Previous:** [03-database](../03-database/)
