package _201_ugly_number_iii

import "testing"

func Test_nthUglyNumber(t *testing.T) {
	type args struct {
		n int
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{3, 2, 3, 5}, 4},
		{"", args{4, 2, 3, 4}, 6},
		{"", args{5, 2, 3, 3}, 8},
		{"", args{7, 7, 7, 7}, 49},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nthUglyNumber(tt.args.n, tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("nthUglyNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
