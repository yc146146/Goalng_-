package main

import (
	"fmt"
	"math/rand"
)



func QuickSort(arr[] int)[]int{
	length := len(arr)
	if length <=1{
		return arr
	}else{
		//n:=length-1 //n>=0 && n<=length-1
		n := rand.Int()%length
		splitdata := arr[0]
		low := make([]int, 0, 0)
		high := make([]int, 0, 0)
		mid := make([]int, 0, 0)
		mid=append(mid, splitdata)
		for i:=0;i<length;i++{
			if i == n{
				continue
			}
			if arr[i] < splitdata{
				low=append(low, arr[i])
			}else if arr[i] > splitdata{
				high=append(high, arr[i])
			}else{
				mid=append(mid, arr[i])
			}
		}
		//切割递归处理
		low, high = QuickSort(low), QuickSort(high)
		myarr := append(append(low, mid...), high...)
		return myarr
	}

}


func QuickSort2(arr[] int)[]int{
	length := len(arr)
	if length <=1{
		return arr
	}else{
		splitdata := arr[0]
		low := make([]int, 0, 0)
		high := make([]int, 0, 0)
		mid := make([]int, 0, 0)
		mid=append(mid, splitdata)
		for i:=1;i<length;i++{
			if arr[i] < splitdata{
				low=append(low, arr[i])
			}else if arr[i] > splitdata{
				high=append(high, arr[i])
			}else{
				mid=append(mid, arr[i])
			}
		}
		//切割递归处理
		low, high = QuickSort2(low), QuickSort2(high)
		myarr := append(append(low, mid...), high...)
		return myarr
	}

}

func main() {
	arr := []int {3,9,2,8,1,7,4,6,5,10}
	//fmt.Println(HeapSortMax(arr))
	fmt.Println(QuickSort2(arr))
}

