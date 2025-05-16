package _652_defuse_the_bomb

import (
	"reflect"
	"testing"
)

func Test_decrypt(t *testing.T) {
	type args struct {
		code []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"a", args{[]int{5, 7, 1, 4}, 3}, []int{12, 10, 16, 13}},
		{"a", args{[]int{2, 4, 9, 3}, -2}, []int{12, 5, 6, 13}},
		{"a", args{[]int{5, 2, 2, 3, 1}, 3}, []int{7, 6, 9, 8, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decrypt(tt.args.code, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
