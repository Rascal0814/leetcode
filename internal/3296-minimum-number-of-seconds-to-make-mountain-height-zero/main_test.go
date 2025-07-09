package _296_minimum_number_of_seconds_to_make_mountain_height_zero

import "testing"

func Test_minNumberOfSeconds(t *testing.T) {
	type args struct {
		mountainHeight int
		workerTimes    []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{"", args{4, []int{2, 1, 1}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minNumberOfSeconds(tt.args.mountainHeight, tt.args.workerTimes); got != tt.want {
				t.Errorf("minNumberOfSeconds() = %v, want %v", got, tt.want)
			}
		})
	}
}
