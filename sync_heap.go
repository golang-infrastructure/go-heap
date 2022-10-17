package heap

import "sync"

// SyncHeap 线程安全的堆
type SyncHeap[T any] struct {
	heap *Heap[T]
	lock sync.RWMutex
}

var _ Interface[any] = &SyncHeap[any]{}

func NewSync[T any](comparator Comparator[T]) *SyncHeap[T] {
	return NewSyncWithOptions(&Options[T]{
		Comparator: comparator,
	})
}

func NewSyncWithOptions[T any](options *Options[T]) *SyncHeap[T] {
	return &SyncHeap[T]{
		heap: NewWithOptions(options),
		lock: sync.RWMutex{},
	}
}

func (x *SyncHeap[T]) Push(valueSlice ...T) {
	x.lock.Lock()
	defer x.lock.Unlock()
	x.heap.Push(valueSlice...)
}

func (x *SyncHeap[T]) Clear() {
	x.lock.Lock()
	defer x.lock.Unlock()
	x.heap.Clear()
}

func (x *SyncHeap[T]) PeekE() (T, error) {
	x.lock.RLock()
	defer x.lock.RUnlock()
	return x.heap.PeekE()
}

func (x *SyncHeap[T]) Peek() T {
	x.lock.RLock()
	defer x.lock.RUnlock()
	return x.heap.Peek()
}

func (x *SyncHeap[T]) PopE() (T, error) {
	x.lock.Lock()
	defer x.lock.Unlock()
	return x.heap.PopE()
}

func (x *SyncHeap[T]) Pop() T {
	x.lock.Lock()
	defer x.lock.Unlock()
	return x.heap.Pop()
}

func (x *SyncHeap[T]) Size() int {
	x.lock.RLock()
	defer x.lock.RUnlock()
	return x.heap.Size()
}

func (x *SyncHeap[T]) IsEmpty() bool {
	x.lock.RLock()
	defer x.lock.RUnlock()
	return x.heap.IsEmpty()
}

func (x *SyncHeap[T]) IsNotEmpty() bool {
	x.lock.RLock()
	defer x.lock.RUnlock()
	return x.heap.IsNotEmpty()
}

func (x *SyncHeap[T]) PopTopN(n int) []T {
	x.lock.Lock()
	defer x.lock.Unlock()
	return x.heap.PopTopN(n)
}

func (x *SyncHeap[T]) PopEach(eachFunc func(v T) bool) {
	x.lock.Lock()
	defer x.lock.Unlock()
	x.heap.PopEach(eachFunc)
}
