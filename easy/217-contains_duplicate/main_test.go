package main

import (
	"reflect"
	"testing"
)

func TestContaisnsDuplicate(t *testing.T) {
	tests := []struct {
		nums     []int
		expected bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := containsDuplicate(tt.nums)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("containsDuplicate(%v) = %t, %v", tt.nums, result, tt.expected)
			}
		})
	}

}
