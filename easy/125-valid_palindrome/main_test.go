package main

import (
	"reflect"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s        string
		expected bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := isPalindrome(tt.s)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("isPalindrome(%v) = %t, %v", tt.s, result, tt.expected)
			}
		})
	}
}
