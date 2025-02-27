package _8_length_of_last_word

import "testing"

func Test_lengthOfLastWord2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{s: "Hello World"}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord2(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLastWord2() = %v, want %v", got, tt.want)
			}
		})
	}
}
