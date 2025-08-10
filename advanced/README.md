# 📚 Go Context

The `context` package in Go helps manage:
- Goroutine lifecycle
- Cancellation
- Deadlines and timeouts
- Request-scoped values

---

## 📖 Why Use Context?

✅ Avoid leaking goroutines  
✅ Allow timeouts on operations (e.g., database, HTTP calls)  
✅ Cleanly propagate cancellation signals  
✅ Pass metadata/values across function calls safely  

---

## 📖 Import

```go
import "context"
