package main

import (
	"testing"
)

func TestHashPNG(t *testing.T) {
	tests := []struct {
		name     string
		file1    string
		file2    string
		shouldEq bool
	}{
		{"a files", "fixtures/a1.png", "fixtures/a2.png", true},
		{"b files", "fixtures/b1.png", "fixtures/b2.png", true},
		{"c files", "fixtures/c1.png", "fixtures/c2.png", true},
		{"d files", "fixtures/d1.png", "fixtures/d2.png", true},
		{"e files", "fixtures/e1.png", "fixtures/e2.png", false},
		{"g files", "fixtures/g1.png", "fixtures/g2.png", true},
		{"f files", "fixtures/f1.png", "fixtures/f2.png", false},
		{"h files", "fixtures/h1.png", "fixtures/h2.png", false},
		{"i files", "fixtures/i1.png", "fixtures/i2.png", false},
		{"j files", "fixtures/j1.png", "fixtures/j2.png", false},
		{"j files", "fixtures/k1.png", "fixtures/k2.png", false},
		{"j files", "fixtures/l1.png", "fixtures/l2.png", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash1, err := HashPNG(tt.file1)
			if err != nil {
				t.Fatalf("Failed to hash %s: %v", tt.file1, err)
			}

			hash2, err := HashPNG(tt.file2)
			if err != nil {
				t.Fatalf("Failed to hash %s: %v", tt.file2, err)
			}

			if tt.shouldEq && hash1 != hash2 {
				t.Errorf("Expected hashes to be equal for %s and %s, but got %s and %s",
					tt.file1, tt.file2, hash1, hash2)
			}

			if !tt.shouldEq && hash1 == hash2 {
				t.Errorf("Expected hashes to be different for %s and %s, but both got %s",
					tt.file1, tt.file2, hash1)
			}
		})
	}
}
func TestQuantizeColor(t *testing.T) {
	tests := []struct {
		name  string
		r     uint32
		g     uint32
		b     uint32
		wantR uint8
		wantG uint8
		wantB uint8
	}{
		{
			name:  "black",
			r:     0x0000,
			g:     0x0000,
			b:     0x0000,
			wantR: 0,
			wantG: 0,
			wantB: 0,
		},
		{
			name:  "white",
			r:     0xffff,
			g:     0xffff,
			b:     0xffff,
			wantR: 128,
			wantG: 128,
			wantB: 128,
		},
		{
			name:  "dark colors quantize to 0",
			r:     0x3fff,
			g:     0x3fff,
			b:     0x3fff,
			wantR: 0,
			wantG: 0,
			wantB: 0,
		},
		{
			name:  "bright colors quantize to 128",
			r:     0x8000,
			g:     0x8000,
			b:     0x8000,
			wantR: 128,
			wantG: 128,
			wantB: 128,
		},
		{
			name:  "mixed values",
			r:     0x2000,
			g:     0x8000,
			b:     0xffff,
			wantR: 0,
			wantG: 128,
			wantB: 128,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR, gotG, gotB := QuantizeColor(tt.r, tt.g, tt.b)
			if gotR != tt.wantR || gotG != tt.wantG || gotB != tt.wantB {
				t.Errorf("QuantizeColor(%d, %d, %d) = (%d, %d, %d), want (%d, %d, %d)",
					tt.r, tt.g, tt.b, gotR, gotG, gotB, tt.wantR, tt.wantG, tt.wantB)
			}
		})
	}
}
