package _226_maximum_candies_allocated_to_k_children

import "testing"

func Test_maximumCandies(t *testing.T) {
	type args struct {
		candies []int
		k       int64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[]int{5, 8, 6}, 3}, 5},
		{"", args{[]int{2, 5}, 11}, 0},
		{"", args{[]int{1, 2, 3, 4, 10}, 5}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumCandies(tt.args.candies, tt.args.k); got != tt.want {
				t.Errorf("maximumCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}
