package main

import (
	"fmt"
	"github.com/CC11001100/go-heap"
	"math/rand"
)

func main() {
	options := &heap.Options[int]{
		Comparator: heap.IntComparator(),
		// 支持N叉堆，默认是2叉堆
		Ary: 4,
	}
	heap := heap.NewWithOptions(options)
	for i := 0; i < 10; i++ {
		n := rand.Int() % 100
		heap.Push(n)
	}
	heapNumSlice := heap.PopToSlice()
	fmt.Println(heapNumSlice) // [10 16 20 21 37 48 49 51 51 58]
}
