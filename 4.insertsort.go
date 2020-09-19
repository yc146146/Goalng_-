package main

import "fmt"

func InsertTest(arr [] int)[] int{
	backup := arr[3]
	//从上个位置循环找到位置插入
	j := 3-1
	for j>=0 && backup<arr[j]{
		//从前往后移动
		arr[j+1] = arr[j]
		j--
	}
	arr[j+1] = backup
	return arr
}

func InsertSort(arr [] int)[]int{
	//数组长度
	length := len(arr)
	if length <= 1{
		//一个元素的数组,直接返回
		return arr
	}else{
		//跳过第一个
		for i:=1;i<length;i++{
			backup := arr[i]
			j:=i-1
			for j>=0 && backup<arr[j]{
				//从前往后移动
				arr[j+1] = arr[j]
				j--
			}
			//插入
			arr[j+1] = backup
			fmt.Println(arr)
		}


		return arr
	}
}


func main() {
	arr := []int {1,19,29,8,3,7,4,6,5,10}
	//fmt.Println(InsertTest(arr))
	fmt.Println(InsertSort(arr))
}
