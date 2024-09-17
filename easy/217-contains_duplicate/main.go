package main

// -------------------- METHOD 3 --------------------------
// Using a hashset to only run the the nums 1 time which will be a O(n)
// ====> FUNC FACT: leetcode performance (most performant solution)
// 			runtime: 	94ms 	-> beats 21.17%
//			memory: 	10.65mb	-> beats 40.80%

func containsDuplicate(nums []int) bool {
	idxMap := make(map[int]int)

	for _, num := range nums {
		if i, ok := idxMap[num]; ok {
			return true
		} else {
			idxMap[num] = i
		}
	}

	return false
}

// -------------------- METHOD 2 --------------------------
// Sorting method O(nlog n)

// func containsDuplicate(nums []int) bool {
// 	sort.Ints(nums)
// 	for i, num := range nums {
// 		if i+1 < len(nums) && num == nums[i+1] {
// 			return true
// 		}
// 	}
// 	return false
// }

// -------------------- METHOD 1 --------------------------
// Brute forcing method O(n) for each number which would be O(n^2)
// ====> FUNC FACT: leetcode performance
// 			runtime: 	1977ms 	-> beats 5.00%
//			memory: 	8.03mb	-> beats 93.37% (best in memory -> but slower over all)

// func containsDuplicate(nums []int) bool {
// 	var mySlice []int

// 	for _, num := range nums {
// 		if contains(mySlice, num) {
// 			return true
// 		} else {
// 			mySlice = append(mySlice, num)
// 		}
// 	}
// 	return false
// }

// func contains(nums []int, target int) bool {
// 	for _, num := range nums {
// 		if num == target {
// 			return true
// 		}
// 	}
// 	return false
// }
