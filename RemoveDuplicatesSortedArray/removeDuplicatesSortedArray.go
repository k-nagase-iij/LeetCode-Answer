package RemoveDuplicatesSortedArray

func removeDuplicates(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums = removeElem(nums, i)
			i--
		}
	}
	return len(nums)
}

func removeElem(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}
