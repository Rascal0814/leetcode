package _529_maximum_count_of_positive_integer_and_negative_integer

import "testing"

func Test_maximumCount(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[]int{-2, -1, -1, 1, 2, 3}}, 3},
		{"", args{[]int{0, 0}}, 0},
		{"", args{[]int{5, 20, 66, 1314}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumCount(tt.args.nums); got != tt.want {
				t.Errorf("maximumCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
