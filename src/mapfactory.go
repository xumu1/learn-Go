package main

import (
	"fmt"
	"testing"
)

func main() {
	fun1(func(n1 int, n2 int) int {
		fmt.Println(n1,n2)
		return n1+n2
	})(5)
}

func test1(t *testing.T) {
	m:=map[int]func(op int)int{}
	m[0] = func(op int) int {
		return op+1
	}
}
func test2(t *testing.T)   {
	fun1(func(n1 int, n2 int) int {
		fmt.Println(n1,n2)
		return n1+n2
	})(5)
}
func fun1(inner func(n1 int,n2 int)int) func(n int) int  {
	return func(n int) int {
		n1 := n
		n2 := 2*n
		i := inner(n1, n2)
		fmt.Println(i)
		return i
	}
}
