package _841_maximum_sum_of_almost_unique_subarray

import "testing"

func Test_maxSum(t *testing.T) {
	type args struct {
		nums []int
		m    int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{"x", args{[]int{2, 6, 7, 3, 1, 7}, 3, 4}, 18},
		{"x", args{[]int{1, 2, 1, 2, 1, 2, 1}, 3, 3}, 0},
		{"x", args{[]int{5, 9, 9, 2, 4, 5, 4}, 1, 3}, 23},
		{"x", args{[]int{1, 2, 2}, 2, 2}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSum(tt.args.nums, tt.args.m, tt.args.k); got != tt.want {
				t.Errorf("maxSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
