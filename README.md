# Go 堆排序 

# Install 

```bash
go get github.com/CC11001100/go-heap
```

# Example
```go
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

	// 非线程安全的堆，仅限于单个goroutine里使用
	//heap :=heap.New(heap.IntComparator())
	heap := heap.NewWithOptions(options)

	// 创建线程安全的堆
	//heap :=heap.NewSync(heap.IntComparator())
	//heap := heap.NewSyncWithOptions(options)
	for i := 0; i < 10; i++ {
		n := rand.Int() % 100
		heap.Push(n)
	}
	heapNumSlice := heap.PopToSlice()
	fmt.Println(heapNumSlice) // [10 16 20 21 37 48 49 51 51 58]
}
```

