package _69_majority_element

import "testing"

func Test_majorityElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{nums: []int{3, 2, 3}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement(tt.args.nums); got != tt.want {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
			if got := majorityElement2(tt.args.nums); got != tt.want {
				t.Errorf("majorityElement2() = %v, want %v", got, tt.want)
			}
		})
	}
}
