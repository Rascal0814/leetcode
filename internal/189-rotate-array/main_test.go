package _89_rotate_array

import (
	"fmt"
	"testing"
)

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
	}{
		{"", args{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 3}},
		{"", args{nums: []int{1, 2}, k: 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate2(tt.args.nums, tt.args.k)
			fmt.Println(tt.args.nums)

		})
	}
}
