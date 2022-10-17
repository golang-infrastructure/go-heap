package heap

// DefaultAryHeap 默认是2叉堆
const DefaultAryHeap = 2

// Options 堆排序的选项
type Options[T any] struct {

	// 必选项，比较器，用来决定堆中数据之前的相对大小
	Comparator Comparator[T]

	// 可选项，堆是几叉堆，默认是二叉堆
	Ary int

	// 可选项，可以从指定的数组初始化堆，不指定的话默认为空堆
	InitSlice []T
}
