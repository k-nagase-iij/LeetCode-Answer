package MaximumSubarray

import "math"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return math.MinInt32
	}

	sum, max := 0, nums[0]
	for _, n := range nums {
		if sum+n < n {
			sum = n
		} else {
			sum += n
		}
		if max < sum {
			max = sum
		}
	}
	return max
}
