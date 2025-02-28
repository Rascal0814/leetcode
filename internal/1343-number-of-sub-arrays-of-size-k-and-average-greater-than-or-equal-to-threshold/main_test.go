package _343_number_of_sub_arrays_of_size_k_and_average_greater_than_or_equal_to_threshold

import "testing"

func Test_numOfSubarrays(t *testing.T) {
	type args struct {
		arr       []int
		k         int
		threshold int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "", args: args{arr: []int{11, 13, 17, 23, 29, 31, 7, 5, 2, 3}, k: 3, threshold: 5}, want: 6},
		{name: "", args: args{arr: []int{1, 1, 1, 1, 1}, k: 1, threshold: 0}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numOfSubarrays(tt.args.arr, tt.args.k, tt.args.threshold); got != tt.want {
				t.Errorf("numOfSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
