package main

import "fmt"

// 1 9 2 8 7 6 4 5



func SelectSortMax(arr[] int) int {
	//数组长度
	length := len(arr)
	if length <= 1{
		//一个元素的数组,直接返回
		return arr[0]
	}else{
		//假定第一个最大
		max:=arr[0]
		for i:=1;i<length;i++{
			//任何一个比我的大的数,最大的
			if arr[i]>max{
				max = arr[i]
			}
		}
		return max
	}
}

func SelectSort(arr[] int)[] int{
	length := len(arr)
	if length <= 1{
		//一个元素的数组,直接返回
		return arr
	}else{
		//只剩1个元素不需要挑选
		for i:=0;i<length-1;i++{
			//标记索引
			min := i
			//每次选出一个极小值
			for j:=i+1; j<length;j++{
				if arr[min] > arr[j]{
					//保存极小值的索引
					min=j
				}
			}
			if i!=min{
				//数据交换
				arr[i],arr[min]=arr[min],arr[i]
			}
		}

	}
	return arr
}

func main1() {
	arr := []int {1,9,2,8,3,7,4,6,5,0}

	//fmt.Println(SelectSortMax(arr))
	fmt.Println(SelectSort(arr))


}
