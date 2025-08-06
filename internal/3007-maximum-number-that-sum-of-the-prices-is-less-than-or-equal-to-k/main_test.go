package _007_maximum_number_that_sum_of_the_prices_is_less_than_or_equal_to_k

import "testing"

func Test_findMaximumNumber(t *testing.T) {
	type args struct {
		k int64
		x int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{"", args{9, 1}, 6},
		{"", args{7, 2}, 9},
		{"", args{19, 6}, 50},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaximumNumber(tt.args.k, tt.args.x); got != tt.want {
				t.Errorf("findMaximumNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
