package main

import "testing"

func Test_add(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "both positive",
			args: args{5, 3},
			want: 8,
		},
		{
			name: "both negative",
			args: args{-5, -3},
			want: -8,
		},
		{
			name: "first positive, second negative",
			args: args{5, -3},
			want: 2,
		},
		{
			name: "first negative, second positiive",
			args: args{-5, 3},
			want: -2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}
