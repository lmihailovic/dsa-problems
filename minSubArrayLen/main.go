package main

import "math"

func minSubArrayLen(target int, nums []int) float64 {
	lb := 0
	rb := 0
	sum := 0
	solution := float64(nums[len(nums)-1])

	for rb < len(nums) {
		if sum < target {
			sum += nums[rb]
			rb += 1
		} else {
			solution = math.Min(solution, float64(rb-lb))
			sum -= lb
			lb += 1
		}
	}

	return solution
}

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	target := 7

	println(int(minSubArrayLen(target, nums)))
}
