package main

import (
	"reflect"
	"strings"
	"testing"
)

// Test FindAllIndexes function
func TestFindAllIndexes(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		substr   string
		expected []int
	}{
		{"Empty substring", "hello", "", nil},
		{"No match", "hello", "xyz", nil},
		{"Single match at start", "hello", "hel", []int{0}},
		{"Single match at end", "hello", "llo", []int{2}},
		{"Multiple matches", "abcabcabc", "abc", []int{0, 3, 6}},
		{"Overlapping not counted", "aaa", "aa", []int{0}},
		{"Substring in word", "a king kitten have kit", "kit", []int{7, 19}},
		{"Single character", "hello", "l", []int{2, 3}},
		{"Full string match", "test", "test", []int{0}},
		{"Case sensitive", "Hello", "hello", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindAllIndexes(tt.str, tt.substr)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("FindAllIndexes(%q, %q) = %v, want %v",
					tt.str, tt.substr, result, tt.expected)
			}
		})
	}
}

// Test ResolveColor function
func TestResolveColor(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// Named colors
		{"Red named", "red", "\033[38;2;255;0;0m"},
		{"Green named", "green", "\033[38;2;0;128;0m"},
		{"Blue named", "blue", "\033[38;2;0;0;255m"},
		{"Yellow named", "yellow", "\033[38;2;255;255;0m"},
		{"Orange named", "orange", "\033[38;2;255;165;0m"},
		{"Case insensitive name", "RED", "\033[38;2;255;0;0m"},
		{"Mixed case name", "Green", "\033[38;2;0;128;0m"},

		// Hex colors
		{"Hex red", "#ff0000", "\033[38;2;255;0;0m"},
		{"Hex green", "#00ff00", "\033[38;2;0;255;0m"},
		{"Hex blue", "#0000ff", "\033[38;2;0;0;255m"},
		{"Hex uppercase", "#FF0000", "\033[38;2;255;0;0m"},

		// RGB colors
		{"RGB red", "rgb(255, 0, 0)", "\033[38;2;255;0;0m"},
		{"RGB green", "rgb(0, 128, 0)", "\033[38;2;0;128;0m"},
		{"RGB no spaces", "rgb(255,0,0)", "\033[38;2;255;0;0m"},

		// HSL colors
		{"HSL red", "hsl(0, 100%, 50%)", "\033[38;2;255;0;0m"},
		{"HSL green", "hsl(120, 100%, 25%)", "\033[38;2;0;128;0m"},

		// Unknown color defaults to white
		{"Unknown color", "unknowncolor", "\033[37m"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ResolveColor(tt.input)
			if result != tt.expected {
				t.Errorf("ResolveColor(%q) = %q, want %q",
					tt.input, result, tt.expected)
			}
		})
	}
}

// Test that all named colors produce valid ANSI codes
func TestColorMapValid(t *testing.T) {
	for name, code := range colors {
		if !strings.HasPrefix(code, "\033[") {
			t.Errorf("Color %q has invalid ANSI prefix: %q", name, code)
		}
		if !strings.HasSuffix(code, "m") {
			t.Errorf("Color %q has invalid ANSI suffix: %q", name, code)
		}
	}
}

// Test hslToRgb conversion
func TestHslToRgb(t *testing.T) {
	tests := []struct {
		name    string
		h, s, l float64
		r, g, b int
	}{
		{"Pure red", 0, 1.0, 0.5, 255, 0, 0},
		{"Pure green", 120, 1.0, 0.5, 0, 255, 0},
		{"Pure blue", 240, 1.0, 0.5, 0, 0, 255},
		{"White", 0, 0, 1.0, 255, 255, 255},
		{"Black", 0, 0, 0, 0, 0, 0},
		{"Gray", 0, 0, 0.5, 128, 128, 128},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, g, b := hslToRgb(tt.h, tt.s, tt.l)
			if r != tt.r || g != tt.g || b != tt.b {
				t.Errorf("hslToRgb(%v, %v, %v) = (%d, %d, %d), want (%d, %d, %d)",
					tt.h, tt.s, tt.l, r, g, b, tt.r, tt.g, tt.b)
			}
		})
	}
}
