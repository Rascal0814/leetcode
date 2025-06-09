package LCP_08_ju_qing_hong_fa_shi_jian

import (
	"reflect"
	"testing"
)

func Test_getTriggerTime(t *testing.T) {
	type args struct {
		increase     [][]int
		requirements [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"", args{[][]int{{2, 8, 4}, {2, 5, 0}, {10, 9, 8}}, [][]int{{2, 11, 3}, {15, 10, 7}, {9, 17, 12}, {8, 1, 14}}}, []int{2, -1, 3, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTriggerTime(tt.args.increase, tt.args.requirements); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTriggerTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
