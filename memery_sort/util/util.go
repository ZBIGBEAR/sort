package util

import (
	"math/rand"
	"sort"
)

func MockData(n int) []int64 {
	var arr []int64
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Int63n(1000))
	}
	return arr
}

func SortData(arr []int64) []int64 {
	copyArr := make([]int64, len(arr))
	for i := range arr {
		copyArr[i] = arr[i]
	}
	sort.Sort(Int64Arr(copyArr))
	return copyArr
}

func CopyData(arr []int64) []int64 {
	copyData := make([]int64, len(arr))
	for i := range arr {
		copyData[i] = arr[i]
	}
	return copyData
}

type Int64Arr []int64

func (arr Int64Arr) Less(i, j int) bool {
	return arr[i] < arr[j]
}

func (arr Int64Arr) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (arr Int64Arr) Len() int {
	return len(arr)
}
