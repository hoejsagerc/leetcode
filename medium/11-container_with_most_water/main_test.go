package main

import (
	"reflect"
	"testing"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		height   []int
		expected int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{2, 1}, 1},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := maxArea(tt.height)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("maxArea(%v) = %v, wanted: %v\n", tt.height, result, tt.expected)
			}
		})
	}
}
