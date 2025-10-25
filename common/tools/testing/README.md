# Testing Guide: Comprehensive Testing Strategies

## Overview

Master unit testing, integration testing, and test-driven development (TDD) across Rust, Go, and Java.

---

## Testing Pyramid

```
       ╱╲
      ╱  ╲
     ╱ E2E ╲        Few, slow, expensive
    ╱────────╲
   ╱          ╲
  ╱Integration╲    Some, medium speed
 ╱──────────────╲
╱                ╲
╱   Unit Tests   ╲  Many, fast, cheap
────────────────────
```

---

## Rust Testing

### Unit Tests

```rust
#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_addition() {
        assert_eq!(add(2, 2), 4);
    }

    #[test]
    #[should_panic(expected = "divide by zero")]
    fn test_divide_by_zero() {
        divide(10, 0);
    }

    #[test]
    fn test_result() -> Result<(), String> {
        if add(2, 2) == 4 {
            Ok(())
        } else {
            Err(String::from("Addition failed"))
        }
    }
}
```

Run: `cargo test`

### Integration Tests

```rust
// tests/integration_test.rs
use my_crate::public_function;

#[test]
fn test_public_api() {
    assert_eq!(public_function(), expected_value);
}
```

### Mocking

```rust
// Using mockall
use mockall::predicate::*;
use mockall::*;

#[automock]
trait Database {
    fn get_user(&self, id: u32) -> Option<User>;
}

#[test]
fn test_with_mock() {
    let mut mock = MockDatabase::new();
    mock.expect_get_user()
        .with(eq(1))
        .returning(|_| Some(User { id: 1, name: "Alice".to_string() }));

    assert_eq!(mock.get_user(1).unwrap().name, "Alice");
}
```

---

## Go Testing

### Unit Tests

```go
// main_test.go
package main

import "testing"

func TestAddition(t *testing.T) {
    result := add(2, 2)
    expected := 4

    if result != expected {
        t.Errorf("Expected %d, got %d", expected, result)
    }
}

func TestTable(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive", 2, 3, 5},
        {"negative", -1, -1, -2},
        {"zero", 0, 0, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Expected %d, got %d", tt.expected, result)
            }
        })
    }
}
```

Run: `go test -v`

### Benchmarks

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        add(2, 3)
    }
}
```

Run: `go test -bench=.`

### Mocking

```go
// Using testify/mock
type MockDB struct {
    mock.Mock
}

func (m *MockDB) GetUser(id int) (*User, error) {
    args := m.Called(id)
    return args.Get(0).(*User), args.Error(1)
}

func TestWithMock(t *testing.T) {
    mockDB := new(MockDB)
    mockDB.On("GetUser", 1).Return(&User{ID: 1, Name: "Alice"}, nil)

    user, _ := mockDB.GetUser(1)
    assert.Equal(t, "Alice", user.Name)

    mockDB.AssertExpectations(t)
}
```

---

## Java Testing

### JUnit 5

```java
import org.junit.jupiter.api.*;
import static org.junit.jupiter.api.Assertions.*;

class CalculatorTest {

    @Test
    void testAddition() {
        assertEquals(4, add(2, 2));
    }

    @Test
    void testDivisionByZero() {
        assertThrows(ArithmeticException.class, () -> divide(10, 0));
    }

    @ParameterizedTest
    @CsvSource({
        "2, 3, 5",
        "-1, -1, -2",
        "0, 0, 0"
    })
    void testAdditionWithParameters(int a, int b, int expected) {
        assertEquals(expected, add(a, b));
    }

    @BeforeEach
    void setUp() {
        // Run before each test
    }

    @AfterEach
    void tearDown() {
        // Run after each test
    }
}
```

Run: `mvn test`

### Mockito

```java
import static org.mockito.Mockito.*;

@Test
void testWithMock() {
    Database mockDB = mock(Database.class);
    when(mockDB.getUser(1)).thenReturn(new User(1, "Alice"));

    User user = mockDB.getUser(1);
    assertEquals("Alice", user.getName());

    verify(mockDB).getUser(1);
}
```

---

## Test Coverage

**Rust:**
```bash
cargo install cargo-tarpaulin
cargo tarpaulin --out Html
```

**Go:**
```bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

**Java:**
```bash
mvn jacoco:report
# Open target/site/jacoco/index.html
```

---

## Best Practices

1. **Test behavior, not implementation**
2. **Keep tests independent**: No shared state
3. **Use descriptive names**: `test_user_login_with_invalid_password`
4. **Follow AAA pattern**: Arrange, Act, Assert
5. **Test edge cases**: Empty input, null, max values
6. **Avoid flaky tests**: No time dependencies, random values

---

**Previous:** [Profiling](../profiling/)
