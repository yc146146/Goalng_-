package pipelineMiddleWare

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"time"
)

//构造时间
var starttime time.Time

func Init() {
	//初始化
	starttime = time.Now()
}

func UseTime() {
	//统计消耗时间
	fmt.Println(time.Since(starttime))
}

//内存排序
func InMemorySort(in <-chan int) <-chan int {
	//新管道
	out := make(chan int, 1024)
	go func() {
		//创建一个数组 存储数据并排序
		data := []int{}
		for v := range in {
			data = append(data, v)
		}
		fmt.Println("数据读取完成", time.Since(starttime))
		//排序
		sort.Ints(data)
		for _, v := range data {
			//压入数据
			out <- v
		}
		//关闭管道
		close(out)
	}()

	return out

}

//合并 两个管道的数据 确保有序 归并有序的数据压入到另一个管道
func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int, 1024)

	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		//归并排序
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1 //取出v1压入 再次读取v1
				v1, ok1 = <-in1
			}else{
				out <- v2 //取出v2压入 再次读取v2
				v2, ok2 = <-in2
			}
		}
		close(out)
		fmt.Println("归并结束")
	}()
	return out
}

//读取数据
func ReaderSource(reader io.Reader, chunksize int)<-chan  int{
	out := make(chan int, 1024)
	go func() {

		buf := make([]byte, 8)
		readsize:=0
		for {
			n, err := reader.Read(buf)
			readsize +=n
			if n>0{
				out<-int(binary.BigEndian.Uint64(buf))
			}
			if err != nil || (chunksize!=-1 && readsize>=chunksize) {
				break
			}
		}


		close(out)
	}()

	return out
}

//写入
func WriterSlink(writer io.Writer, in <-chan int){
	for v:=range in {
		//64位8字节
		buf:=make([]byte, 8)
		//字节转换
		binary.BigEndian.PutUint64(buf, uint64(v))
		writer.Write(buf)//写入
	}

}


//随机数数组
func RandomSource(count int)<-chan int{
	out := make(chan int)
	go func(){
		for i:=0;i<count;i++{
			out<-rand.Int()
		}
		//关闭管道
		close(out)
	}()
	return out
}

//多路合并
func MergeN(inputs...<-chan int)<-chan int{
	if len(inputs) == 1{
		return inputs[0]
	}else{
		m := len(inputs)/2
		//递归
		return Merge(MergeN(inputs[:m]...),MergeN(inputs[:m]...))
	}
}

//
func ArraySource(num...int)<-chan int{
	var out=make(chan int)
	go func() {
		for _,v := range num{
			out<-v

		}
		close(out)
	}()

	return out
}