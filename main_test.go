package main

import "testing"

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		///test cases
		{
			name: "success",
		},
		{
			name: "fail",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_init(t *testing.T) {
	tests := []struct {
		name string
	}{
		///test cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			init()
		})
	}
}
