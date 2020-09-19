package main

import (
	"fmt"
	"strings"
)

func SelectSortMaxString(arr[] string) string {
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
			//if arr[i]>max{
			//	max = arr[i]
			//}
			if strings.Compare(arr[i], max)>0{
				max = arr[i]
			}
		}
		return max
	}
}

func SelectSortString(arr[] string)[] string{
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
				//if arr[min] > arr[j]{
				//	//保存极小值的索引
				//	min=j
				//}

				if strings.Compare(arr[min], arr[j]) <0{
					min = j
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

func main() {
	arr := []string{"c","a","b","x","z","m","n","d","f"}


	fmt.Println(SelectSortMaxString(arr))
	fmt.Println(SelectSortString(arr))
}