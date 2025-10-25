# Project 2: Concurrent Web Server

## Overview

Build a production-ready HTTP web server that demonstrates concurrency, networking, and request handling. This project will teach you how each language approaches concurrent I/O and server architecture.

**Difficulty:** Intermediate
**Time Estimate:** 8-12 hours
**Concepts Covered:**
- HTTP protocol fundamentals
- Concurrent request handling
- Routing and middleware
- Static file serving
- JSON APIs
- Error handling in async contexts

---

## Project Goals

1. **Master concurrency patterns**: Handle multiple requests simultaneously
2. **Understand HTTP**: Parse requests, build responses
3. **Build REST APIs**: Create JSON endpoints
4. **Serve static files**: HTML, CSS, JavaScript, images
5. **Implement middleware**: Logging, authentication, CORS
6. **Performance testing**: Load testing and optimization

---

## Functional Requirements

### Basic Features

```bash
# Health check endpoint
GET /health
Response: 200 OK { "status": "healthy" }

# Hello world
GET /
Response: 200 OK "Hello, World!"

# Echo endpoint
POST /echo
Body: { "message": "hello" }
Response: 200 OK { "message": "hello" }

# Serve static files
GET /static/index.html
Response: HTML content

# RESTful API
GET    /api/users       # List users
GET    /api/users/:id   # Get user by ID
POST   /api/users       # Create user
PUT    /api/users/:id   # Update user
DELETE /api/users/:id   # Delete user
```

### Requirements

- Handle 1000+ concurrent connections
- Support GET, POST, PUT, DELETE methods
- Parse JSON request bodies
- Return JSON responses
- Serve static files with correct MIME types
- Request logging middleware
- Graceful shutdown
- Configuration via environment variables

---

## Architecture

### Request Flow

```
Client Request
    │
    ▼
┌─────────────────┐
│  TCP Listener   │  Listen on port (e.g., 8080)
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│   Connection    │  Accept connection
│     Handler     │  (spawns new thread/goroutine/async task)
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│   Middleware    │  Logging, auth, CORS
│     Chain       │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│     Router      │  Match route pattern
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│    Handler      │  Business logic
│    Function     │
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│    Response     │  Status code, headers, body
└─────────────────┘
```

---

## Implementation Guide

### Phase 1: Simple HTTP Server

Start with a basic server that responds to all requests.

<details>
<summary>Rust (using tokio + hyper)</summary>

```toml
[dependencies]
tokio = { version = "1", features = ["full"] }
hyper = { version = "0.14", features = ["full"] }
```

```rust
use hyper::service::{make_service_fn, service_fn};
use hyper::{Body, Request, Response, Server, StatusCode};
use std::convert::Infallible;
use std::net::SocketAddr;

async fn handle_request(req: Request<Body>) -> Result<Response<Body>, Infallible> {
    match (req.method(), req.uri().path()) {
        (&hyper::Method::GET, "/") => {
            Ok(Response::new(Body::from("Hello, World!")))
        }
        (&hyper::Method::GET, "/health") => {
            Ok(Response::new(Body::from(r#"{"status":"healthy"}"#)))
        }
        _ => {
            let mut response = Response::new(Body::from("Not Found"));
            *response.status_mut() = StatusCode::NOT_FOUND;
            Ok(response)
        }
    }
}

#[tokio::main]
async fn main() {
    let addr = SocketAddr::from(([127, 0, 0, 1], 8080));

    let make_svc = make_service_fn(|_conn| async {
        Ok::<_, Infallible>(service_fn(handle_request))
    });

    let server = Server::bind(&addr).serve(make_svc);

    println!("Server running on http://{}", addr);

    if let Err(e) = server.await {
        eprintln!("Server error: {}", e);
    }
}
```
</details>

<details>
<summary>Go (using net/http)</summary>

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, World!"))
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "status": "healthy",
    })
}

func main() {
    http.HandleFunc("/", handleRoot)
    http.HandleFunc("/health", handleHealth)

    addr := ":8080"
    fmt.Printf("Server running on http://localhost%s\n", addr)
    log.Fatal(http.ListenAndServe(addr, nil))
}
```
</details>

<details>
<summary>Java (using Spring Boot)</summary>

```xml
<dependency>
    <groupId>org.springframework.boot</groupId>
    <artifactId>spring-boot-starter-web</artifactId>
    <version>3.1.0</version>
</dependency>
```

```java
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.*;

@SpringBootApplication
@RestController
public class WebServerApplication {

    @GetMapping("/")
    public String root() {
        return "Hello, World!";
    }

    @GetMapping("/health")
    public Map<String, String> health() {
        return Map.of("status", "healthy");
    }

