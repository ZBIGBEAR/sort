/*
基本思想：
父节点比左右两个孩子都要小/大
*/

package select_sort

func HeapSort(data []int64) []int64 {
	count := len(data)
	// 生成堆
	for i := count / 2; i >= 0; i-- {
		sink(data, i, count)
	}

	// 输出
	for i := range data {
		// 对顶元素与最后一个元素交换
		data[0], data[count-1-i] = data[count-1-i], data[0]
		// 对长度减1
		// 重新调整堆
		sink(data, 0, count-1-i)
	}

	return data
}

// 下沉数组[0,length]的第idx个元素。每次比较父节点和2个孩子节点，找出最大值放在父节点。如果找到了则要递归往下找
func sink(data []int64, idx, length int) {
	parent := idx // 指向要下沉的元素
	leftIdx := idx*2 + 1
	rightIdx := idx*2 + 2
	if leftIdx < length && data[leftIdx] > data[parent] {
		parent = leftIdx
	}

	if rightIdx < length && data[rightIdx] > data[parent] {
		parent = rightIdx
	}

	if parent != idx {
		// 交换
		data[idx], data[parent] = data[parent], data[idx]
		// 递归下沉
		sink(data, parent, length)
	}
}
