package controller

import "testing"

func TestRegisterGraphQL(t *testing.T) {
	type args struct {
		c GraphQL
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterGraphQL(tt.args.c)
		})
	}
}
