package swap_sort

import (
	"awesomeProject/sort/memery_sort/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := util.MockData(10000)
	result := util.SortData(arr)
	BubbleSort(arr)
	assert.Equal(t, result, arr)
}
