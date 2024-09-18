package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
		{[]int{-1, -1}, 1, []int{-1}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := topKFrequent(tt.nums, tt.k)
			sort.Ints(result)
			sort.Ints(tt.expected)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("topKFrequent(%v, %d) = %v, %v", tt.nums, tt.k, result, tt.expected)
			}
		})
	}
}
