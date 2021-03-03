package swap_sort

func BubbleSort(arr []int64) []int64 {
	for i:=len(arr)-1;i>=0;i--{
		swap := false
		for j:=0;j<i;j++{
			if arr[j]>arr[j+1]{
				swap=true
				tmp := arr[j]
				arr[j]=arr[j+1]
				arr[j+1]=tmp
			}
		}
		if !swap {
			break
		}
	}
	return arr
}