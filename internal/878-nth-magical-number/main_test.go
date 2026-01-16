package _78_nth_magical_number

import "testing"

func Test_nthMagicalNumber(t *testing.T) {
	type args struct {
		n int
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{5, 2, 4}, 10},
		{"", args{1000000000, 40000, 40000}, 999720007},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nthMagicalNumber(tt.args.n, tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("nthMagicalNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gcd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{5, 6}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gcd(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}
