package _875_minimum_size_subarray_in_infinite_array

import "testing"

func Test_minSizeSubarray(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[]int{1, 2, 3}, 5}, 2},
		{"", args{[]int{1, 1, 1, 2, 3}, 4}, 2},
		{"", args{[]int{2, 4, 6, 8}, 3}, -1},
		{"", args{[]int{1, 2, 3, 2, 1, 5, 3, 4, 5}, 53}, 19},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSizeSubarray(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("minSizeSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
