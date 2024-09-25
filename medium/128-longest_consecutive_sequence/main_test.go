package main

import (
	"reflect"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := longestConsecutive(tt.nums)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("longestConsecutive(%v) = %d, want %d", tt.nums, result, tt.expected)
			}
		})
	}
}
