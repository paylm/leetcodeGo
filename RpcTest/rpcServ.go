package main

import (
	"net"
	"fmt"
	"os"
	"net/rpc"
	"errors"
)

type Calc int

/*
Go RPC的函数只有符合下面的条件才能够被远程访问，不然会被忽略
1 函数必须是导出的（首字母大写）
2 必须有两个导出类型的参数
3 第一个参数是接受的参数，第二个参数是返回给客户端的参数，第二个参数必须是指针类型的
4 函数还必须有一个返回值error
 */
func (c *Calc)Add(input []int,res *int) error{
	if input == nil{
		return errors.New("input is nil")
	}
	for _,k := range input{
		*res = *res+k
	}
	return nil
}

func (c *Calc)Rotate(input [][]int,res *[][]int) error{
	n := len(input) - 1
	for i := 0; i < n/2+1; i++ {
		for j := i; j < n && j+i < n; j++ {
			k := input[i][j]
			input[i][j] = input[n-j][i]
			input[n-j][i] = input[n-i][n-j]
			input[n-i][n-j] = input[j][n-i]
			input[j][n-i] = k
		}
	}
	*res = input
	return nil
}

func main(){

	rpc.Register(new(Calc))

	tcpAddr,err:=net.ResolveTCPAddr("tcp",":1234")
	if err != nil {
		fmt.Println("错误了哦")
		os.Exit(1)
	}
	listener,err:=net.ListenTCP("tcp",tcpAddr)
	if err != nil{
		fmt.Println("listen fail")
		os.Exit(1)
	}
	defer listener.Close()

	for{
		//需要自己控制连接，当有客户端连接上来后，我们需要把这个连接交给rpc 来处理
		conn,err:=listener.Accept()
		if err != nil{
			fmt.Printf("conn from %v , err:%v\n",conn.RemoteAddr().String(),err)
		}
		rpc.ServeConn(conn)
	}
}