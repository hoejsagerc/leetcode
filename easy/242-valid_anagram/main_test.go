package main

import (
	"reflect"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := isAnagram(tt.s, tt.t)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("isAnagram(%s, %s) = %t, %t", tt.s, tt.t, result, tt.expected)
			}
		})
	}
}
