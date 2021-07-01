package main

import "fmt"

func searchRange(nums []int, target int) []int {
	var x []int

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			x = append(x, i)
		}
	}

	if len(x) == 0 {
		x = append(x, -1)
		x = append(x, -1)
	} else if len(x) == 1 {
		x = append(x, x[0])
	}
	x1 := append(x[:1], x[len(x)-1])
	return x1
}

func main() {
	nums := []int{}
	target := 5
	result := searchRange(nums, target)
	fmt.Println(result)
}
