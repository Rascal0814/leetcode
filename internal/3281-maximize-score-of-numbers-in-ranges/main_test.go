package _281_maximize_score_of_numbers_in_ranges

import "testing"

func Test_maxPossibleScore(t *testing.T) {
	type args struct {
		start []int
		d     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[]int{6, 0, 3}, 2}, 4},
		{"", args{[]int{2, 6, 13, 13}, 5}, 5},
		{"", args{[]int{0, 9, 2, 9}, 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPossibleScore(tt.args.start, tt.args.d); got != tt.want {
				t.Errorf("maxPossibleScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
