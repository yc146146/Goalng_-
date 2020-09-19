package main

import "fmt"


func cocktail(arr []int)[]int{
	//每次循环,正向冒泡一次
	for i:=0;i<len(arr)/2;i++{
		left := 0
		right := len(arr)-1

		for left <= right{
			if arr[left] > arr[left+1]{
				arr[left], arr[left+1] = arr[left+1], arr[left]
			}
			left++
			if arr[right-1] > arr[right]{
				arr[right-1], arr[right] = arr[right], arr[right-1]
			}
			right--
		}
		fmt.Println(arr)
	}
	return arr
}



func main() {
	arr := []int {1,9,2,8,3,7,4,6,5,0}

	//fmt.Println(SelectSortMax(arr))
	fmt.Println(cocktail(arr))


}
