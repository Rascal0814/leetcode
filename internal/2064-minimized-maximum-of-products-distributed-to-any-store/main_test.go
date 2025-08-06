package _064_minimized_maximum_of_products_distributed_to_any_store

import "testing"

func Test_minimizedMaximum(t *testing.T) {
	type args struct {
		n          int
		quantities []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{6, []int{11, 6}}, 3},
		{"", args{1, []int{1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimizedMaximum(tt.args.n, tt.args.quantities); got != tt.want {
				t.Errorf("minimizedMaximum() = %v, want %v", got, tt.want)
			}
		})
	}
}
