package pipelineMiddleWare

import (
	"bufio"
	"net"
)

//给我一个ip端口 127.0.0.1:8090
func NetWordkWrite(addr string, in <-chan int){
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	go func() {
		defer listen.Close()

		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		writer := bufio.NewWriter(conn)
		defer writer.Flush()
		WriterSlink(writer, in)

	}()

}

//给我一个端口 读取数据
func NetWordkRead(addr string)<-chan int{
	out := make (chan int)

	go func() {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			panic(err)
		}

		r := ReaderSource(bufio.NewReader(conn), -1)
		for v:=range r{
			//压入数据
			out<-v
		}

		close(out)
	}()
	return out
}
