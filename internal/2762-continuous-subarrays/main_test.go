package _762_continuous_subarrays

import "testing"

func Test_continuousSubarrays(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{"", args{[]int{5, 4, 2, 4}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := continuousSubarrays(tt.args.nums); got != tt.want {
				t.Errorf("continuousSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
