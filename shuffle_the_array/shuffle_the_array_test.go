package main

import (
	"reflect"
	"testing"
)

func Test_shuffle(t *testing.T) {
	type args struct {
		nums []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "3パターン",
			args: args{nums: []int{2, 5, 1, 3, 4, 7}, n: 3},
			want: []int{2, 3, 5, 4, 1, 7},
		},
		{
			name: "2パターン",
			args: args{nums: []int{2, 5, 1, 3}, n: 2},
			want: []int{2, 1, 5, 3},
		},
		{
			name: "2パターン",
			args: args{nums: []int{1, 1, 2, 2}, n: 2},
			want: []int{1, 2, 1, 2},
		},
		{
			name: "0パターン",
			args: args{nums: []int{1, 1, 4, 5, 2, 2, 4, 8, 3, 5}, n: 5},
			want: []int{1, 2, 1, 4, 4, 8, 5, 3, 2, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shuffle(tt.args.nums, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shuffle() = %v, want %v", got, tt.want)
			}
		})
	}
}
