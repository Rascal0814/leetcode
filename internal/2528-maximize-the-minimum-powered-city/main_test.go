package _528_maximize_the_minimum_powered_city

import "testing"

func Test_maxPower(t *testing.T) {
	type args struct {
		stations []int
		r        int
		k        int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{"", args{[]int{1, 2, 4, 5, 0}, 1, 2}, 5},
		{"", args{[]int{4, 2}, 1, 1}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPower(tt.args.stations, tt.args.r, tt.args.k); got != tt.want {
				t.Errorf("maxPower() = %v, want %v", got, tt.want)
			}
		})
	}
}
