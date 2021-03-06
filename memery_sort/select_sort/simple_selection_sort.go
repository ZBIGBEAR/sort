/*
基本思想：
在要排序的一组数中，选出最小（或者最大）的一个数与第1个位置的数交换；
然后在剩下的数当中再找最小（或者最大）的与第2个位置的数交换，
依次类推，直到第n-1个元素（倒数第二个数）和第n个元素（最后一个数）比较为止。
*/

package select_sort

func SimpleSelectionSort(arr []int64) []int64 {
	for i:=0;i<len(arr)-1;i++{
		minIndex := i
		for j:=i+1;j<len(arr);j++{
			if arr[minIndex]>arr[j]{
				minIndex=j
			}
		}
		if minIndex>=len(arr){
			break
		}
		arr[minIndex],arr[i]=arr[i],arr[minIndex]
	}
	return arr
}