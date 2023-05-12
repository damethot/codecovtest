package main

import "testing"

func Test_hello(t *testing.T) {
	type args struct {
		b bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "True",
			args: args{
				b: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hello(tt.args.b)
		})
	}
}
