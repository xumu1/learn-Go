package main

import "testing"

func main() {

}

func test1(t *testing.T) {
	m:=map[int]func(op int)int{}
	m[0] = func(op int) int {
		return op+1
	}
}
