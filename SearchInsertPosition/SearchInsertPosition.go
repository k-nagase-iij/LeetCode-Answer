package SearchInsertPosition

func searchInsert(nums []int, target int) int {
	var i int
	for i = 0; i < len(nums); {
		if nums[i] >= target {
			return i
		}
		i++
	}
	return i
}
