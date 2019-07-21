package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

func (u User) Info() {
	fmt.Printf("who:%s\t age:%d\n", u.Name, u.Age)
}

func (u User) Login() {
	fmt.Printf("who:%s\t age:%d login\n", u.Name, u.Age)
}

func ReflectFunc(T interface{}) {
	getType := reflect.TypeOf(T)
	getValue := reflect.ValueOf(T)
	fmt.Printf("t type:%v,value:%v\n", getType, getValue)
	// 获取方法字段
	// 1. 先获取interface的reflect.Type，然后通过NumField进行遍历
	// 2. 再通过reflect.Type的Field获取其Field
	// 3. 最后通过Field的Interface()得到对应的value
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)
		fmt.Printf("%s: %v \n", m.Name, m.Type)
		//call this func
		methodValue := getValue.MethodByName(m.Name)
		args := make([]reflect.Value, 0)
		methodValue.Call(args)
	}
}

func main() {
	fmt.Println("vim-go")
	var num float64 = 4.14
	fmt.Println("typeof:", reflect.TypeOf(num))
	fmt.Println("valueof:", reflect.ValueOf(num))
	p := &num
	fmt.Println("p typeof:", reflect.TypeOf(p))
	fmt.Println("p valueof:", reflect.ValueOf(p))

	u := User{Name: "ppl", Age: 26}
	ReflectFunc(u)
}
