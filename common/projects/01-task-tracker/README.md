# Project 1: Command-Line Task Tracker

## Overview

Build a command-line task tracker that demonstrates fundamental programming concepts: data structures, file I/O, error handling, and basic CRUD operations.

**Difficulty:** Beginner
**Time Estimate:** 4-8 hours
**Concepts Covered:**
- Data structures (structs, classes)
- File I/O and persistence
- Error handling
- Command-line argument parsing
- Testing

---

## Project Goals

By completing this project, you will:

1. **Understand data modeling**: Design a `Task` type with appropriate fields
2. **Master file I/O**: Persist tasks to JSON format
3. **Handle errors gracefully**: Deal with file operations, parsing errors
4. **Parse command-line arguments**: Build an intuitive CLI interface
5. **Write tests**: Ensure correctness with unit tests
6. **Compare implementations**: See how each language approaches the same problem

---

## Functional Requirements

### Core Features

Your task tracker must support these operations:

```bash
# Add a new task
$ task add "Buy groceries"
Task added successfully (ID: 1)

# List all tasks
$ task list
ID  Status      Task
1   todo        Buy groceries
2   in-progress Write report
3   done        Read chapter 5

# Update task status
$ task update 1 in-progress
Task 1 updated to "in-progress"

# Mark task as done
$ task done 2
Task 2 marked as done

# Delete a task
$ task delete 3
Task 3 deleted

# List tasks by status
$ task list todo
$ task list in-progress
$ task list done
```

### Data Model

Each task should have:

```
Task:
  - id: unique identifier (integer)
  - description: task text (string)
  - status: one of [todo, in-progress, done]
  - createdAt: timestamp
  - updatedAt: timestamp
```

### Persistence

- Store tasks in a JSON file (`tasks.json`)
- File location: current directory or user's home directory
- Handle missing file (create on first use)
- Validate JSON on read

---

## Technical Specifications

### File Format

```json
{
  "tasks": [
    {
      "id": 1,
      "description": "Buy groceries",
      "status": "todo",
      "createdAt": "2025-01-15T10:30:00Z",
      "updatedAt": "2025-01-15T10:30:00Z"
    },
    {
      "id": 2,
      "description": "Write report",
      "status": "in-progress",
      "createdAt": "2025-01-15T11:00:00Z",
      "updatedAt": "2025-01-15T14:20:00Z"
    }
  ],
  "nextId": 3
}
```

### Command-Line Interface

```
task <command> [arguments]

Commands:
  add <description>       Add a new task
  list [status]          List all tasks or filter by status
  update <id> <status>   Update task status
  done <id>              Mark task as done
  delete <id>            Delete a task
  help                   Show help message
```

### Error Handling

Handle these scenarios gracefully:

1. **File not found**: Create new file with empty task list
2. **Invalid JSON**: Show error, offer to reset file
3. **Invalid task ID**: Display error message
4. **Invalid status**: Show valid options
5. **Invalid command**: Show help message

---

## Implementation Guide

### Phase 1: Data Structures

**Define the Task type:**

<details>
<summary>Rust Hints</summary>

```rust
use serde::{Deserialize, Serialize};
use chrono::{DateTime, Utc};

#[derive(Debug, Serialize, Deserialize, Clone)]
pub enum TaskStatus {
    Todo,
    InProgress,
    Done,
}

#[derive(Debug, Serialize, Deserialize, Clone)]
pub struct Task {
    pub id: u32,
    pub description: String,
    pub status: TaskStatus,
    pub created_at: DateTime<Utc>,
    pub updated_at: DateTime<Utc>,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct TaskStore {
    pub tasks: Vec<Task>,
    pub next_id: u32,
}
```

**Dependencies (Cargo.toml):**
```toml
[dependencies]
serde = { version = "1.0", features = ["derive"] }
serde_json = "1.0"
chrono = { version = "0.4", features = ["serde"] }
```
</details>

<details>
<summary>Go Hints</summary>

```go
package main

import "time"

type TaskStatus string

const (
    StatusTodo       TaskStatus = "todo"
    StatusInProgress TaskStatus = "in-progress"
    StatusDone       TaskStatus = "done"
)

type Task struct {
    ID          int        `json:"id"`
    Description string     `json:"description"`
    Status      TaskStatus `json:"status"`
    CreatedAt   time.Time  `json:"createdAt"`
    UpdatedAt   time.Time  `json:"updatedAt"`
}

type TaskStore struct {
    Tasks  []Task `json:"tasks"`
    NextID int    `json:"nextId"`
}
```
</details>

<details>
<summary>Java Hints</summary>

