package main

import (
	"fmt"
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
	fmt.Println(subsets([]int{1, 2, 3}))
}

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	backtracking(0, nums, &path, &res)
	return res
}

func backtracking(start int, nums []int, path *[]int, result *[][]int) {
	tmp := make([]int, len(*path))
	for i := range *path {
		tmp[i] = (*path)[i]
	}
	*result = append(*result, tmp)

	for i := start; i < len(nums); i++ {
		*path = append(*path, nums[i])
		backtracking(i+1, nums, path, result)
		*path = (*path)[:len(*path)-1]
	}
}
