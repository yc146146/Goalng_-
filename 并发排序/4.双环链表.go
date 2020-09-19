package main

import "fmt"

type CircleLink struct {
	Id int //数据编号
	Data interface{} //数据
	Prev *CircleLink //上一个节点
	Next *CircleLink //下一个节点
}

func InitHeadNode(data interface{}) * CircleLink{
	return &CircleLink{1,data,nil,nil}
}

//重置头结点
func (head *CircleLink) ResetHeadNode(data interface{}){
	if head.Id == 0{
		head.Id =1
	}

	head.Data = data
}
//判断头节点是否为空
func (head *CircleLink)IsHeadEmpty()bool{
	return head.Next == nil && head.Prev==nil
}

//判断链表是否为空
func (head *CircleLink)IsEmpty()bool{
	return head.Data == nil && head.Next==nil && head.Prev==nil

}

func (head * CircleLink)GetLastNode()*CircleLink{
	curnode := head
	if !head.IsHeadEmpty(){
		for {
			if curnode.Next == head{
				break
			}
			curnode = curnode.Next
		}
	}
	return curnode
}

func (head * CircleLink) Addnode(newnode * CircleLink){
	if head.IsHeadEmpty(){
		head.Next = newnode
		head.Prev = newnode
		newnode.Prev = head
		newnode.Next = head
		return
	}

	//备份第一个数据
	curnode := head
	//标志数据添加到末尾
	flag := false
	for true {
		if curnode == head.Prev{
			break //已经是最后一个节点了
		}else if curnode.Next.Id > newnode.Id{
			//标志下数据应该插入到前列
			flag = true
			break
		}else if curnode.Next.Id == newnode.Id{
			fmt.Println("数据已存在\\n")
			return
		}
		curnode = curnode.Next
	}

	if flag{
	//	最后一个节点 前插
		newnode.Next = curnode.Next
		newnode.Prev = curnode

		curnode.Next.Prev = newnode
		curnode.Next = newnode

	}else{
	//	最后一个节点 后插
		newnode.Prev = curnode
		newnode.Next = curnode.Next
		curnode.Next = newnode
		head.Prev = newnode
	}
}

func (head * CircleLink) Findnodebyid(id int)(* CircleLink, bool){
	if head.IsHeadEmpty() && head.Id == id{
		return head, true
	}else if head.IsHeadEmpty() && head.Id != id{
		return &CircleLink{}, false
	}

	curnode := head
	flag := false
	for {
		if curnode.Id == id{
			flag = true
			break
		}
		if curnode == head.Prev{
			break
		}
		curnode = curnode.Next
	}

	if !flag{
		return &CircleLink{}, false
	}
	return curnode, true
}

func (head * CircleLink) Deletenodebyid(id int)(bool){
	if head.IsEmpty(){
		fmt.Println("空链表无法删除\\n")
		return false
	}

	node, isok := head.Findnodebyid(id)

	if isok{
		//删除第一个节点
		if node == head{
			if head.IsHeadEmpty(){
				head.Next = nil
				head.Prev = nil
				head.Data = nil
				head.Id = 0
				return true
			}

			if head.Next.Next == head{
				nextnode := head.Next
				head.Id = nextnode.Id
				head.Data = nextnode.Data
				head.Prev = nil
				head.Next = nil
				return true

			}
			//双环列表 始终保留第一个节点
			//移动下一个节点为头结点
			nextNodetmp := head.Next
			head.Data = nextNodetmp.Data
			head.Id = nextNodetmp.Id
			head.Next = nextNodetmp.Next
			nextNodetmp.Next.Prev = head
			return true
		}

		//删除最后一个节点
		if node == head.GetLastNode(){
			if node.Prev == head && node.Next==head{
				head.Prev=nil
				head.Next=nil
				return true
			}
			head.Prev = node.Prev
			node.Prev.Next = head
			return true
		}

		//	处理中间节点
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
		return true


	}else{
		fmt.Println("数据找不到")
	}

	return isok
}

func (head * CircleLink) Changenodebyid(id int, data interface{})(bool){
	node, isok := head.Findnodebyid(id)
	if isok{
		node.Data = data

	}else{
		fmt.Println("数据找不到!\\n")
	}

	return isok
}

func (head * CircleLink)Showall(){
	if head.IsEmpty(){
		fmt.Println("空链表无法显示\\n")
		return
	}

	if head.IsHeadEmpty(){
		fmt.Println(head.Id, head.Data, head.Prev, head.Next)
		return
	}
	curnnode := head
	for{
		fmt.Println(curnnode.Id, curnnode.Data, curnnode.Prev, curnnode.Next)
		if curnnode == head.Prev{
			break
		}

		curnnode = curnnode.Next
	}
}

func main() {
	linknode:=InitHeadNode("a")
	node1 := &CircleLink{2,"b",nil,nil}
	node2 := &CircleLink{3,"c",nil,nil}
	node3 := &CircleLink{4,"d",nil,nil}
	node4 := &CircleLink{5,"e",nil,nil}
	fmt.Println("------------------------------------")
	linknode.Showall()
	linknode.Addnode(node1)
	fmt.Println("------------------------------------")
	linknode.Showall()
	linknode.Addnode(node2)
	fmt.Println("------------------------------------")
	linknode.Showall()
	linknode.Addnode(node3)
	fmt.Println("------------------------------------")
	linknode.Showall()
	linknode.Addnode(node4)
	fmt.Println("------------------------------------")
	linknode.Showall()

	fmt.Println("------------------------------------")
	linknode.Deletenodebyid(2)
	linknode.Showall()

	fmt.Println("------------------------------------")
	linknode.Changenodebyid(3,"x")
	linknode.Showall()

}