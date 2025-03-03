package _090_k_radius_subarray_averages

import (
	"reflect"
	"testing"
)

func Test_getAverages(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		//{-1, -1, -1, 37, 32, 34, 5, 2, 6}
		{"", args{[]int{7, 4, 3, 9, 1, 8, 5, 2, 6}, 3}, []int{-1, -1, -1, 5, 4, 4, -1, -1, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAverages2(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAverages() = %v, want %v", got, tt.want)
			}
		})
	}
}
