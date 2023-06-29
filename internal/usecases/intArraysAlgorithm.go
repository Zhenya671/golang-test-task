package usecases

import "sort"

func FindingMissingNumbersInAnUnsortedArray(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	maxValue := nums[len(nums)-1]
	encountered := make([]bool, maxValue+1)
	result := []int{}

	for _, num := range nums {
		encountered[num] = true
	}

	for i := 1; i <= maxValue; i++ {
		if !encountered[i] {
			result = append(result, i)
		}
	}

	return result
}
