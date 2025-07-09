package _048_earliest_second_to_mark_indices_i

import "testing"

func Test_earliestSecondToMarkIndices(t *testing.T) {
	type args struct {
		nums          []int
		changeIndices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[]int{2, 2, 0}, []int{2, 2, 2, 2, 3, 2, 2, 1}}, 8},
		{"", args{[]int{0, 1}, []int{2, 2, 2}}, -1},
		{"", args{[]int{1, 3}, []int{1, 1, 1, 2, 1, 1, 1}}, 6},
		{"", args{[]int{0, 2, 3, 0}, []int{2, 4, 1, 3, 3, 3, 3, 3, 3, 2, 1}}, 10},
		{"", args{[]int{0}, []int{1, 1, 1}}, 1},
		{"", args{[]int{0, 1}, []int{1, 1, 1, 2, 2}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := earliestSecondToMarkIndices(tt.args.nums, tt.args.changeIndices); got != tt.want {
				t.Errorf("earliestSecondToMarkIndices() = %v, want %v", got, tt.want)
			}
		})
	}
}
