package main

import (
	"fmt"
	"math/rand"
	"time"
)




func isOrder(list []int)bool{
	for i:=1;i<len(list);i++{
		if list[i-1]>list[i]{
			return false

		}
	}
	return true
}

func randList(list []int){
	//新建一个数组
	data := make([]int, len(list))
	copy(data, list)
	//定义随机数种子
	rand.Seed(time.Now().UnixNano())
	//随机选择一个切片
	index := rand.Perm(len(list))
	for i,k:=range index{
		list[i] = data[k]
	}
}



func main() {
	list := []int{1,9,2,8,3,7,4,5}
	fmt.Println(list)
	count := 0
	for {
		if isOrder(list){
			fmt.Println("排序完成", list)
			break
		}else{
			randList(list)
			count++
		}
	}
	fmt.Println(count)

}
