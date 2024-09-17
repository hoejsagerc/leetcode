package main

func twoSum(nums []int, target int) []int {
	// create a new hashmap
	indexMap := make(map[int]int)

	// loop through the nums array
	for i, num := range nums {
		// get the diff between the target and the num in array
		diff := target - num

		// check if the diff exists in the hashmap if the difference between the current
		// element and the target already exists in the hashmap, then we know index of that value
		if idx, ok := indexMap[diff]; ok {
			// return array with the index of the found element in the hashmap
			// and the current element in the arrays index
			return []int{idx, i}
		} else {
			// if difference was not found in hashmap, then add the current element from the array
			// to the hashmap with the element as the key, and the elementes index as the value
			indexMap[num] = i
		}
	}
	// return nil if nothing found
	return nil
}
