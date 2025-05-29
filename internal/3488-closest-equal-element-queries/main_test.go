package _488_closest_equal_element_queries

import (
	"reflect"
	"testing"
)

func Test_solveQueries(t *testing.T) {
	type args struct {
		nums    []int
		queries []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"", args{[]int{1, 3, 1, 4, 1, 3, 2}, []int{0, 3, 5}}, []int{2, -1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveQueries(tt.args.nums, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solveQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
