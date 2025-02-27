package _74_h_index

import "testing"

func Test_hIndex(t *testing.T) {
	type args struct {
		citations []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{citations: []int{3, 0, 6, 1, 5}}, 3},
		{"", args{citations: []int{3, 0, 6, 1, 5, 9}}, 3},
		{"", args{citations: []int{1, 3, 1}}, 1},
		{"", args{citations: []int{100}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hIndex(tt.args.citations); got != tt.want {
				t.Errorf("hIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
