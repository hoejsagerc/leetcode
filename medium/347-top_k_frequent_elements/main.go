package main

import (
	"sort"
)

//TODO - this can be done with bucket sort -- look at needcode video

func topKFrequent(nums []int, k int) []int {
	// create new hashmap for storing the number as key and the count as the value
	idxMap := make(map[int]int)
	result := []int{}

	// for each number in nums add number as key and if already exists then bumb with 1
	for _, v := range nums {
		idxMap[v]++
	}

	// since we need the most frequent i will sort the keys(the actual number) based on the value(the count of the number)
	sortedMap := sortKeys(idxMap)

	// loop through the sorted slice of keys(numbers)
	for i, value := range sortedMap {
		// if the index is lower than frequenzy "k" then append the key(number) of the sorted slice of keys to the result
		if i < k {
			result = append(result, value)
		} else {
			// if we are above the frequenzy then break the loop
			break
		}
	}

	return result
}

func sortKeys(hashMap map[int]int) []int {
	// create a new slice on ints with the lenght of the haspmap
	keys := make([]int, 0, len(hashMap))
	// for each key(number) in the hashmap add the number to the slice
	for k := range hashMap {
		keys = append(keys, k)
	}

	// then sort the slice based on the values in the hashmap. so add the lowest value(count of number) to the right
	// so push the number which has the highest count to the rigt/top of the slice
	sort.Slice(keys, func(i, j int) bool {
		return hashMap[keys[i]] > hashMap[keys[j]]
	})

	return keys
}
