package select_sort

import (
	"sort/memery_sort/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStraightInsertSort(t *testing.T) {
	arr := util.MockData(1000)
	result := util.SortData(arr)
	SimpleSelectionSort(arr)
	assert.Equal(t, result, arr)
}

func TestHeadSort(t *testing.T) {
	arr := util.MockData(1000)
	result := util.SortData(arr)
	arr = HeapSort(arr)
	assert.Equal(t, result, arr)
}
