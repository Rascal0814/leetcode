package _290_convert_binary_number_in_a_linked_list_to_integer

import "testing"

func Test_getDecimalValue(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{name: "x", args: args{head: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		}}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDecimalValue(tt.args.head); got != tt.want {
				t.Errorf("getDecimalValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
