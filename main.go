package main

import "fmt"

// containsDuplicate returns true if any value appears more than once in nums.
func containsDuplicate(nums []int) bool {
	
	seen := make(map[int]struct{}, len(nums))
	for _, x := range nums {
		if _, ok := seen[x]; ok {
			return true
		}
		seen[x] = struct{}{}
	}
	return false
}

func main() {
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{}, false},
		{[]int{99, 99}, true},
		{[]int{-1, -2, -3, -1}, true},
	}

	for _, t := range tests {
		got := containsDuplicate(t.nums)
		fmt.Printf("containsDuplicate(%v) = %v (want %v)\n", t.nums, got, t.want)
	}
}
