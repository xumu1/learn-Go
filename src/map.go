package main

import "testing"

func test(t *testing.T) {
	m := map[string]int{"1":1,"2":2}
	if _,ok := m["s"];ok{
		t.Log("fff")
	}
}
