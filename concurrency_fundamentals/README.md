# 📚 Go Goroutines

**Goroutines** are lightweight, independently executing functions managed by the Go runtime.

✅ They allow you to run multiple tasks concurrently within the same program — like threads, but cheaper.

---

## 📖 What is a Goroutine?

A **goroutine** is a function that runs concurrently with other functions.

They’re:
- Lightweight (a few KBs of stack memory)
- Multiplexed onto OS threads by Go's scheduler
- Created using the `go` keyword

---

## 📖 ✅ Syntax

```go
go functionName()
```

# 📚 Concurrency vs Parallelism in Go

Understanding **concurrency** and **parallelism** is key to writing efficient, scalable, and clean Go applications.

Go was designed for **concurrency first**, but can achieve **parallelism** too.

---

## 📖 What is Concurrency?

> **Concurrency is about dealing with many things at once.**

It’s a way to structure your program so multiple tasks can progress independently — even if not simultaneously.

✅ Go uses **goroutines** and its **scheduler** to manage concurrency.

---

## 📖 What is Parallelism?

> **Parallelism is about doing many things at the exact same time.**

It requires multiple CPU cores and physically runs different tasks simultaneously.

✅ Go can do this when multiple CPU cores are available and when `GOMAXPROCS` is set.

---

## 📦 Real-World Analogy

| Concept      | Analogy                                                 |
|:---------------|:----------------------------------------------------------|
| **Concurrency** | A single chef cooking multiple dishes by switching tasks quickly |
| **Parallelism** | Multiple chefs cooking different dishes at the same time |

---

## 📖 Go Code Example

### Concurrent (but not necessarily parallel)

```go
go task1()
go task2()
// The Go scheduler switches between them as needed
```

# 📚 Go Channels

**Channels** are Go’s built-in concurrency-safe communication mechanism.  
They let you safely pass data between goroutines without needing locks or shared memory.

---

## 📖 What is a Channel?

A **channel** is a typed conduit through which you can send and receive values between goroutines.

It guarantees **synchronization and safety** — when one goroutine sends data into a channel, it’s blocked until another goroutine receives it, and vice versa (unless it’s a buffered channel).

---

## 📖 ✅ Channel Syntax  

### 📦 Declare a Channel  

```go
var ch chan int
```

## 📖 ✅ Buffered Channels
A buffered channel allows a limited number of values to be queued without blocking.
```go
ch := make(chan int, 3)
```