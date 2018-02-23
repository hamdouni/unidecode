package main

import "testing"

func Test_isFile(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{"file does not exist", "unknownfile", false},
		{"directories are not allowed", "/", false},
		{"permissions are required", "/root/.bashrc", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isFile(tt.arg); got != tt.want {
				t.Errorf("isFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
