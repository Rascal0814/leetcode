package _080_range_frequency_queries

import "testing"

func TestRangeFreqQuery_Query(t *testing.T) {
	type fields struct {
		origin []int
	}
	type args struct {
		left  int
		right int
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
		{"", fields{origin: []int{2, 2, 1, 2, 2}}, args{2, 4, 1}, 1},
		{"", fields{origin: []int{2, 2, 1, 2, 2}}, args{1, 3, 1}, 1},
		{"", fields{origin: []int{2, 2, 1, 2, 2}}, args{0, 3, 1}, 1},
		{"", fields{origin: []int{1, 1, 1, 2, 2}}, args{0, 1, 2}, 0},
		{"", fields{origin: []int{1, 1, 1, 2, 2}}, args{0, 2, 1}, 3},
		{"", fields{origin: []int{1, 1, 1, 2, 2}}, args{3, 3, 2}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := Constructor(tt.fields.origin)
			if got := rr.Query(tt.args.left, tt.args.right, tt.args.value); got != tt.want {
				t.Errorf("Query() = %v, want %v", got, tt.want)
			}
		})
	}
}
