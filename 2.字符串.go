package main

import (
	"fmt"
	"strings"
	"time"
)

func main2() {
	fmt.Println(3>2)
	//字符串地址的比较
	fmt.Println("a">"b")
	//相等0 不等-1 +1
	//首先比较第一个字母, 左边小于右边-1 大于右边+1
	//第一个比较不成功 比较第二个
	fmt.Println(strings.Compare("a","b"))
	fmt.Println(strings.Compare("a","a"))
	fmt.Println(strings.Compare("a","a1"))
	fmt.Println(strings.Compare("a2","a1"))
}

func main3() {
	fmt.Println("a">"b")
	fmt.Println("a"<"b")
	fmt.Println("a"=="b")

	pa := "a3"
	pb := "a2"

	//go 1.1 1.3 版本 比较地址
	//go 1.10优化
	fmt.Println("pa", &pa,"pb", &pb)
	fmt.Println(pa < pb)

}
//974400

func main() {
	startTime := time.Now().UnixNano()
	//fmt.Println(4<<10)
	fmt.Println(10*1024)

	endTime := time.Now().UnixNano()

	fmt.Println(endTime - startTime)
}