package heap

import (
	"fmt"
	"strings"
)

type Heap[T any] struct {

	// 用于存储完全二叉树的数组
	heapSlice []T
	heapSize  int

	// 相关排序选项
	options *Options[T]
}

var _ Interface[any] = &Heap[any]{}

func New[T any](comparator Comparator[T]) *Heap[T] {
	return NewWithOptions[T](&Options[T]{
		Comparator: comparator,
		Ary:        DefaultAryHeap,
	})
}

func NewWithOptions[T any](options *Options[T]) *Heap[T] {

	// 最小必须得是二叉堆
	if options.Ary <= 1 {
		options.Ary = DefaultAryHeap
	}

	heap := &Heap[T]{
		options:   options,
		heapSlice: make([]T, 0),
		heapSize:  0,
	}

	if len(options.InitSlice) != 0 {
		heap.Push(options.InitSlice...)
	}

	return heap
}

func (x *Heap[T]) Push(valueSlice ...T) {
	for _, value := range valueSlice {
		if x.heapSize < len(x.heapSlice) {
			x.heapSlice[x.heapSize] = value
		} else {
			x.heapSlice = append(x.heapSlice, value)
		}
		x.heapSize++
		x.up(x.Size() - 1)
	}
}

// Clear 把堆清空
func (x *Heap[T]) Clear() {
	x.heapSlice = nil
}

// PeekE 获取堆顶的元素，但并不弹出
func (x *Heap[T]) PeekE() (T, error) {
	if x.IsEmpty() {
		var zero T
		return zero, ErrHeapIsEmpty
	} else {
		return x.heapSlice[0], nil
	}
}

func (x *Heap[T]) Peek() T {
	v, _ := x.PeekE()
	return v
}

// PopE 弹出堆顶的元素，如果堆是空的，则返回error
func (x *Heap[T]) PopE() (T, error) {
	if x.IsEmpty() {
		var zero T
		return zero, ErrHeapIsEmpty
	}
	v := x.heapSlice[0]
	x.swap(0, x.Size()-1)
	x.heapSize--
	x.down(0, x.Size())
	return v, nil
}

func (x *Heap[T]) Pop() T {
	v, _ := x.PopE()
	return v
}

func (x *Heap[T]) Size() int {
	return x.heapSize
}

func (x *Heap[T]) IsEmpty() bool {
	return x.Size() == 0
}

// IsNotEmpty 堆是否非空
func (x *Heap[T]) IsNotEmpty() bool {
	return x.Size() != 0
}

// PopTopN n为正数时将前N个元素弹出，-1表示全部
func (x *Heap[T]) PopTopN(n int) []T {
	resultSlice := make([]T, 0)
	for n > 0 && x.IsNotEmpty() {
		resultSlice = append(resultSlice, x.Pop())
		n--
	}
	return resultSlice
}

func (x *Heap[T]) PopToSlice() []T {
	return x.PopTopN(x.Size())
}

func (x *Heap[T]) PopEach(eachFunc func(v T) bool) {
	for x.IsNotEmpty() {
		if !eachFunc(x.Pop()) {
			return
		}
	}
}

// ------------------------------------------------ ---------------------------------------------------------------------

// 把站定下表的节点往上提
func (x *Heap[T]) up(index int) {
	for {
		parentIndex := (index - 1) / x.options.Ary
		if parentIndex == index || x.options.Comparator(x.heapSlice[parentIndex], x.heapSlice[index]) < 0 {
			break
		}
		x.swap(parentIndex, index)
		index = parentIndex
	}
}

// 把给定下标的节点往下沉
// TODO 2022-10-18 23:52:37 考虑下表int溢出的问题
func (x *Heap[T]) down(parentIndex, n int) bool {
	currentParentIndex := parentIndex
	for {
		// 把自己当做是最小的
		minChildIndex := currentParentIndex
		// 然后看看孩子节点中能不能找得到比自己更小的
		for i := 0; i < x.options.Ary; i++ {
			childIndex := x.options.Ary*currentParentIndex + i + 1
			// 到达完全二叉树的最后一个节点了，后面就没必要再继续了
			if childIndex >= n {
				break
			}
			if x.options.Comparator(x.heapSlice[childIndex], x.heapSlice[minChildIndex]) < 0 {
				minChildIndex = childIndex
			}
		}

		if minChildIndex == -1 || minChildIndex == currentParentIndex {
			break
		}

		x.swap(currentParentIndex, minChildIndex)
		currentParentIndex = minChildIndex
	}
	return currentParentIndex > parentIndex
}

// 交换两个下标位置的元素
func (x *Heap[T]) swap(i, j int) {
	x.heapSlice[i], x.heapSlice[j] = x.heapSlice[j], x.heapSlice[i]
}

// ------------------------------------------------ ---------------------------------------------------------------------

// ExportDotLanguage 导出为dot language，拿出去画图以便可视化
func (x *Heap[T]) ExportDotLanguage() string {
	// digraph G1 {
	//    a -- b;
	//    a -- d;
	//    b -- c;
	//    d -- c;
	//}
	sb := strings.Builder{}
	sb.WriteString("digraph G1 { \n")
	for i := 0; i < x.Size(); i++ {
		parentIndex := i
		for j := 0; j < x.options.Ary; j++ {
			childIndex := parentIndex*x.options.Ary + j + 1
			if childIndex >= x.Size() || childIndex == parentIndex {
				continue
			}
			sb.WriteString(fmt.Sprintf("    \"%d:%v\" -> \"%d:%v\"; \n", parentIndex, x.heapSlice[parentIndex], childIndex, x.heapSlice[childIndex]))
		}
	}
	sb.WriteString("}")
	return sb.String()
}

//func (x *Heap[T]) Show() {
//
//}

// ------------------------------------------------ ---------------------------------------------------------------------
