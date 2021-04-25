package main

import "testing"

func Test_numJewelsInStones(t *testing.T) {
	type args struct {
		jewels string
		stones string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{jewels: "aA", stones: "aAAbbbb"},
			want: 3,
		}, {
			name: "test",
			args: args{jewels: "z", stones: "ZZZ"},
			want: 0,
		},
		{
			name: "test",
			args: args{jewels: "A", stones: "AABBaaCD"},
			want: 2,
		},
		{
			name: "test",
			args: args{jewels: "a", stones: "aaa"},
			want: 3,
		}, {
			name: "test",
			args: args{jewels: "a", stones: "AAAAA"},
			want: 0,
		}, {
			name: "test",
			args: args{jewels: "aZCD", stones: "aazzccddZ"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numJewelsInStones(tt.args.jewels, tt.args.stones); got != tt.want {
				t.Errorf("numJewelsInStones() = %v, want %v", got, tt.want)
			}
		})
	}
}
