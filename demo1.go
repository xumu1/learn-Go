package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	//(a) 指针类型（Pointer）
	//(b) 数组类型
	//(c) 结构化类型(struct)
	//(d) Channel 类型
	//(e) 函数类型
	//(f) 切片类型
	//(g) 接口类型（interface）
	//(h) Map 类型

	// 需要分配内存
	//var pointer1 *int
	//pointer1=new(int)
	//*pointer1 = 1

	// 无需
	//var a1 []int
	//a1 = append(a1, 1)
	//println(a1)

	// 无需
	//type people struct{name string}
	//var pp1 people
	//pp1.name = "5"
	//println(pp1.name)

	//var chan1 chan int
	//println(chan1)

	//var fun1 func()
	//fun1()

	//var interface1 interface{}
	//println(interface1)

	// 需要初始化
	//var map1 map[string]string
	//println(map1)
	//var map1 map[string]string
	//map1["hh"]="hh"

}
