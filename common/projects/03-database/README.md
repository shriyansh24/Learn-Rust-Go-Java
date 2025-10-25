# Project 3: Key-Value Database Engine

## Overview

Build a simple but functional key-value database with persistence, indexing, and query capabilities. Learn about data structures, file formats, and database fundamentals.

**Difficulty:** Advanced
**Time Estimate:** 16-24 hours
**Concepts Covered:** B-trees, WAL, ACID properties, serialization, indexing, transactions

---

## Core Features

```bash
# Set a key-value pair
db.set("user:1", {"name": "Alice", "age": 30})

# Get a value
db.get("user:1")  # Returns: {"name": "Alice", "age": 30}

# Delete a key
db.delete("user:1")

# Check if key exists
db.exists("user:1")  # Returns: false

# List keys with prefix
db.keys("user:")  # Returns: ["user:2", "user:3", ...]

# Transaction support
db.begin_transaction()
db.set("balance", 100)
db.set("pending", 50)
db.commit()  # or db.rollback()
```

---

## Architecture

### Storage Engine

```
┌─────────────────────────────────┐
│      Application API            │
├─────────────────────────────────┤
│      Transaction Layer          │  ACID guarantees
├─────────────────────────────────┤
│      Index Layer (B-Tree)       │  Fast lookups
├─────────────────────────────────┤
│      Storage Engine             │
│  ┌──────────┬──────────────┐   │
│  │   WAL    │  Data File   │   │
│  └──────────┴──────────────┘   │
└─────────────────────────────────┘
```

**Components:**
1. **Write-Ahead Log (WAL)**: Durability & crash recovery
2. **Data File**: Sorted string table (SSTable) or B-tree
3. **Index**: In-memory hash map or B-tree for fast lookups
4. **Memtable**: In-memory buffer for writes
5. **Compaction**: Merge/cleanup data files

---

## Implementation Phases

### Phase 1: In-Memory Store
Simple HashMap-based storage (no persistence).

### Phase 2: File Persistence
Append-only log with sequential reads.

### Phase 3: Indexing
Build in-memory index for O(1) lookups.

### Phase 4: Write-Ahead Log
Ensure durability with WAL.

### Phase 5: Transactions
Support ACID transactions with isolation.

### Phase 6: Compaction
Remove deleted entries, merge files.

---

## File Format Example

**WAL Entry:**
```
[Timestamp][Operation][Key Length][Key][Value Length][Value][Checksum]
```

**SSTable Format:**
```
Header:
  - Magic number: 0xDB01
  - Version: 1
  - Number of entries: 1000
  - Index offset: 8192

Data Blocks:
  [Key1][Value1]
  [Key2][Value2]
  ...

Index:
  [Key1 offset][Key2 offset]...
```

---

## Enhancement Ideas

- Range queries: `db.range("user:1", "user:999")`
- TTL (Time-To-Live): `db.set("session", data, ttl=3600)`
- Compression: LZ4/Snappy for values
- Replication: Master-slave architecture
- Distributed: Consistent hashing + Raft consensus

---

## Resources

- [Log-Structured Merge Trees](https://en.wikipedia.org/wiki/Log-structured_merge-tree)
- [B-tree Visualization](https://www.cs.usfca.edu/~galles/visualization/BTree.html)
- [Designing Data-Intensive Applications](https://dataintensive.net/)

---

**Next:** [04-compiler](../04-compiler/)
