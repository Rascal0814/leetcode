package _83_move_zeroes

import (
	"fmt"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "[0,1,0,3,12]",
			args: args{
				nums: []int{0, 1, 0, 3, 12},
			},
			want: []int{1, 3, 12, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//moveZeroes(tt.args.nums)
			//fmt.Println(tt.args.nums)
			moveZeroesPoint(tt.args.nums)
			fmt.Println(tt.args.nums)
		})
	}
}
