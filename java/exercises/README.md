# Java Exercises

Progressive exercises to learn Java from basics to advanced topics.

## Structure

Each exercise is a Maven project:
- `README.md`: Problem description
- `src/main/java/`: Your implementation
- `src/test/java/`: JUnit tests
- `pom.xml`: Maven configuration

## Running Exercises

```bash
# Navigate to exercise directory
cd exercises/01-variables

# Compile
mvn compile

# Run tests
mvn test

# Run specific test
mvn test -Dtest=VariablesTest

# Check style
mvn checkstyle:check
```

## Exercise List

### Fundamentals
1. **variables** - Variables and data types
2. **methods** - Method definition and overloading
3. **control-flow** - If, for, while, switch
4. **arrays** - Arrays and iteration

### OOP
5. **classes** - Classes and objects
6. **inheritance** - Extending classes
7. **polymorphism** - Method overriding
8. **interfaces** - Interface implementation
9. **abstract-classes** - Abstract classes

### Collections
10. **arraylist** - ArrayList operations
11. **hashmap** - HashMap usage
12. **streams** - Stream API

### Advanced
13. **generics** - Generic classes and methods
14. **lambda** - Lambda expressions
15. **threads** - Multi-threading
16. **completablefuture** - Async programming

## Tips

- Use an IDE (IntelliJ IDEA, Eclipse, VS Code)
- Read JavaDoc comments
- Follow naming conventions (camelCase)
- Write meaningful test cases
- Use Java 17+ features

## Resources

- [Java Tutorial](https://docs.oracle.com/javase/tutorial/)
- [Effective Java](https://www.oreilly.com/library/view/effective-java/9780134686097/)
- Lessons in `../lessons/`

## Progress

Update `../progress.json` as you complete exercises.
