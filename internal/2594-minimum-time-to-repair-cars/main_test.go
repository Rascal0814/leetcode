package _594_minimum_time_to_repair_cars

import "testing"

func Test_repairCars(t *testing.T) {
	type args struct {
		ranks []int
		cars  int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{"", args{[]int{4, 2, 3, 1}, 10}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repairCars(tt.args.ranks, tt.args.cars); got != tt.want {
				t.Errorf("repairCars() = %v, want %v", got, tt.want)
			}
		})
	}
}
