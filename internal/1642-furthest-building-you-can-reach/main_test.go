package _642_furthest_building_you_can_reach

import "testing"

func Test_furthestBuilding(t *testing.T) {
	type args struct {
		heights []int
		bricks  int
		ladders int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[]int{4, 2, 7, 6, 9, 14, 12}, 5, 1}, 4},
		{"", args{[]int{17, 16, 5, 10, 10, 14, 7}, 74, 6}, 6},
		{"", args{[]int{7, 5, 13}, 0, 0}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := furthestBuilding(tt.args.heights, tt.args.bricks, tt.args.ladders); got != tt.want {
				t.Errorf("furthestBuilding() = %v, want %v", got, tt.want)
			}
		})
	}
}
