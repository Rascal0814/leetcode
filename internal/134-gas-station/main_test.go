package _34_gas_station

import "testing"

func Test_canCompleteCircuit(t *testing.T) {
	type args struct {
		gas  []int
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{gas: []int{1, 2, 3, 4, 5}, cost: []int{3, 4, 5, 1, 2}}, 3},
		{"", args{gas: []int{2, 3, 4}, cost: []int{3, 4, 3}}, -1},
		{"", args{gas: []int{0, 0, 0}, cost: []int{0, 0, 0}}, -1},
		{"", args{gas: []int{1, 1, 1}, cost: []int{1, 1, 1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canCompleteCircuit(tt.args.gas, tt.args.cost); got != tt.want {
				t.Errorf("canCompleteCircuit() = %v, want %v", got, tt.want)
			}
		})
	}
}
