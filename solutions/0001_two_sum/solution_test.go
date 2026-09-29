package twosum

import "testing"

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{name: "exemplo 1", nums: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
		{name: "exemplo 2", nums: []int{3, 2, 4}, target: 6, want: []int{1, 2}},
		{name: "exemplo 3", nums: []int{3, 3}, target: 6, want: []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)
			if len(got) != len(tt.want) {
				t.Fatalf("twoSum(%v, %d) = %v, esperado %v", tt.nums, tt.target, got, tt.want)
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Fatalf("twoSum(%v, %d) = %v, esperado %v", tt.nums, tt.target, got, tt.want)
				}
			}
		})
	}
}
