package main

import "fmt"


func ShellSortStep(arr[] int, start int, gap int){
	length := len(arr)
	//插入排序的变种
	for i:=start+gap;i<length;i+=gap{
		backup := arr[i]
		j:=i-gap
		for j>=0 && backup<arr[j]{
			//从前往后移动
			arr[j+gap] = arr[j]
			j-=gap
		}
		//插入
		arr[j+gap] = backup
		fmt.Println(arr)
	}
}

func ShellSort(arr[] int)[]int{
	length := len(arr)
	if length <= 1 {
		return arr

	} else {
		gap := length/2
		for gap >0{
			//处理每个元素的步长
			for i:=0;i<gap;i++{
				ShellSortStep(arr, i, gap)
			}
			//gap--
			gap/=2
		}

		return arr
	}

}


func main() {
	arr := []int{3, 9, 2, 8, 1, 7, 4, 6, 5, 10}

	fmt.Println(ShellSort(arr))
}
