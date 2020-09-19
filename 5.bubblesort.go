package main

import "fmt"


func BubbleFindMax(arr []int)int{
	length := len(arr)
	if length <=1{
		return arr[0]
	}else{
		for i:=0;i<length-1;i++{
			//两两比较
			if arr[i] > arr[i+1]{
				arr[i],arr[i+1] = arr[i+1],arr[i]
			}
		}
		return arr[length - 1]
	}
}

func BubbleSort(arr[]int)[]int{
	length := len(arr)
	if length <=1{
		return arr
	}else{
		//只剩1个不需要冒泡了
		for i:=0;i<length-1;i++{
			isneedexchange := false
			for j:=0;j<length-1-i;j++{
				//两两比较
				if arr[j] > arr[j+1]{
					arr[j],arr[j+1] = arr[j+1],arr[j]
					isneedexchange = true
				}
			}
			if !isneedexchange{

				break
			}
			fmt.Println(arr)
		}
		return arr
	}
}




func main() {
	arr := []int {11,9,2,8,3,7,4,6,5,10}
	fmt.Println(BubbleFindMax(arr))
	fmt.Println(BubbleSort(arr))
}

