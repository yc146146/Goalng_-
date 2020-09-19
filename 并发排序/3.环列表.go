package main

import "fmt"

//单列表的节点
type node struct {
	num int
	next *node

}

var head, tail *node

func Addnode(n*node){
	if tail == nil{
		head = n
		n.next = head
		tail = n

	}else{
		tail.next = n
		n.next = head
		tail = n

	}
}

func showlist(head *node){
	if head == nil{
		return
	}else{
		//循环环列表
		for head.next != nil && head != tail{
			fmt.Println(head.num)
			head = head.next
		}
		fmt.Println(head.num)
	}
}

//从第k个, 循环第num个 留下最后一个
func jose(k, num int){
	count := 1
	for i:=0;i<=k-1;i++{
		head = head.next
		tail = tail.next

	}

	for {
		//开始记录次数
		count++
		head = head.next
		tail = tail.next

		if count == num{
			fmt.Println(head.num,"出局")
			tail.next = head.next
			head=head.next

			count = 1
		}

		if head == tail{
			fmt.Println(head.num, "最后一个")
			break
		}
	}
}

func main() {
	for i:=0;i<10;i++{
		n := &node{i, nil}
		Addnode(n)
	}
	showlist(head)
	jose(3,3)
}
