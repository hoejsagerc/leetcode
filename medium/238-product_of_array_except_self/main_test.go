package main

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
		{[]int{0, 0}, []int{0, 0}},
		{[]int{0, 4, 0}, []int{0, 0, 0}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := productExceptSelf(tt.nums)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("productExceptSelf(%v) = %v; want %v", tt.nums, result, tt.expected)
			}
		})
	}
}
