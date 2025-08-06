package _648_sell_diminishing_valued_colored_balls

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		inventory []int
		orders    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[]int{2, 5}, 4}, 14},
		{"", args{[]int{3, 5}, 6}, 19},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.inventory, tt.args.orders); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
