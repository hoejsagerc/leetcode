package main

import "strconv"

func main() {
}

func minBitFlips(start int, goal int) int {
	diff := start ^ goal                                          // Using XOR to find the difference between start and goal
	count := 0                                                    // Count the number of 1s in the difference
	for i := 0; i < len(strconv.FormatInt(int64(diff), 2)); i++ { // Convert the difference to binary and iterate through each bit
		if diff&(1<<i) != 0 { // Counting from the right to the left, if the bit is 1, increment the count
			count++
		}
	}
	return count // Return the count
}
