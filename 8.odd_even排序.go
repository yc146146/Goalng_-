package main

import "fmt"



func OddEven(arr []int) []int {

	//奇数位 偶数位 都不需要排序是就是有序
	isSorted := false


	for ;isSorted==false; {
		isSorted = true
		//奇数位
		for i:= 1;i<len(arr)-1;i=i+2{
			if arr[i]>arr[i+1]{
				arr[i],arr[i+1]=arr[i+1],arr[i]
				isSorted=false
			}
		}
		fmt.Println("a",arr)
		//偶数位
		for i:= 0;i<len(arr)-1;i=i+2{
			if arr[i]>arr[i+1]{
				arr[i],arr[i+1]=arr[i+1],arr[i]
				isSorted=false
			}
		}
		fmt.Println("b",arr)
	}
	return arr
}



func main() {
	arr := []int {3,9,2,8,1,7,4,6,5,10}
	//fmt.Println(HeapSortMax(arr))
	fmt.Println(OddEven(arr))
}

