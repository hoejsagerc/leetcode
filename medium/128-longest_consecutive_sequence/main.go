package main

// pro solution
func longestConsecutive(nums []int) int {
	set := make(map[int]bool)

	// initializing the set with the nums
	for _, num := range nums {
		set[num] = true
	}

	res := 0
	for _, num := range nums {
		if set[num-1] {
			continue
		}

		sequence, temp := 1, num+1
		// same as whie set[temp] exists -> while loop
		for set[temp] {
			sequence++
			temp++
		}

		res = max(sequence, res)

		// if sequence is higher than half of the lenght of nums, then this
		// can only be longest sequence
		if sequence > len(nums)/2 {
			break
		}
	}
	return res
}

// func longestConsecutive(nums []int) int {
// 	myMap := make(map[int]bool)

// 	// we add all the numbers to a map to make it log(1) to lookup
// 	// if a number exists
// 	for _, num := range nums {
// 		myMap[num] = true
// 	}

// 	result := 0
// 	// looping trough the array of nums
// 	for _, num := range nums {

// 		// if the num -1 exists in the map then we now that num is not the first number in a sequence
// 		// since num -1 exists in the array, so we can skip this num
// 		if _, ok := myMap[num-1]; ok {
// 			continue
// 		}

// 		count := 1
// 		// currentNum is used to avoid repeating steps since it will start from the num in the range
// 		// if we where to use i := 0, then we would loop from 0 all the way through.
// 		currentNum := num
// 		// while myMap[i] exists then run otherwise break (golang type while loop)
// 		for {
// 			// if currentNum +1 exists in the map, then add count++ and currentNum++
// 			if _, exists := myMap[currentNum+1]; exists {
// 				count++
// 				currentNum++
// 			} else {
// 				// else the sequence has stopped and we break
// 				break
// 			}
// 		}

// 		// if the count is higher than result the update the result
// 		if count > result {
// 			result = count
// 		}
// 	}
// 	return result
// }
