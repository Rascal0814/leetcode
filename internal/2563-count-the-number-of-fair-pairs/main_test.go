package _563_count_the_number_of_fair_pairs

import "testing"

func Test_countFairPairs(t *testing.T) {
	type args struct {
		nums  []int
		lower int
		upper int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{"", args{[]int{0, 1, 7, 4, 4, 5}, 4, 6}, 6},
		{"", args{[]int{0, 0, 0, 0, 0, 0}, 0, 0}, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countFairPairs(tt.args.nums, tt.args.lower, tt.args.upper); got != tt.want {
				t.Errorf("countFairPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
