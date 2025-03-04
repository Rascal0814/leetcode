package _379_minimum_recolors_to_get_k_consecutive_black_blocks

import "testing"

func Test_minimumRecolors(t *testing.T) {
	type args struct {
		blocks string
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"x", args{"BWBBWWBBBWBWWWBWWBBWBWBBWBB", 11}, 4},
		{"x", args{"BWWWBB", 6}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumRecolors(tt.args.blocks, tt.args.k); got != tt.want {
				t.Errorf("minimumRecolors() = %v, want %v", got, tt.want)
			}
		})
	}
}
