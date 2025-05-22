package _904_shortest_and_lexicographically_smallest_beautiful_string

import (
	"fmt"
	"testing"
)

func Test_shortestBeautifulSubstring(t *testing.T) {
	a := "1101"
	b := "1011"
	fmt.Println(a > b)
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"", args{"100011001", 3}, "11001"},
		{"", args{"111111110010001010", 11}, "11111111001000101"},
		{"", args{"001110101101101111", 10}, "10101101101111"},
		{"", args{"110101000010110101", 3}, "1011"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestBeautifulSubstring(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("shortestBeautifulSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
