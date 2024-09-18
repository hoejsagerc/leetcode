package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		strs     []string
		expected [][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
		{[]string{""}, [][]string{{""}}},
		{[]string{"a"}, [][]string{{"a"}}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := groupAnagrams(tt.strs)
			// sorting the two slices because leetcode does not count for the order
			sortAnagramGroups(result)
			sortAnagramGroups(tt.expected)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("groupAnagrams(%v) = %v, %v", tt.strs, result, tt.expected)
			}
		})
	}
}

// Helper function to sort the outer and inner slices
func sortAnagramGroups(groups [][]string) {
	// Sort each inner slice
	for _, group := range groups {
		sort.Strings(group)
	}
	// Sort the outer slice based on the first element of each inner slice
	sort.Slice(groups, func(i, j int) bool {
		if len(groups[i]) == 0 || len(groups[j]) == 0 {
			return len(groups[i]) < len(groups[j])
		}
		return groups[i][0] < groups[j][0]
	})
}
