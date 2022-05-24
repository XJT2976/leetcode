package main

import (
	"fmt"
	"main/leetcode69"
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
	fmt.Println(leetcode69.MySqrt(2))
}
