package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
	"yinchen.com/文件/9.分布式排序/pipelineMiddleWare"
)

func createNetworkPiple(filename string, filesize int, chunkCount int) <-chan int{

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//管道装随机数
	mypipe := pipelineMiddleWare.RandomSource(filesize/8)

	writer := bufio.NewWriter(file)
	//
	pipelineMiddleWare.WriterSlink(writer, mypipe)
	//刷新
	writer.Flush()

	//大小
	chunksize := filesize/chunkCount
	sortAddr := []string{}
	pipelineMiddleWare.Init()

	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}

	for i:=0;i<chunkCount;i++{
		//移动文件指针位置
		file.Seek(int64(i*chunksize), 0)
		source := pipelineMiddleWare.ReaderSource(bufio.NewReader(file), chunksize)
		addr := ":"+strconv.Itoa(7000+i)
		pipelineMiddleWare.NetWordkWrite(addr, pipelineMiddleWare.InMemorySort(source))
		//地址复制
		sortAddr = append(sortAddr, addr)
	}

	sortresults := []<-chan int{}
	for _,addr := range sortAddr{
		sortresults = append(sortresults, pipelineMiddleWare.NetWordkRead(addr))

	}

	return pipelineMiddleWare.MergeN(sortresults...)

}



//本地 多线程 分布式

//多线程 调用中间件完成
func createPipleline(filename string, filesize int, chunkCount int) <-chan int{
	//var filename = "data.in"

	//var count = 1000000

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//管道装随机数
	mypipe := pipelineMiddleWare.RandomSource(filesize/8)

	writer := bufio.NewWriter(file)
	//
	pipelineMiddleWare.WriterSlink(writer, mypipe)
	//刷新
	writer.Flush()



	//数量
	chunkSize := filesize/chunkCount
	//排序结果 一个数据每一个元素是个管道
	sortResults := []<-chan int{}

	pipelineMiddleWare.Init()

	for i:=0;i<chunkCount;i++{
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		//跳到文件指针
		file.Seek(int64(i*chunkSize), 0)
		source := pipelineMiddleWare.ReaderSource(bufio.NewReader(file), chunkSize)

		sortResults = append(sortResults, pipelineMiddleWare.InMemorySort(source))
	}
	return pipelineMiddleWare.MergeN(sortResults...)
}

//写入文件
func writetofile(in <-chan int, filename string){
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	pipelineMiddleWare.WriterSlink(writer, in)
}

//显示文件
func showfile(filename string){
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	p:= pipelineMiddleWare.ReaderSource(bufio.NewReader(file), -1)

	counter := 0
	for v := range p{
		fmt.Println(v)
		counter++
		if counter>1000{
			break
		}
	}


}

//生成随机数组
func main1() {

	go func() {
		time.Sleep(time.Second*50)
	}()



	p := createPipleline("big.in", 800000, 4)
	writetofile(p,"big.out")
	showfile("big.out")

}

func main() {
	go func() {
		time.Sleep(time.Second*50)
	}()



	p := createNetworkPiple("big.in", 800000, 4)
	writetofile(p,"big.out")
	showfile("big.out")
}