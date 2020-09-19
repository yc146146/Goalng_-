package main

import "fmt"

type TreeNode struct {
	//数据
	element interface{}
	//左右节点
	left *TreeNode
	right *TreeNode
	npl int

}

//有限队列
type PQ *TreeNode


//开普一个节点
func NewLeftHeap(element interface{})PQ{
	head := new(TreeNode)
	head.element=element
	head.left=nil
	head.right=nil
	head.npl=0
	return PQ(head)

}

func MergeSort(H1, H2 PQ)PQ{
	if H1.left==nil{
		H1.left=H2

	}else{
		H1.right=Merge(H1.right, H2)
		//层级的互换
		if H1.left.npl < H1.right.npl{
			H1.left, H1.right = H1.right, H1.left
		}

		//层级递增
		H1.npl = H1.right.npl+1
	}
	return H1
}

//确保有序
func Merge(H1, H2 PQ)PQ{
	if H1==nil{
		return H2
	}

	if H2 == nil{
		return H1
	}

	//>极大 <极小
	if H1.element.(int)<H2.element.(int){
		return MergeSort(H1,H2)
	}else{
		return MergeSort(H2,H1)
	}
}

func Insert(data interface{}, H PQ)PQ{
	insertnode := new(TreeNode)
	insertnode.element=data
	insertnode.left=nil
	insertnode.right=nil
	insertnode.npl=0

	//插入用归并实现
	H=Merge(insertnode, H)
	return H

}

func DeleteMin(H PQ)(PQ, interface{}){
	if H==nil{
		return nil,nil

	}else{
		leftHeap := H.left
		rightHeap := H.right
		value := H.element
		H=nil
		return Merge(leftHeap, rightHeap), value
	}
}

//遍历树
func PrintHQ(H *TreeNode){
	if H==nil{
		return
	}

	PrintHQ(H.left)
	PrintHQ(H.right)
	fmt.Println(H.element, "   ")
}

func main() {
	H:=NewLeftHeap(3)
	H=Insert(2, H)
	H=Insert(1, H)
	H=Insert(4, H)



	PrintHQ(H)
	H,data :=DeleteMin(H)
	fmt.Println(data)
	H,data2 :=DeleteMin(H)
	fmt.Println(data2)
}