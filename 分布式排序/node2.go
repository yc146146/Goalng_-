package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"sort"
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

func MsgHandler(conn net.Conn){

	buf := make([]byte, 16)
	defer conn.Close()
	//数组 保存数据
	arr := []int{}
	for {
		n, err:=conn.Read(buf)
		//fmt.Println("n", n)
		if err != nil{
			fmt.Println("conn close", err)
			return
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
				sort.Ints(arr)
				fmt.Println("排序完成", arr)
				//写入
				mybstart := IntTobytes(0)
				mybstart = append(mybstart, IntTobytes(0)...)
				conn.Write(mybstart)

				for i:=0; i<len(arr);i++{
					mybdata := IntTobytes(1)
					mybdata = append(mybdata, IntTobytes(arr[i])...)
					conn.Write(mybdata)
					fmt.Println(mybdata)
				}

				mybend := IntTobytes(0)
				mybend = append(mybend, IntTobytes(1)...)
				conn.Write(mybend)



				arr = make([]int, 0, 0)
			}

			fmt.Println(arr)
		}


	}

}


func main() {
	server_listener, err := net.Listen("tcp", "127.0.0.1:7002")

	if err != nil{
		fmt.Println(err)
		panic(err)
	}



	//延迟关闭
	defer server_listener.Close()

	//接受消息
	for {
		new_conn, err := server_listener.Accept()
		if err != nil{
			fmt.Println(err)
			panic(err)
		}

		//处理客户端消息
		go MsgHandler(new_conn)
	}

}