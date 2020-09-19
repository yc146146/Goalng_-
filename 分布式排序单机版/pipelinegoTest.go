package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"yinchen.com/文件/9.分布式排序/pipelineMiddleWare"
)

func main1z() {
	//写入
	var filename = "data.in"

	var count = 1000000

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//管道装随机数
	mypipe := pipelineMiddleWare.RandomSource(count)


	writer := bufio.NewWriter(file)
	//
	pipelineMiddleWare.WriterSlink(writer, mypipe)
	//刷新
	writer.Flush()

	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	mypipeX := pipelineMiddleWare.ReaderSource(bufio.NewReader(file), -1)
	counter := 0
	for v := range mypipeX{
		fmt.Println(v)
		counter++
		if counter>1000{
			break
		}
	}



}

func main() {
	go func() {
		myp:= pipelineMiddleWare.Merge(
			pipelineMiddleWare.InMemorySort(pipelineMiddleWare.ArraySource(3,9,2,1,10)),
			pipelineMiddleWare.InMemorySort(pipelineMiddleWare.ArraySource(13,19,12,11,110)),
		)

		for v:=range myp{
			fmt.Println(v)
		}
	}()

	time.Sleep(time.Second*10)

}