package swap_sort

import (
	"awesomeProject/sort/memery_sort/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFastSort(t *testing.T) {
	arr := util.MockData(10000)
	result := util.SortData(arr)
	FastSort(arr)
	assert.Equal(t, result, arr)
}
