package ex6

import (
	"testing"
)

func TestWriter_Write(t *testing.T) {
	tests := []struct {
		name    string
		writer  Writer
		wantErr bool
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.writer.Write(); (err != nil) != tt.wantErr {
				t.Errorf("Writer.Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
