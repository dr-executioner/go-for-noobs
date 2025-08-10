# 📚 Go Type System

This folder covers Go's type system beyond basic data types, focusing on how you can create your own types and work with them.

---

## 📖 Type Declarations

You can create **new, distinct types** based on existing types using the `type` keyword.

```go
type Name string
type Age int

```

## 📖 Type Aliases

A type alias gives an existing type a new name. It behaves exactly like the original type.

```go
type ID = int
```

## 📖 Method Sets on Custom Types

If you create a new type (not an alias), you can attach methods to it.

```go
type Celsius float64

func (c Celsius) ToFahrenheit() float64 {
	return float64(c)*9/5 + 32
}
```

##📖 Zero Values for Custom Types

A custom type inherits the zero value of its underlying type.
```go
type Age int

var a Age // Zero value is 0
```

# 📚 Go Type Assertions & Type Switches

Go uses interfaces to express general behavior, but sometimes you need to check what the actual concrete type is inside an `interface{}`.  
This is where **type assertions** and **type switches** come in.

---

## 📖 Type Assertions

A **type assertion** is used to extract the underlying concrete value from an interface.

### ✅ Syntax

```go
value, ok := x.(T)
```

x → interface value

T → type you want to check for

ok → boolean result (true if it matches, false if it doesn’t)

# 📚 Go Interfaces

An **interface** in Go defines a set of method signatures (behavior), and any type that implements those methods is said to satisfy the interface — no explicit declaration needed.

Go uses interfaces for polymorphism, abstraction, and clean dependency injection.

---

## 📖 Basic Interface Syntax

```go
type Shape interface {
	Area() float64
}
```