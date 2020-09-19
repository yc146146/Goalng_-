package main

import "fmt"

//排序数组 排序列号 级别 线程数量
func QuickSortThread(arr[]int, lastarr chan int, level int, thread int){
	//没加深一级别,多一个线程
	level = level*2

	if len(arr)==0{
		close(lastarr)
		return
	}else if len(arr) == 1{
		//唯一的数据放入管道
		lastarr <- arr[0]
		close(lastarr)
		return
	}else{
		less := make([]int, 0)
		greater := make([]int, 0)
		midder := make([]int, 0)

		//取得第一个数据
		left := arr[0]
		midder = append(midder, left)

		for i:=1;i<len(arr);i++{
			if arr[i]<left{
				less=append(less,arr[i])
			}else if arr[i]>left{
				greater=append(greater,arr[i])
			}else{
				midder=append(midder,arr[i])
			}
		}

		//防止数组的管道
		left_ch :=make(chan int, len(less))

		right_ch :=make(chan int, len(greater))
		//如果线程超过执行数量 顺序调用否则并发调用
		if level <= thread{
			go QuickSortThread(less, left_ch,level, thread)
			go QuickSortThread(greater, right_ch,level, thread)
		}else{
			QuickSortThread(less, left_ch,level, thread)
			QuickSortThread(greater, right_ch,level, thread)
		}

		for i:=range left_ch{
			lastarr<-i
		}

		for _,v :=range midder{
			lastarr<-v
		}

		for i:=range right_ch{
			lastarr<-i
		}

		close(lastarr)

		return

	}
}



func main() {
	arr := []int{1,9,2,8,3,7,6,4,5,10}
	lastarr := make(chan int)

	go QuickSortThread(arr, lastarr, 1,10)
	//显示管道每一个数据
	for v := range(lastarr){
		fmt.Println(v)
	}
}
