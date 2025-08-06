package _143_maximum_points_inside_the_square

import "testing"

func Test_maxPointsInsideSquare(t *testing.T) {
	type args struct {
		points [][]int
		s      string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[][]int{{2, 2}, {-1, -2}, {-4, 4}, {-3, 1}, {3, -3}}, "abdca"}, 2},
		{"", args{[][]int{{-19, 14}, {20, 7}, {17, 9}, {3, 15}, {-6, -4}}, "bbade"}, 4},
		{"", args{[][]int{{0, 1}, {0, 0}}, "aa"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPointsInsideSquare(tt.args.points, tt.args.s); got != tt.want {
				t.Errorf("maxPointsInsideSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
