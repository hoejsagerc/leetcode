package main

func maxArea(height []int) int {
	// handling if only two entries in the height array
	if len(height) == 2 {
		area := 0
		// condition is needed for always calculating the area based on the lowest
		// column(line) of the two heights
		if height[0] <= height[1] {
			area = height[0] * 1
		} else {
			area = height[1] * 1
		}

		return area
	}

	result := 0
	left, right := 0, len(height)-1
	// two pointer
	for left < right {
		area := 0
		// if the left height is higher or equal to the right then we calculate the area from the
		// right column (since we need to do it from the smallest).
		// when we know the left pointer is the highest or eqaul then we can move the right pointer
		// 1 step to the left
		if height[left] >= height[right] {
			area = height[right] * (right - left)
			right--
		} else if height[left] < height[right] {
			area = height[left] * (right - left)
			left++
		}

		// if the area calculated is higher then new result
		if area > result {
			result = area
		}
	}

	return result
}
