/* 基本思想：
将一个记录插入到已排序好的有序表中，从而得到一个新，记录数增1的有序表。
即：先将序列的第1个记录看成是一个有序的子序列，然后从第2个记录逐个进行插入，直至整个序列有序为止。
*/
package insert_sort

func StraightInsertSort(arr []int64) []int64{
	for i:=1;i<len(arr);i++{
		tmp := arr[i]
		for j:=i-1;j>=0;j--{
			if tmp<arr[j]{
				arr[j+1]=arr[j]
				if j==0{
					arr[0]=tmp
					break
				}
			}else{
				arr[j+1]=tmp
				break
			}
		}
	}
	return arr
}