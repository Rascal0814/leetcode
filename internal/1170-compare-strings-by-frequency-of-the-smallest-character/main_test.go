package _170_compare_strings_by_frequency_of_the_smallest_character

import (
	"reflect"
	"testing"
)

func Test_numSmallerByFrequency(t *testing.T) {
	type args struct {
		queries []string
		words   []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"", args{[]string{"cbd"}, []string{"zaaaz"}}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSmallerByFrequency(tt.args.queries, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numSmallerByFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}
