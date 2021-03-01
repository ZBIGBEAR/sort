package insert_sort

import (
	"awesomeProject/sort/memery_sort/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStraightInsertSort(t *testing.T) {
	arr := util.MockData(1000)
	result := util.SortData(arr)
	StraightInsertSort(arr)
	assert.Equal(t, result, arr)
}
