package memery_sort

import (
	"awesomeProject/sort/memery_sort/insert_sort"
	"awesomeProject/sort/memery_sort/select_sort"
	"awesomeProject/sort/memery_sort/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	MaxCount = 10000
)

var (
	arr []int64
	expectedResult []int64
)

func init() {
	arr = util.MockData(MaxCount)
	expectedResult = util.SortData(arr)
}

func BenchmarkStraightInsertSort(b *testing.B){
	for i:=0;i<b.N;i++{
		result := insert_sort.StraightInsertSort(util.CopyData(arr))
		assert.Equal(b, expectedResult, result)
	}
}

func BenchmarkSimpleSelectionSort(b *testing.B){
	for i:=0;i<b.N;i++{
		result := select_sort.SimpleSelectionSort(util.SortData(arr))
		assert.Equal(b, expectedResult, result)
	}
}
