package swap_sort

func FastSort(arr [] int64) []int64 {
	fastSort(arr,0,len(arr)-1)
	return arr
}

func fastSort(arr []int64, begin,end int) {
	if begin>=end{
		return
	}
	midIndex := begin
	i,j := begin+1, end
	for i<=j{
		for j>=i && arr[j]>=arr[midIndex] {
			j--
		}
		if j<i{
			break
		}
		arr[midIndex],arr[j]=arr[j],arr[midIndex]
		midIndex=j

		for i<=j && arr[i]<=arr[midIndex] {
			i++
		}
		if i>j{
			break
		}
		arr[midIndex],arr[i]=arr[i],arr[midIndex]
		midIndex=i
	}
	fastSort(arr,begin,midIndex-1)
	fastSort(arr,midIndex+1,end)
}