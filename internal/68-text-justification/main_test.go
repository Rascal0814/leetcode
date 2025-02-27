package _8_text_justification

import (
	"reflect"
	"testing"
)

func Test_fullJustify(t *testing.T) {
	type args struct {
		words    []string
		maxWidth int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{"", args{words: []string{"This", "is", "an", "example", "of", "text", "justification."}, maxWidth: 16}, []string{"This    is    an", "example  of text", "justification.  "}},
		{"", args{words: []string{"What", "must", "be", "acknowledgment", "shall", "be"}, maxWidth: 16}, []string{"What   must   be", "acknowledgment  ", "shall be        "}},
		{"", args{words: []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, maxWidth: 20}, []string{"Science  is  what we", "understand      well", "enough to explain to", "a  computer.  Art is", "everything  else  we", "do                  "}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fullJustify(tt.args.words, tt.args.maxWidth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fullJustify() = %v, want %v", got, tt.want)
			}
		})
	}
}
