# Go Exercises

Hands-on exercises to master Go programming.

## Structure

Each exercise contains:
- `README.md`: Problem description
- `exercise.go`: Your solution
- `exercise_test.go`: Tests
- `solution.go`: Reference solution (don't peek!)

## Running Exercises

```bash
# Navigate to exercise directory
cd exercises/01-variables

# Run tests
go test -v

# Run with coverage
go test -cover

# Format code
go fmt
```

## Exercise List

### Basics
1. **variables** - Variable declaration and types
2. **functions** - Functions and multiple returns
3. **control-structures** - If, for, switch
4. **slices** - Working with slices

### Structures
5. **structs** - Define and use structs
6. **methods** - Attach methods to types
7. **interfaces** - Interface implementation
8. **errors** - Error handling patterns

### Concurrency
9. **goroutines** - Launch goroutines
10. **channels** - Communicate via channels
11. **select** - Multiplexing channels
12. **context** - Cancellation and timeouts

### Standard Library
13. **json** - JSON marshaling
14. **http** - HTTP client/server
15. **io** - Input/output operations

## Tips

- Read error messages carefully
- Use `go doc` to read documentation
- Run `go vet` to catch common mistakes
- Follow [Effective Go](https://go.dev/doc/effective_go)

## Progress

Track your progress in `../progress.json`.
