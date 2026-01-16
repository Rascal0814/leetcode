package _93_preimage_size_of_factorial_zeroes_function

import "testing"

func Test_preimageSizeFZF(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{0}, 5},
		{"", args{1}, 5},
		{"", args{5}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preimageSizeFZF(tt.args.k); got != tt.want {
				t.Errorf("preimageSizeFZF() = %v, want %v", got, tt.want)
			}
		})
	}
}
