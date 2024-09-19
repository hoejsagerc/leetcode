package main

// -------------------- METHOD 2 ------------------------------
// Using Bucket Sort

func topKFrequent(nums []int, k int) []int {
	// creating a map for counting the occurences of a number
	m := make(map[int]int)

	// populate the count map where the key is the number from nums, and the value is it's count
	// example: [1 1 1 2 2 3] => [1:3 2:2 3:1]
	for _, val := range nums {
		m[val]++
	}

	// create a new bucket array where the index represents the frequency
	// and the value at each index is a list of numbers with that frequency
	bucket := make([][]int, len(nums)+1)

	// looping over the count map m with the key -> the number occured and value -> the count of that number
	for key, val := range m {
		// append the key (the number from nums) to the count of occurences of the number
		bucket[val] = append(bucket[val], key)
	}

	// creating a new array for the result which has the length of the nums array
	result := make([]int, 0, k)

	// Iterate over the bucket array in reverse order to get the most frequent elements first
	for i := len(bucket) - 1; i >= 0; i-- {
		// Iterate over the numbers in the current bucket
		for _, val := range bucket[i] {
			// if k > 0 meaning that there are still more elements we need to include in the frequency
			if k > 0 {
				result = append(result, val)
				k--
			}
		}
	}

	return result
}

// -------------------- METHOD 1 ------------------------------

// func topKFrequent(nums []int, k int) []int {
// 	// create new hashmap for storing the number as key and the count as the value
// 	idxMap := make(map[int]int)
// 	result := []int{}

// 	// for each number in nums add number as key and if already exists then bumb with 1
// 	for _, v := range nums {
// 		idxMap[v]++
// 	}

// 	// since we need the most frequent i will sort the keys(the actual number) based on the value(the count of the number)
// 	sortedMap := sortKeys(idxMap)

// 	// loop through the sorted slice of keys(numbers)
// 	for i, value := range sortedMap {
// 		// if the index is lower than frequenzy "k" then append the key(number) of the sorted slice of keys to the result
// 		if i < k {
// 			result = append(result, value)
// 		} else {
// 			// if we are above the frequenzy then break the loop
// 			break
// 		}
// 	}

// 	return result
// }

// func sortKeys(hashMap map[int]int) []int {
// 	// create a new slice on ints with the lenght of the haspmap
// 	keys := make([]int, 0, len(hashMap))
// 	// for each key(number) in the hashmap add the number to the slice
// 	for k := range hashMap {
// 		keys = append(keys, k)
// 	}

// 	// then sort the slice based on the values in the hashmap. so add the lowest value(count of number) to the right
// 	// so push the number which has the highest count to the rigt/top of the slice
// 	sort.Slice(keys, func(i, j int) bool {
// 		return hashMap[keys[i]] > hashMap[keys[j]]
// 	})

// 	return keys
// }
