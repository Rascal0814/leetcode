package main

import "testing"

func Test_longestPalindromeTwice(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "babad",
			args: args{s: "babad"},
			want: "bab",
		},
		{
			name: "cbbd",
			args: args{s: "cbbd"},
			want: "bb",
		},
		{
			name: "aaaa",
			args: args{s: "aaaa"},
			want: "aaaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindromeTwice(tt.args.s); got != tt.want {
				t.Errorf("longestPalindromeTwice() = %v, want %v", got, tt.want)
			}
		})
	}
}