    public static void main(String[] args) {
        SpringApplication.run(WebServerApplication.class, args);
    }
}
```
</details>

### Phase 2: RESTful API

Implement CRUD operations for a User resource.

**User Model:**
```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "createdAt": "2025-01-15T10:30:00Z"
}
```

**Endpoints:**
```
GET    /api/users       - List all users
GET    /api/users/:id   - Get user by ID
POST   /api/users       - Create new user (body: name, email)
PUT    /api/users/:id   - Update user (body: name, email)
DELETE /api/users/:id   - Delete user
```

**Implementation Tips:**
- Use in-memory HashMap/Dict for storage (Phase 3 will add database)
- Return proper HTTP status codes (200, 201, 404, 400)
- Validate input (email format, required fields)
- Use UUID or auto-incrementing IDs

### Phase 3: Middleware

Implement middleware for:

1. **Request Logging**
   ```
   [2025-01-15 10:30:45] GET /api/users - 200 OK (15ms)
   ```

2. **CORS Headers**
   ```
   Access-Control-Allow-Origin: *
   Access-Control-Allow-Methods: GET, POST, PUT, DELETE
   ```

3. **Authentication** (Simple token-based)
   ```
   Authorization: Bearer <token>
   ```

<details>
<summary>Rust Middleware Example</summary>

```rust
use std::time::Instant;

async fn logging_middleware(
    req: Request<Body>,
    next: impl Fn(Request<Body>) -> Response<Body>,
) -> Response<Body> {
    let method = req.method().clone();
    let path = req.uri().path().to_string();
    let start = Instant::now();

    let response = next(req);

    let elapsed = start.elapsed();
    println!("[{}] {} {} - {:?}",
        chrono::Utc::now().format("%Y-%m-%d %H:%M:%S"),
        method,
        path,
        elapsed
    );

    response
}
```
</details>

### Phase 4: Static File Serving

Serve files from a `public/` directory:

```
public/
├── index.html
├── style.css
├── app.js
└── images/
    └── logo.png
```

**Requirements:**
- Detect MIME type from file extension
- Handle 404 for missing files
- Security: Prevent directory traversal (`../etc/passwd`)
- Caching headers for performance

### Phase 5: Performance & Testing

**Load Testing:**
```bash
# Using Apache Bench
ab -n 10000 -c 100 http://localhost:8080/

# Using wrk
wrk -t12 -c400 -d30s http://localhost:8080/

# Using hey
hey -n 10000 -c 100 http://localhost:8080/
```

**Metrics to Track:**
- Requests per second (RPS)
- Average latency
- 99th percentile latency
- Error rate
- Memory usage

**Unit Tests:**
- Test each endpoint
- Test error cases (404, 400, 500)
- Test middleware execution order
- Test concurrent requests

---

## Enhancement Ideas

### Level 2: Intermediate

1. **WebSocket Support**: Real-time bidirectional communication
2. **File Uploads**: Handle multipart/form-data
3. **Rate Limiting**: Prevent abuse (per IP or per user)
4. **Compression**: Gzip response bodies
5. **HTTPS/TLS**: SSL certificate handling

### Level 3: Advanced

1. **Database Integration**: PostgreSQL, MySQL, or MongoDB
2. **Caching Layer**: Redis for session/data caching
3. **JWT Authentication**: Stateless auth with tokens
4. **API Versioning**: `/api/v1/users`, `/api/v2/users`
5. **OpenAPI/Swagger**: Auto-generate API documentation

### Level 4: Expert

1. **GraphQL Server**: Alternative to REST
2. **gRPC Support**: Binary protocol with protobuf
3. **Microservices**: Split into auth, user, order services
4. **Message Queue**: RabbitMQ or Kafka integration
5. **Observability**: Prometheus metrics, distributed tracing

---

## Performance Comparison

After implementing all three versions, benchmark them:

### Test Setup
- 10,000 requests
- 100 concurrent connections
- Mixed workload (70% reads, 30% writes)

### Expected Results (Approximate)

| Metric | Rust (Tokio) | Go (net/http) | Java (Spring Boot) |
|--------|--------------|---------------|---------------------|
| **RPS** | ~50,000 | ~40,000 | ~20,000 |
| **Avg Latency** | 2ms | 3ms | 5ms |
| **P99 Latency** | 10ms | 15ms | 50ms |
| **Memory** | 10 MB | 50 MB | 150 MB |
| **CPU** | 100% (1 core) | 150% (1.5 cores) | 200% (2 cores) |

**Note:** These are rough estimates. Actual performance depends on hardware, implementation details, and JVM tuning.

---

## Concurrency Models Explained

### Rust: Async/Await with Tokio

```
Single-threaded event loop (or thread pool)
┌────────────────────────────────┐
│         Tokio Runtime          │
│  ┌──────┐ ┌──────┐ ┌──────┐  │
│  │Task 1│ │Task 2│ │Task N│  │
│  └──────┘ └──────┘ └──────┘  │
│         Event Loop             │
└────────────────────────────────┘

