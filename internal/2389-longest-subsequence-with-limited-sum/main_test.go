package _389_longest_subsequence_with_limited_sum

import (
	"reflect"
	"testing"
)

func Test_answerQueries(t *testing.T) {
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
		{"", args{[]int{4, 5, 2, 1}, []int{3, 10, 21}}, []int{2, 3, 4}},
		{"", args{[]int{736411, 184882, 914641, 37925, 214915}, []int{331244, 273144, 118983, 118252, 305688, 718089, 665450}}, []int{2, 2, 1, 1, 2, 3, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := answerQueries(tt.args.nums, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("answerQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
