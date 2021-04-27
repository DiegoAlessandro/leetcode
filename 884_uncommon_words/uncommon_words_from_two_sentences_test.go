package main

import (
	"reflect"
	"testing"
)

func Test_uncommonFromSentences(t *testing.T) {
	type args struct {
		A string
		B string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test",
			args: args{A: "this apple is sweet", B: "this apple is sour"},
			want: []string{"sweet", "sour"},
		}, {
			name: "test",
			args: args{A: "apple apple", B: "banana"},
			want: []string{"banana"},
		}, {
			name: "test",
			args: args{A: "apple", B: "banana banana"},
			want: []string{"apple"},
		}, {
			name: "test",
			args: args{A: "apple apple", B: "banana banana"},
			want: nil,
		}, {
			name: "test",
			args: args{A: "fd kss fd", B: "fd fd kss"},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uncommonFromSentences(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("uncommonFromSentences() = %v, want %v", got, tt.want)
			}
		})
	}
}
