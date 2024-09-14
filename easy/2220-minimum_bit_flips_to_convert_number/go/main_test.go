package main

import "testing"

func TestMinBitFlips(t *testing.T) {
	tests := []struct {
		start int
		goal  int
		want  int
	}{
		{10, 7, 3},
		{3, 4, 3},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := minBitFlips(tt.start, tt.goal)
			if got != tt.want {
				t.Errorf("minBitFlips(%d, %d) = %d; want %d", tt.start, tt.goal, got, tt.want)
			}
		})
	}
}
