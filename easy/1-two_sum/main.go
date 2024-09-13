package main

func twoSum(nums []int, target int) []int {
	numToIndexMap := make(map[int]int)

	for i, num := range nums {
		var diff int = target - num
		if index, ok := numToIndexMap[diff]; ok {
			return []int{index, i}
		} else {
			numToIndexMap[num] = i
		}
	}
	return nil
}
