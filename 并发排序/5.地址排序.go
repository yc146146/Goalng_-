package main

import (
	"fmt"
)

func QuickSort(arr []int, addr []int) []int {
	length := len(addr)
	if length <= 1 {
		return addr
	} else {
		//n:=length-1 //n>=0 && n<=length-1
		//n := rand.Int()%length

		splitaddr := addr[0]
		lowaddr := make([]int, 0, 0)
		highaddr := make([]int, 0, 0)
		midaddr := make([]int, 0, 0)
		midaddr = append(midaddr, splitaddr)

		split := arr[0]
		low := make([]int, 0, 0)
		high := make([]int, 0, 0)
		mid := make([]int, 0, 0)
		mid = append(mid, split)

		for i := 1; i < length; i++ {
			if arr[i] < split {
				lowaddr = append(lowaddr, addr[i])
				low = append(low, arr[i])
			} else if arr[i] > split {
				highaddr = append(highaddr, addr[i])
				high = append(high, arr[i])
			} else {
				midaddr = append(midaddr, splitaddr)
				mid = append(mid, split)
			}
		}
		//切割递归处理
		lowaddr, highaddr = QuickSort(low, lowaddr), QuickSort(high, highaddr)
		myarr := append(append(lowaddr, midaddr...), highaddr...)

		//返回地址
		return myarr
	}

}

func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 6, 4, 5, 10}
	arraddr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(arr)
	fmt.Println(arraddr)
	fmt.Println("-------------------------")

	arraddr = QuickSort(arr, arraddr)

	fmt.Println(arr)
	fmt.Println(arraddr)
	fmt.Println("-------------------------")

	for _, v := range arraddr {
		fmt.Println(arr[v])
	}

}
