package _695_maximum_erasure_value

import "testing"

func Test_maximumUniqueSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[]int{10000, 1, 10000, 1, 1, 1, 1, 1, 1}}, 10001},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumUniqueSubarray(tt.args.nums); got != tt.want {
				t.Errorf("maximumUniqueSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
