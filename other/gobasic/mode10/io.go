package main

type Reader interface {
	Read(p []byte) (n int, err error)
}
type Writer interface {
	Write(p []byte) (n int, err error)
}
type Closer interface {
	Close() error
}

// ReadWriteCloser 此接口包含三个接口（组合接口）
type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}

func main() {

}
