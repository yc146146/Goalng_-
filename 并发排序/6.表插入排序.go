package main

import "fmt"

const Int_MAX=int(^uint(0)>>1)


type Node struct {
	value int
	next int

}

//集合
var NL []Node


func InitList(arr []int){
	var node Node
	node = Node{Int_MAX, 1}
	NL = append(NL, node)
	for i:=1; i<=len(arr);i++{
		node = Node{arr[i-1], 0}
		NL = append(NL, node)
	}

	fmt.Println(NL)
}


func ListSort(){
	var i, low, high int
	for i=2;i<len(NL);i++{
		low=0
		high=NL[0].next
		//寻找一个相邻的数组
		for NL[high].value<NL[i].value{
			low=high
			high=NL[high].next
		}
		NL[low].next = i
		NL[i].next = high

	}
	fmt.Println(NL)

}

func Arrange(){
	p:=NL[0].next
	for i:=1;i<len(NL);i++{
		for p<i{
			p=NL[p].next
		}
		//下一个排序的记录
		q:=NL[p].next
		if p!=i{
			NL[p].value, NL[i].value = NL[i].value,NL[p].value
			NL[p].next = NL[i].next
			NL[i].next = p
		}
		p=q
	}
	for i:=1;i<len(NL);i++{
		fmt.Print(" ",NL[i].value)
	}
}

func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 6, 4, 5, 10}
	//初始化
	InitList(arr)
	//排序
	ListSort()

	Arrange()
}