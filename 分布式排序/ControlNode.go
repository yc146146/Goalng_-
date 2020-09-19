package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
	"time"
	"yinchen.com/文件/9.分布式排序/pipelineMiddleWare"
)


func IntTobytes(n int)[]byte{
	data := int64(n)
	byteBuffer := bytes.NewBuffer([]byte{})
	binary.Write(byteBuffer, binary.BigEndian, data)
	return byteBuffer.Bytes()
}

func BytesToInt(bts []byte)int{
	byteBuffer := bytes.NewBuffer(bts)
	var data int64
	binary.Read(byteBuffer, binary.BigEndian, &data)
	return int(data)
}



func ServerMsgHandler(conn net.Conn) <-chan int{
	out := make(chan int, 1024)
	buf := make([]byte, 16)
	defer conn.Close()

	//数组 保存数据
	arr := []int{}
	for {
		n, err:=conn.Read(buf)
		//fmt.Println("n", n)
		if err != nil{
			fmt.Println("server conn close", err)
			return nil
		}
		//clientip := conn.RemoteAddr()

		if n == 16{

			data1 := BytesToInt(buf[:len(buf)/2])
			data2 := BytesToInt(buf[len(buf)/2:])

			//fmt.Println("data1",data1)
			//fmt.Println("data2",data2)
			if data1 == 0 && data2 == 0{
				arr = make([]int, 0, 0)
			}

			if data1 == 1 {
				arr = append(arr, data2)
			}
			if data1 == 0 && data2 == 1{
				fmt.Println("数组接受完成", arr)
				for i:=0;i<len(arr);i++{
					//数组压入管道
					out <-arr[i]
				}
				close(out)
				return out

				arr = make([]int, 0, 0)
			}
		}




	}
	return nil
}

func SendArray(arr []int, conn net.Conn){


	length := len(arr)
	// 0 0 开始传输
	//1 1

	//-1 0
	//-1 1

	mybstart := IntTobytes(0)
	mybstart = append(mybstart, IntTobytes(0)...)
	conn.Write(mybstart)

	for i:=0; i<length;i++{
		mybdata := IntTobytes(1)
		mybdata = append(mybdata, IntTobytes(arr[i])...)
		conn.Write(mybdata)
		//fmt.Println(mybdata)
	}


	mybend := IntTobytes(0)
	mybend = append(mybend, IntTobytes(1)...)
	conn.Write(mybend)

	//fmt.Println(mybend)
}

func main(){

	arrlist := [][]int {{1,9,2,8,7,4,5,6,10,4,6,4},{11,19,12,81,71,14,15,16,101,14,16,111}}

	sortResults := []<-chan int{}
	//last := make(chan int, 1024)

	for i:=0;i<2;i++ {
		tpcaddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:700" + strconv.Itoa(1+i))
		if err != nil{
			//panic(err)
			fmt.Println("server err:",err)
		}

		conn, err := net.DialTCP("tcp", nil, tpcaddr)
		if err != nil{
			panic(err)
		}

		SendArray(arrlist[i], conn)
		sortResults = append(sortResults,ServerMsgHandler(conn))


	}

	last := pipelineMiddleWare.Merge(sortResults[0],sortResults[1])
	for v := range last{
		fmt.Print(" ", v)
	}
	time.Sleep(time.Second*30)

}

