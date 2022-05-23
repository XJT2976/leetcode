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
	result := leetcode1.TwoSum([]int{3, 2, 4}, 6)
	fmt.Println(result)
}
