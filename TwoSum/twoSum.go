package main

import "fmt"

func twoSum(nums []int, target int) []int {
	index := make(map[int]int, len(nums))
	for i, b := range nums {
		if j, ok := index[target-b]; ok {
			return []int{j, i}
		}
		index[b] = i
	}
	return nil
}

func main() {
	nums := []int{3, 2, 4}
	target := 6
	fmt.Printf("answer: %v", twoSum(nums, target))
}
