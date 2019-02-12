package main

import (
	"fmt"
	"os"
	"net/rpc"
	"time"
)

func main(){
	fmt.Println("start cli")

	client,err := rpc.Dial("tcp","127.0.0.1:1234")
	if err != nil{
		fmt.Printf("connect 127.0.0.1:1234 fail ,err:%v\n",err)
		os.Exit(1)
	}
	//rpc call Add
	res := 0
	input := []int{1,2,3,4,5}
	err = client.Call("Calc.Add",input,&res)
	if err != nil{
		fmt.Printf("Calc.Add fail ,err:%v\n",err)
		os.Exit(0)
	}
	fmt.Printf("Calc.Add input:%v,res:%d\n",input,res)

	//rpc call Rotate
	res1 := [][]int{}
	input1 := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	err = client.Call("Calc.Rotate",input1,&res1)
	if err != nil{
		fmt.Printf("Calc.Rotate fail ,err:%v\n",err)
		os.Exit(1)
	}
	fmt.Printf("Calc.Rotate input:%v,res:%v\n",input1,res1)

	//异步调用
	divCall:=client.Go("Calc.Rotate",input1,&res1,nil)

	//使用select模型监听通道有数据时执行，否则执行后续程序
	for {
		select {
		case <-divCall.Done:
			fmt.Printf("Calc.Rotate res=>%v\n",res1)
		default:
			fmt.Println("继续向下执行....")
			time.Sleep(time.Second * 1)
		}
	}
}