```java
import java.time.Instant;
import com.fasterxml.jackson.annotation.JsonProperty;

public enum TaskStatus {
    TODO("todo"),
    IN_PROGRESS("in-progress"),
    DONE("done");

    private final String value;

    TaskStatus(String value) {
        this.value = value;
    }

    public String getValue() {
        return value;
    }
}

public class Task {
    private int id;
    private String description;
    private TaskStatus status;
    @JsonProperty("createdAt")
    private Instant createdAt;
    @JsonProperty("updatedAt")
    private Instant updatedAt;

    // Constructors, getters, setters...
}

public class TaskStore {
    private List<Task> tasks;
    private int nextId;

    // Methods...
}
```

**Dependencies (Maven pom.xml):**
```xml
<dependency>
    <groupId>com.fasterxml.jackson.core</groupId>
    <artifactId>jackson-databind</artifactId>
    <version>2.15.2</version>
</dependency>
```
</details>

### Phase 2: File I/O

Implement functions to:
- Load tasks from JSON file
- Save tasks to JSON file
- Handle file not found (create default)
- Handle invalid JSON (error message)

<details>
<summary>Rust Example</summary>

```rust
use std::fs;
use std::path::Path;

impl TaskStore {
    pub fn load(path: &Path) -> Result<Self, Box<dyn std::error::Error>> {
        if !path.exists() {
            return Ok(TaskStore {
                tasks: Vec::new(),
                next_id: 1,
            });
        }

        let data = fs::read_to_string(path)?;
        let store: TaskStore = serde_json::from_str(&data)?;
        Ok(store)
    }

    pub fn save(&self, path: &Path) -> Result<(), Box<dyn std::error::Error>> {
        let data = serde_json::to_string_pretty(self)?;
        fs::write(path, data)?;
        Ok(())
    }
}
```
</details>

### Phase 3: Core Operations

Implement CRUD operations:

```
add(description)    → Create new task
list(filter)        → Read all/filtered tasks
update(id, status)  → Update task status
delete(id)          → Delete task
```

### Phase 4: CLI Parsing

Parse command-line arguments and dispatch to appropriate functions.

<details>
<summary>Rust: Use clap crate</summary>

```toml
[dependencies]
clap = { version = "4.0", features = ["derive"] }
```

```rust
use clap::{Parser, Subcommand};

#[derive(Parser)]
#[command(name = "task")]
#[command(about = "A simple task tracker")]
struct Cli {
    #[command(subcommand)]
    command: Commands,
}

#[derive(Subcommand)]
enum Commands {
    Add { description: String },
    List { status: Option<String> },
    Update { id: u32, status: String },
    Done { id: u32 },
    Delete { id: u32 },
}
```
</details>

<details>
<summary>Go: Use flag package or cobra</summary>

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        printUsage()
        return
    }

    command := os.Args[1]

    switch command {
    case "add":
        if len(os.Args) < 3 {
            fmt.Println("Error: Missing description")
            return
        }
        addTask(os.Args[2])
    case "list":
        filter := ""
        if len(os.Args) > 2 {
            filter = os.Args[2]
        }
        listTasks(filter)
    // ... other commands
    }
}
```
</details>

<details>
<summary>Java: Use picocli</summary>

```xml
<dependency>
    <groupId>info.picocli</groupId>
    <artifactId>picocli</artifactId>
    <version>4.7.4</version>
</dependency>
```

```java
import picocli.CommandLine;
import picocli.CommandLine.Command;

@Command(name = "task", mixinStandardHelpOptions = true)
public class TaskCLI implements Runnable {

    @Command(name = "add")
    public void add(@Parameters(index = "0") String description) {
        // Implementation
    }

    @Command(name = "list")
    public void list(@Parameters(index = "0", defaultValue = "") String filter) {
        // Implementation
    }

    public static void main(String[] args) {
        int exitCode = new CommandLine(new TaskCLI()).execute(args);
        System.exit(exitCode);
    }
}
```
</details>

### Phase 5: Testing

Write tests for:
- Task creation
- Status updates
- File I/O (use temporary files)
- Error handling

<details>
<summary>Rust Testing</summary>

```rust
#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_add_task() {
        let mut store = TaskStore {
            tasks: Vec::new(),
            next_id: 1,
        };

        let task = store.add_task("Test task".to_string());
        assert_eq!(task.id, 1);
        assert_eq!(task.description, "Test task");
        assert_eq!(task.status, TaskStatus::Todo);
    }

    #[test]
    fn test_update_task_status() {
        // ...
    }
}
```

Run tests:
```bash
cargo test
```
</details>

<details>
<summary>Go Testing</summary>

```go
package main

import (
    "testing"
)

