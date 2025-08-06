package _439_minimize_maximum_of_array

import "testing"

func Test_minimizeArrayValue(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[]int{3, 7, 1, 6}}, 5},
		{"", args{[]int{13, 13, 20, 0, 8, 9, 9}}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimizeArrayValue(tt.args.nums); got != tt.want {
				t.Errorf("minimizeArrayValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
