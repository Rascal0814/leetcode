package _8_Merge_Sorted_Array

import "testing"

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
	}{
		{"", args{
			nums1: []int{1, 2, 3, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
		},
		},
		{
			"", args{
				nums1: []int{0},
				m:     0,
				nums2: []int{1},
				n:     1,
			},
		},
		{
			"", args{
				nums1: []int{1},
				m:     1,
				nums2: []int{},
				n:     0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
		})
	}
}
