package main

import "fmt"

func main() {
	nums := []int{3, 0, 1}

	missing := missingNumber(nums)
	fmt.Println(nums)
	fmt.Println(missing)
}

func missingNumber(nums []int) int {
	arrLen := len(nums)
	sum := arrLen * (arrLen + 1) / 2

	for _, v := range nums {
		sum -= v
	}

	return sum
}
