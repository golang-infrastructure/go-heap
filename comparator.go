package heap

import "strings"

// Comparator 用于比较堆中元素的大小
type Comparator[T any] func(a T, b T) int

func StringComparator() Comparator[string] {
	return strings.Compare
}

func IntComparator() Comparator[int] {
	return func(a int, b int) int {
		return a - b
	}
}

func UintComparator() Comparator[uint] {
	return func(a uint, b uint) int {
		return int(a) - int(b)
	}
}

func Float64Comparator() Comparator[float64] {
	return func(a float64, b float64) int {
		return int((a - b) * 1000)
	}
}

func Float32Comparator() Comparator[float32] {
	return func(a float32, b float32) int {
		return int((a - b) * 1000)
	}
}
