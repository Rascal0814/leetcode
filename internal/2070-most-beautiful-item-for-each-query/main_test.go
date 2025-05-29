package _070_most_beautiful_item_for_each_query

import (
	"reflect"
	"testing"
)

func Test_maximumBeauty(t *testing.T) {
	type args struct {
		items   [][]int
		queries []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{[][]int{{1, 2}, {3, 2}, {2, 4}, {5, 6}, {3, 5}}, []int{1, 2, 3, 4, 5, 6}}, []int{2, 4, 5, 5, 6, 6}},
		{"", args{[][]int{{1000000000, 1}}, []int{1}}, []int{0}},
		{"", args{[][]int{
			{193, 732},
			{781, 962},
			{864, 954},
			{749, 627},
			{136, 746},
			{478, 548},
			{640, 908},
			{210, 799},
			{567, 715},
			{914, 388},
			{487, 853},
			{533, 554},
			{247, 919},
			{958, 150},
			{193, 523},
			{176, 656},
			{395, 469},
			{763, 821},
			{542, 946},
			{701, 676},
		}, []int{885, 1445, 1580, 1309, 205, 1788, 1214, 1404, 572, 1170, 989, 265, 153, 151, 1479, 1180, 875, 276, 1584}}, []int{962, 962, 962, 962, 746, 962, 962, 962, 946, 962, 962, 919, 746, 746, 962, 962, 962, 919, 962}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumBeauty(tt.args.items, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maximumBeauty() = %v, want %v", got, tt.want)
			}
		})
	}
}
