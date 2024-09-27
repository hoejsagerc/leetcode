package main

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		nums     []int
		expected [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{[]int{0, 1, 1}, [][]int{}},
		{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
		{[]int{-1, 0, 1, 1, 2}, [][]int{{-2, 0, 2}, {-2, 1, 1}}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := threeSum(tt.nums)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("threeSum(%v) = %v, wanted: %v\n", tt.nums, result, tt.expected)
			}
		})
	}
}
