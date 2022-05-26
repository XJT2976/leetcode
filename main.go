package main

import (
	"fmt"
	"main/leetcode18"
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
	fmt.Println(leetcode18.FourSum([]int{-3, -2, -1, 0, 0, 1, 2, 3}, 0))
}
