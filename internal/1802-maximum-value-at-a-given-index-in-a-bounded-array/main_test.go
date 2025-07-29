package _802_maximum_value_at_a_given_index_in_a_bounded_array

import "testing"

func Test_maxValue(t *testing.T) {
	type args struct {
		n      int
		index  int
		maxSum int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{4, 2, 6}, 2},
		{"", args{6, 1, 10}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxValue(tt.args.n, tt.args.index, tt.args.maxSum); got != tt.want {
				t.Errorf("maxValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
