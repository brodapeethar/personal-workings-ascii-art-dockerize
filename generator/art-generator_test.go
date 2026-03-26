package artGenerator

import (
	"testing"
)

func TestGenerateArt(t *testing.T) {
	// 1. Define "Table-Driven Tests" (a list of inputs and expected outputs)
	tests := []struct {
		name       string
		input      string
		bannerPath string
		wantErr    bool
	}{
		{"Empty Input", "", "../banners/standard.txt", false},
		{"New Line Only", "\\n", "../banners/standard.txt", false},
		{"Invalid Banner Path", "hello", "missing.txt", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateArt(tt.input, tt.bannerPath)

			// Check if we expected an error and got one
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateArt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// For the "Empty Input" case, ensure result is empty
			if tt.input == "" && got != "" {
				t.Errorf("Expected empty string for empty input, got %q", got)
			}
		})
	}
}
