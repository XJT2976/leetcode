package main

import (
	"fmt"
	"main/leetcode1"
)

type Sum interface {
	Add() error
}

type TestSum struct {
}

func (t *TestSum) Add() error {
	return nil
}

func main() {
	var s Sum
	s = &TestSum{}
	_ = s.Add()
	leetcode1.TwoSum()
	fmt.Println("hello world")
}
