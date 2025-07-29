package _898_maximum_number_of_removable_characters

import "testing"

func Test_maximumRemovals(t *testing.T) {
	type args struct {
		s         string
		p         string
		removable []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{"abcacb", "ab", []int{3, 1, 0}}, 2},
		{"", args{"abcab", "abc", []int{0, 1, 2, 3, 4}}, 0},
		{"", args{"qobftgcueho", "obue", []int{5, 3, 0, 6, 4, 9, 10, 7, 2, 8}}, 7},
		{"", args{"qlevcvgzfpryiqlwy", "qlecfqlw", []int{12, 5}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumRemovals(tt.args.s, tt.args.p, tt.args.removable); got != tt.want {
				t.Errorf("maximumRemovals() = %v, want %v", got, tt.want)
			}
		})
	}
}
