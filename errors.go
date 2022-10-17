package heap

import "errors"

var (
	// ErrHeapIsEmpty 堆是空的时候不能进行某些操作，会返回此错误
	ErrHeapIsEmpty = errors.New("heap is empty")
)
