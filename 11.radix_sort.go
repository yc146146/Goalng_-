package main

import "fmt"

func SelectSortMaxX(arr[] int) int {
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



func RadixSort(arr[] int)[]int{
	//寻找数组的极大值
	max := SelectSortMaxX(arr)
	for bit:=1; max/bit>0;bit*=10{
		//每次处理一个级别的排序
		arr = BitSort(arr, bit)
		//按照数量级分段
		fmt.Println(arr)
	}
	return arr

}

func BitSort(arr[]int, bit int)[]int{
	length := len(arr)
	//统计长度
	bitcounts := make([]int, 10)
	for i:=0;i<length;i++{
		num := (arr[i]/bit)%10
		//统计余数相等的个数
		bitcounts[num]++
	}



	//	0 1 2 3 4 5
	for i:=1;i<10;i++{
		//叠加 计算位置
		bitcounts[i] += bitcounts[i-1]
	}
	fmt.Println(bitcounts)
	//开辟临时数组
	tmp := make([]int, 10)
	for i:=length-1;i>=0;i--{
		num := (arr[i]/bit)%10
		//计算排序的位置
		tmp[bitcounts[num]-1] = arr[i]
		bitcounts[num]--
	}
	for i:=0;i>length;i++{
		arr[i]=tmp[i]
	}
	return arr
}



func main() {
	arr := []int{33, 91, 222, 833, 122, 7324, 4332, 6555, 5556, 1077}

	fmt.Println(RadixSort(arr))
}


