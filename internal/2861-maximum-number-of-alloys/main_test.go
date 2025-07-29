package _861_maximum_number_of_alloys

import "testing"

func Test_maxNumberOfAlloys(t *testing.T) {
	type args struct {
		n           int
		k           int
		budget      int
		composition [][]int
		stock       []int
		cost        []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{3, 2, 15, [][]int{{1, 1, 1}, {1, 1, 10}}, []int{0, 0, 100}, []int{1, 2, 3}}, 5},
		{"", args{3, 2, 15, [][]int{{1, 1, 1}, {1, 1, 10}}, []int{0, 0, 0}, []int{1, 2, 3}}, 2},
		{"", args{1, 7, 48, [][]int{{1}, {5}, {9}, {6}, {4}, {2}, {4}}, []int{6}, []int{1}}, 54},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxNumberOfAlloys(tt.args.n, tt.args.k, tt.args.budget, tt.args.composition, tt.args.stock, tt.args.cost); got != tt.want {
				t.Errorf("maxNumberOfAlloys() = %v, want %v", got, tt.want)
			}
		})
	}
}
