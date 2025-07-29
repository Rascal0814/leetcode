package _982_find_longest_special_substring_that_occurs_thrice_ii

import "testing"

func Test_maximumLength(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{"aaaa"}, 2},
		{"", args{"abcaba"}, 1},
		{"", args{"aada"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumLength(tt.args.s); got != tt.want {
				t.Errorf("maximumLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
