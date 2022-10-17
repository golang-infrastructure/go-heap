package heap

// Interface 用于定义堆提供的API
type Interface[T any] interface {
	Clear()

	Push(v ...T)

	PeekE() (T, error)
	Peek() T

	PopE() (T, error)
	Pop() T

	Size() int

	IsEmpty() bool
	IsNotEmpty() bool

	PopTopN(n int) []T
	PopEach(eachFunc func(v T) bool)
}
