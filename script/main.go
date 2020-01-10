package main

import "fmt"

type Calc struct {
	value1 int
	value2 int
}

// 普通の関数
func Add(c Calc) int {
	return c.value1 + c.value2
}

// メソッド
func (c Calc) Add() int {
	return c.value1 + c.value2
}

func main() {
	q := Calc{value1: 1, value2: 2}
	fmt.Println(Add(q))

	p := Calc{value1: 2, value2: 3}
	fmt.Println(p.Add())
}
