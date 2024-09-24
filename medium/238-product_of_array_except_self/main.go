package main

func productExceptSelf(nums []int) []int {
	length := len(nums)
	res := make([]int, length)

	// starting by setting the prefix to 1
	prefix := 1
	// starting from left to right in the array nums
	for i := 0; i < length; i++ {
		// setting first entry in the res array to the prefix
		res[i] = prefix
		// then multiplying the first number in the nums array with the prefix and
		// setting prefix to this number
		prefix *= nums[i]
	}

	// starting by setting the postfix to 1
	postfix := 1
	// starting from right to left in the array of nums
	for i := length - 1; i > -1; i-- {
		// setting the "last" entry in the nums array by multiplying the entry with the postfix
		res[i] *= postfix
		// the setting the postfix by multiplying current postfix with the "last" entry in the nums array
		postfix *= nums[i]
	}

	return res
}
