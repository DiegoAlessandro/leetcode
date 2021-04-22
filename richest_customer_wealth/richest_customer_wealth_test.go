package main

import "testing"

//Problems: https://leetcode.com/problems/richest-customer-wealth/submissions/
func Test_maximumWealth(t *testing.T) {
	type args struct {
		accounts [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "所持金が同じユーザ",
			args: args{accounts: [][]int{{1, 2, 3}, {3, 2, 1}}},
			want: 6,
		},
		{
			name: "所持金が最大のユーザ",
			args: args{accounts: [][]int{{2, 3, 4}, {7, 3}}},
			want: 10,
		},
		{
			name: "所持金がゼロのユーザ",
			args: args{accounts: [][]int{{0, 0}, {1}}},
			want: 1,
		},
		{
			name: "全員が所持金ゼロ",
			args: args{accounts: [][]int{{0}, {0}}},
			want: 0,
		},
		{
			name: "全員が所持金がマイナス",
			args: args{accounts: [][]int{{0}, {-1}}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumWealth(tt.args.accounts); got != tt.want {
				t.Errorf("maximumWealth() = %v, want %v", got, tt.want)
			}
		})
	}
}
