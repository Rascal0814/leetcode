package _052_grumpy_bookstore_owner

import "testing"

func Test_maxSatisfied(t *testing.T) {
	type args struct {
		customers []int
		grumpy    []int
		minutes   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[]int{4, 10, 10}, []int{1, 1, 0}, 2}, 24},
		{"", args{[]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSatisfied(tt.args.customers, tt.args.grumpy, tt.args.minutes); got != tt.want {
				t.Errorf("maxSatisfied() = %v, want %v", got, tt.want)
			}
		})
	}
}
