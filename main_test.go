package main

import (
	"testing"
)

func TestGeneratePassword(t *testing.T) {
	tests := []struct {
		name       string
		length     int
		complexity int
		wantErr    bool
	}{
		{"invalid length", -1, 1, true},
		{"invalid length", 8192, 1, true},
		{"invalid length", 0, 1, true},
		{"invalid password complexity level", 20, 0, true},
		{"valid", 20, 1, false},
		{"valid", 20, 2, false},
		{"valid", 20, 3, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := generatePassword(tt.length, tt.complexity)
			if (err != nil) != tt.wantErr {
				t.Errorf("generatePassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}