func TestAddTask(t *testing.T) {
    store := &TaskStore{
        Tasks:  []Task{},
        NextID: 1,
    }

    task := store.AddTask("Test task")
    if task.ID != 1 {
        t.Errorf("Expected ID 1, got %d", task.ID)
    }
    if task.Description != "Test task" {
        t.Errorf("Expected 'Test task', got '%s'", task.Description)
    }
}
```

Run tests:
```bash
go test
```
</details>

<details>
<summary>Java Testing (JUnit 5)</summary>

```java
import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;

class TaskStoreTest {

    @Test
    void testAddTask() {
        TaskStore store = new TaskStore();
        Task task = store.addTask("Test task");

        assertEquals(1, task.getId());
        assertEquals("Test task", task.getDescription());
        assertEquals(TaskStatus.TODO, task.getStatus());
    }
}
```

Run tests:
```bash
mvn test
```
</details>

---

## Enhancement Ideas

Once you've completed the basic version, try these extensions:

### Level 2: Intermediate

1. **Priority Levels**: Add high/medium/low priority
2. **Due Dates**: Set and display due dates
3. **Search**: Find tasks by keyword
4. **Color Output**: Colorize status in terminal
5. **Task Categories**: Group tasks by project/category

### Level 3: Advanced

1. **Recurring Tasks**: Support daily/weekly tasks
2. **Subtasks**: Nested tasks
3. **Task Dependencies**: Block tasks until dependencies complete
4. **Export/Import**: CSV, Markdown formats
5. **Sync**: Cloud storage (S3, Google Drive)

### Level 4: Expert

1. **Web Interface**: REST API + frontend
2. **Multi-user**: User accounts and permissions
3. **Real-time Sync**: WebSocket updates
4. **Mobile App**: React Native or Flutter
5. **AI Assistant**: Natural language task creation

---

## Learning Checkpoints

After completing this project, you should be able to answer:

1. **How does each language handle JSON serialization?**
   - Rust: serde traits
   - Go: struct tags
   - Java: Jackson annotations

2. **How does error handling differ?**
   - Rust: `Result<T, E>` and `?` operator
   - Go: `error` return value
   - Java: Exceptions

3. **How are timestamps represented?**
   - Rust: `chrono::DateTime<Utc>`
   - Go: `time.Time`
   - Java: `java.time.Instant`

4. **How are command-line args parsed?**
   - Rust: clap (derive macros)
   - Go: flag or cobra
   - Java: picocli (annotations)

---

## Project Structure

```
01-task-tracker/
├── rust/
│   ├── Cargo.toml
│   ├── src/
│   │   ├── main.rs
│   │   ├── task.rs
│   │   └── storage.rs
│   └── tests/
│       └── integration_test.rs
├── go/
│   ├── go.mod
│   ├── main.go
│   ├── task.go
│   ├── storage.go
│   └── main_test.go
├── java/
│   ├── pom.xml
│   └── src/
│       ├── main/
│       │   └── java/
│       │       └── tasktracker/
│       │           ├── Main.java
│       │           ├── Task.java
│       │           └── TaskStore.java
│       └── test/
│           └── java/
│               └── tasktracker/
│                   └── TaskStoreTest.java
└── README.md (this file)
```

---

## Comparison Matrix

After implementing in all three languages, compare:

| Aspect | Rust | Go | Java |
|--------|------|-----|------|
| **Lines of Code** | ? | ? | ? |
| **Binary Size** | ? | ? | N/A (JAR) |
| **Compilation Time** | ? | ? | ? |
| **Runtime Performance** | ? | ? | ? |
| **Error Handling Style** | Result | error return | Exceptions |
| **JSON Library** | serde_json | encoding/json | Jackson |
| **CLI Parsing** | clap | flag/cobra | picocli |
| **Testing Framework** | Built-in | Built-in | JUnit |

---

## Resources

### Rust
- [Serde JSON Docs](https://docs.serde.rs/serde_json/)
- [Clap Documentation](https://docs.rs/clap/)
- [Rust Book: Error Handling](https://doc.rust-lang.org/book/ch09-00-error-handling.html)

### Go
- [JSON and Go](https://go.dev/blog/json)
- [Effective Go: Errors](https://go.dev/doc/effective_go#errors)
- [Testing in Go](https://go.dev/doc/tutorial/add-a-test)

### Java
- [Jackson Tutorial](https://www.baeldung.com/jackson)
- [Picocli Quick Start](https://picocli.info/#_quick_start)
- [JUnit 5 User Guide](https://junit.org/junit5/docs/current/user-guide/)

---

## Submission Checklist

- [ ] Core CRUD operations work
- [ ] Tasks persist to JSON file
- [ ] Error handling for invalid inputs
- [ ] At least 5 unit tests
- [ ] README with build/run instructions
- [ ] Code is well-commented
- [ ] Follows language conventions (rustfmt, gofmt, checkstyle)

---

**Next Project:** [02-web-server](../02-web-server/) - Build a concurrent HTTP server
