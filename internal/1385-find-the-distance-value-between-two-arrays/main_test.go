package _385_find_the_distance_value_between_two_arrays

import "testing"

func Test_findTheDistanceValue(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
		d    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[]int{4, 5, 8}, []int{10, 9, 1, 8}, 2}, 2},
		{"", args{[]int{1, 4, 2, 3}, []int{-4, -3, 6, 10, 20, 30}, 3}, 2},
		{"", args{[]int{-8, -7}, []int{4, 10, -4, 5, 2}, 55}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTheDistanceValue(tt.args.arr1, tt.args.arr2, tt.args.d); got != tt.want {
				t.Errorf("findTheDistanceValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
