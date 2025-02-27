package _09_minimum_size_subarray_sum

import "testing"

func Test_minSubArrayLen(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"x", args{15, []int{1, 2, 3, 4, 5}}, 5},
		{"x", args{11, []int{1, 1, 1, 1, 1, 1, 1, 1}}, 0},
		{"x", args{213, []int{12, 28, 83, 4, 25, 26, 25, 2, 25, 25, 25, 12}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.args.target, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
