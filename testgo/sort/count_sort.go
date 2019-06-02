package sort

// Sort a slice array of integers from 1-k (excluded)
func CountSort(nums []int32, k int32) []int32 {
	// count is as long as the range k of integers that we are expecting to
	// be present in the nums array.
	count := make([]int32, k)
	for i := 0; i < len(nums); i++ {
		count[nums[i]] += 1
	}

	// here we want to transform the count array to be a positions array where
	// each element tells us from which position we should start the next element
	for i := 1; i < len(count); i++ {
		count[i] = count[i-1] + count[i]
	}

	// in the final loop and starting from the end we want to reposition all e
	// elements based on where they should be based on the count array
	result := make([]int32, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		result[count[nums[i]]-1] = nums[i]
		count[nums[i]] -= 1
	}

	return result
}
