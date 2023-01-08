package heap

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sort"
	"testing"
)

func TestHeap_Clear(t *testing.T) {

}

func TestHeap_IsEmpty(t *testing.T) {

}

func TestHeap_IsNotEmpty(t *testing.T) {

}

func TestHeap_Peek(t *testing.T) {

}

func TestHeap_PeekE(t *testing.T) {

}

func TestHeap_Pop(t *testing.T) {

}

func TestHeap_PopE(t *testing.T) {

}

func TestHeap_PopEach(t *testing.T) {

}

func TestHeap_PopToSlice(t *testing.T) {

}

func TestHeap_PopTopN(t *testing.T) {

}

func TestHeap_Push(t *testing.T) {
}

func TestHeap_Size(t *testing.T) {

}

func TestHeap_down(t *testing.T) {

}

func TestHeap_swap(t *testing.T) {

}

func TestHeap_up(t *testing.T) {

}

func TestNewWithComparator(t *testing.T) {

}

func TestNewWithOptions(t *testing.T) {
	options := &Options[int]{
		Comparator: IntComparator(),
		Ary:        4,
	}
	heap := NewWithOptions(options)
	numSlice := make([]int, 0)
	for i := 0; i < 100; i++ {
		n := rand.Int() % 100
		heap.Push(n)
		numSlice = append(numSlice, n)
	}
	heapNumSlice := heap.PopToSlice()
	sort.Ints(numSlice)
	assert.Equal(t, numSlice, heapNumSlice)
}

func TestHeap_ExportDotLanguage(t *testing.T) {
	options := &Options[int]{
		Comparator: IntComparator(),
		Ary:        10,
	}
	heap := NewWithOptions(options)
	for i := 0; i < 20; i++ {
		n := rand.Int() % 1000
		heap.Push(n)
	}

	for heap.IsNotEmpty() {
		fmt.Println(heap.ExportDotLanguage())
		fmt.Println("-------------------------------------------------------------------------")
		heap.Pop()
	}

}

func TestHeap_MarshalJSON(t *testing.T) {

	options := &Options[int]{
		Comparator: IntComparator(),
		Ary:        10,
	}
	heap := NewWithOptions(options)
	for i := 0; i < 20; i++ {
		n := rand.Int() % 1000
		heap.Push(n)
	}

	t.Log(heap.MarshalJSON())

}