- Async tasks yield at .await points
- Cooperative multitasking
- Very low overhead per connection
```

### Go: Goroutines

```
M:N threading (goroutines:OS threads)
┌────────────────────────────────┐
│     Go Runtime Scheduler       │
│  G1  G2  G3  G4 ... (10,000+) │
└──────┬───────┬──────────────────┘
       │       │
       ▼       ▼
    ┌─────┬─────┐
    │  M1 │  M2 │  (OS Threads)
    └─────┴─────┘

- One goroutine per connection
- Preemptive scheduling
- ~2 KB overhead per goroutine
```

### Java: Thread Pool

```
Fixed thread pool with request queue
┌────────────────────────────────┐
│   Tomcat Thread Pool (200)     │
│  T1  T2  T3 ... T200           │
└────────┬───────────────────────┘
         │
         ▼
┌────────────────────────────────┐
│      Request Queue             │
│  R1 → R2 → R3 → ... → RN      │
└────────────────────────────────┘

- One thread per request (blocking)
- Thread reuse via pool
- ~1 MB overhead per thread
- Virtual threads (Java 21+) change this!
```

---

## Common Pitfalls

1. **Not handling errors properly**
   - Always return appropriate status codes
   - Log errors for debugging
   - Don't expose internal errors to clients

2. **Blocking the event loop (Rust/Node.js)**
   - Use `tokio::task::spawn_blocking` for CPU-intensive work
   - Don't do heavy computation in async functions

3. **Memory leaks**
   - Go: Unbounded channels or goroutine leaks
   - Java: Not closing connections, excessive caching
   - Rust: Circular Rc references (rare)

4. **Improper shutdown**
   - Always implement graceful shutdown
   - Finish in-flight requests before exiting
   - Close database connections

---

## Project Structure

```
02-web-server/
├── rust/
│   ├── Cargo.toml
│   ├── src/
│   │   ├── main.rs
│   │   ├── routes/
│   │   │   ├── mod.rs
│   │   │   └── users.rs
│   │   ├── middleware/
│   │   │   ├── mod.rs
│   │   │   └── logging.rs
│   │   └── models/
│   │       └── user.rs
│   └── public/
│       └── index.html
├── go/
│   ├── go.mod
│   ├── main.go
│   ├── handlers/
│   │   └── users.go
│   ├── middleware/
│   │   └── logging.go
│   ├── models/
│   │   └── user.go
│   └── public/
│       └── index.html
└── java/
    ├── pom.xml
    └── src/main/
        ├── java/
        │   └── webserver/
        │       ├── WebServerApplication.java
        │       ├── controllers/
        │       │   └── UserController.java
        │       ├── models/
        │       │   └── User.java
        │       └── middleware/
        │           └── LoggingFilter.java
        └── resources/
            └── static/
                └── index.html
```

---

## Resources

### Rust
- [Tokio Tutorial](https://tokio.rs/tokio/tutorial)
- [Hyper Guide](https://hyper.rs/guides/)
- [Axum Framework](https://github.com/tokio-rs/axum) (higher-level)

### Go
- [Writing Web Applications](https://go.dev/doc/articles/wiki/)
- [net/http Package](https://pkg.go.dev/net/http)
- [Gin Framework](https://github.com/gin-gonic/gin) (higher-level)

### Java
- [Spring Boot Guide](https://spring.io/guides/gs/rest-service/)
- [Building REST APIs](https://www.baeldung.com/rest-with-spring-series)
- [Java HTTP Server (built-in)](https://docs.oracle.com/en/java/javase/17/docs/api/jdk.httpserver/com/sun/net/httpserver/HttpServer.html)

---

## Submission Checklist

- [ ] Server starts and listens on configurable port
- [ ] Health check endpoint works
- [ ] Full CRUD API for users
- [ ] Static file serving with correct MIME types
- [ ] Request logging middleware
- [ ] Graceful shutdown on SIGTERM/SIGINT
- [ ] Load test results documented
- [ ] At least 10 integration tests
- [ ] Comparison notes across languages

---

**Previous:** [01-task-tracker](../01-task-tracker/)
**Next:** [03-database](../03-database/) - Build a simple database engine
