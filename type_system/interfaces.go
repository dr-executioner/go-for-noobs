package typesystem

// Multiple Interfaces
type Writer interface {
	Write(p []byte) (n int, err error)
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}


// Empty interface equivalent to any
var x interface{} = "hello"


// Attatching a method to a interface
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

//When the value inside the interface is nil or no concrete type, calling methods on it panics.