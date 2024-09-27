package main

import "sort"

func threeSum(nums []int) [][]int {
	results := [][]int{}
	// taking note of this sorting function in golang's std lib
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {

		//check to prevent repeats
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		//iterating throgh the array nums such that;
		//1. target is the negative value of the current element which sums up with the
		//    current element to zero
		//2. left is the element right next to the current element
		//3. right is the farthest or last element in the array

		//this helps because the array is sorted
		target, left, right := -nums[i], i+1, len(nums)-1

		//to continue until left is less than right(i.e all elements have been used)
		for left < right {

			sum := nums[left] + nums[right]

			//check to see if the sum of the two elements in equal to the target which would
			// up with nums[i] to equal zero
			// i can do this because if the sum of the two pointers, equal to the indexed value, then i would have 0 for the target
			if sum == target {
				//if so, append the triplet and keep going
				results = append(results, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				//check to avoid duplicates
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}

				// if the sum is greater than the target, then we need to come lefttward to numbers
				// less value to find a smaller sum
			} else if sum > target {
				right--

				// else, we need to find numbers of greater value which would mean going rightward
			} else if sum < target {
				left++
			}
		}
	}
	return results
}
