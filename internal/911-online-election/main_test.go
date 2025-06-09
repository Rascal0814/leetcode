package _11_online_election

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		persons []int
		times   []int
		q       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[]int{0, 1, 1, 2, 3, 0, 3, 4, 4, 2, 0, 3, 1, 2, 4}, []int{16, 25, 32, 33, 37, 44, 51, 52, 54, 65, 70, 73, 83, 89, 100}, 53}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			constructor := Constructor(tt.args.persons, tt.args.times)
			if got := constructor.Q(tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}
