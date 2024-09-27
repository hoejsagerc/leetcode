package main

// two pointers method where i move from each side, but i only move a single
// pointer at the time.
// this method is more effecient in terms of runtime, but requires more memory
func twoSum(numbers []int, target int) []int {
	res := []int{}
	// initializes the right pointer at the start, and the left pointer at the end
	left, right := 0, len(numbers)-1

	// while left is less than right
	for left < right {
		// calculate the sum of the two numbers
		calc := numbers[left] + numbers[right]

		// if the sum is equal to the target, then add result
		if calc == target {
			res = append(res, left+1)
			res = append(res, right+1)
			break
		}

		// if the sum is above the target then only move the rigth cursor
		if calc > target {
			right--
		}

		// if the sum is less than the target then only move the left cursor
		if calc < target {
			left++
		}
	}
	return res
}

// this method is much slower, but beats 94% of the solutions in memory usage
// here i am bruteforcing the problem by keeping the left pointer at a number and then
// trying all other numbers. And if not found then i move both pointers 1 step do the right
// and do the same thing again.
func firstAttempt(numbers []int, target int) []int {
	res := []int{}
	// creating two pointers. left starting from the first indecies in the
	// numbers array. and right starting from the next number
	left, right := 0, 1

	// while left is less than or equal to the length of numbers -1
	for left <= len(numbers)-1 {
		// if the two numbers where pointers currently are == target
		// then return the result
		if numbers[left]+numbers[right] == target {
			res = append(res, left+1)
			res = append(res, right+1)
			break
		}

		// if the right is equal to the lenght of numbers -1 then move
		// left cursor 1 to the right and start the right cursor where the left
		// is now located
		if right == len(numbers)-1 {
			left++
			right = left
		}

		// move the right cursor one step to the right
		right++
	}

	return res
}
