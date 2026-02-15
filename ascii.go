package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// Named color map using ANSI 24-bit (RGB) escape codes
var colors = map[string]string{
	"crimson":   "\033[38;2;220;20;60m",
	"red":       "\033[38;2;255;0;0m",
	"darkred":   "\033[38;2;139;0;0m",
	"green":     "\033[38;2;0;128;0m",
	"darkgreen": "\033[38;2;0;100;0m",
	"lime":      "\033[38;2;0;255;0m",
	"yellow":    "\033[38;2;255;255;0m",
	"gold":      "\033[38;2;255;215;0m",
	"orange":    "\033[38;2;255;165;0m",
	"blue":      "\033[38;2;0;0;255m",
	"darkblue":  "\033[38;2;0;0;139m",
	"lightblue": "\033[38;2;172;216;230m",
	"cyan":      "\033[38;2;0;255;255m",
	"purple":    "\033[38;2;128;0;128m",
	"magenta":   "\033[38;2;255;0;255m",
	"violet":    "\033[38;2;238;130;238m",
	"pink":      "\033[38;2;255;192;203m",
	"brown":     "\033[38;2;139;69;16m",
	"white":     "\033[37m",
	"black":     "\033[38;2;0;0;0m",
	"gray":      "\033[38;2;128;128;128m",
	"grey":      "\033[38;2;128;128;128m",
}

// ResolveColor converts a color string to an ANSI escape code.
// Supports: named colors, #RRGGBB hex, rgb(R,G,B), hsl(H,S%,L%)
func ResolveColor(color string) string {
	lower := strings.ToLower(strings.TrimSpace(color))

	// Check named colors first
	if code, ok := colors[lower]; ok {
		return code
	}

	// Try #RRGGBB hex format
	if strings.HasPrefix(lower, "#") && len(lower) == 7 {
		r, err1 := strconv.ParseInt(lower[1:3], 16, 64)
		g, err2 := strconv.ParseInt(lower[3:5], 16, 64)
		b, err3 := strconv.ParseInt(lower[5:7], 16, 64)
		if err1 == nil && err2 == nil && err3 == nil {
			return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
		}
	}

	// Try rgb(R, G, B) format
	rgbRe := regexp.MustCompile(`^rgb\(\s*(\d+)\s*,\s*(\d+)\s*,\s*(\d+)\s*\)$`)
	if matches := rgbRe.FindStringSubmatch(lower); matches != nil {
		r, _ := strconv.Atoi(matches[1])
		g, _ := strconv.Atoi(matches[2])
		b, _ := strconv.Atoi(matches[3])
		if r <= 255 && g <= 255 && b <= 255 {
			return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
		}
	}

	// Try hsl(H, S%, L%) format
	hslRe := regexp.MustCompile(`^hsl\(\s*(\d+)\s*,\s*(\d+)%\s*,\s*(\d+)%\s*\)$`)
	if matches := hslRe.FindStringSubmatch(lower); matches != nil {
		h, _ := strconv.Atoi(matches[1])
		s, _ := strconv.Atoi(matches[2])
		l, _ := strconv.Atoi(matches[3])
		if h <= 360 && s <= 100 && l <= 100 {
			r, g, b := hslToRgb(float64(h), float64(s)/100.0, float64(l)/100.0)
			return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
		}
	}

	// Default to white if unrecognized
	return "\033[37m"
}

// hslToRgb converts HSL values to RGB (0-255)
func hslToRgb(h, s, l float64) (int, int, int) {
	h = math.Mod(h, 360) / 360.0

	var r, g, b float64
	if s == 0 {
		r, g, b = l, l, l
	} else {
		var q float64
		if l < 0.5 {
			q = l * (1 + s)
		} else {
			q = l + s - l*s
		}
		p := 2*l - q
		r = hueToRgb(p, q, h+1.0/3.0)
		g = hueToRgb(p, q, h)
		b = hueToRgb(p, q, h-1.0/3.0)
	}
	return int(math.Round(r * 255)), int(math.Round(g * 255)), int(math.Round(b * 255))
}

// hueToRgb is a helper for HSL to RGB conversion
func hueToRgb(p, q, t float64) float64 {
	if t < 0 {
		t += 1
	}
	if t > 1 {
		t -= 1
	}
	if t < 1.0/6.0 {
		return p + (q-p)*6*t
	}
	if t < 1.0/2.0 {
		return q
	}
	if t < 2.0/3.0 {
		return p + (q-p)*(2.0/3.0-t)*6
	}
	return p
}

func AsciiArt(line string, sample []string, color, substring string, noSubStr bool) {
	// Split input by newlines (both literal and escaped)
	re := regexp.MustCompile(`(\n|\\n)`)
	words := re.Split(line, -1)

	// If input consists only of newlines, print empty lines without 8-row blocks
	onlyNl := true
	for _, word := range words {
		if word != "" {
			onlyNl = false
		}
	}
	if onlyNl {
		for i := 0; i < len(words)-1; i++ {
			fmt.Printf("\n")
		}
	}

	// Resolve color once
	colorCode := ResolveColor(color)

	// Print each part on separate lines
	for _, word := range words {
		if word == "" {
			fmt.Printf("\n")
		} else {
			PrintAscii(word, sample, colorCode, substring, noSubStr)
		}
	}
}

// PrintAscii renders a line of text using the ASCII art sample
func PrintAscii(line string, sample []string, colorCode, substring string, noSubStr bool) {
	indexes := FindAllIndexes(line, substring)

	// Each ASCII art character is 8 lines tall
	for i := 0; i < 8; i++ {
		row := ""

		for j, letter := range line {
			if int(letter) < 32 || int(letter) > 127 { // Skip non-ASCII
				continue
			}

			// Find the character's row in the sample
			index := (int(letter)-32)*9 + 1 + i
			result := sample[index]

			// Color the character if it's part of the matching substring
			for _, ind := range indexes {
				if j >= ind && j < ind+len(substring) {
					result = colorCode + result + "\033[0m"
				}
			}

			// Color the entire string if no substring was specified
			if noSubStr {
				result = colorCode + result + "\033[0m"
			}

			row += result
		}

		// Print the row if it has content
		if row != "" {
			fmt.Printf("%s\n", row)
		}
	}
}

// FindAllIndexes returns starting indexes of all non-overlapping occurrences of substr in str
func FindAllIndexes(str, substr string) []int {
	var indexes []int
	if substr == "" {
		return nil
	}
	for i := 0; ; {
		idx := strings.Index(str[i:], substr)
		if idx == -1 {
			break
		}
		indexes = append(indexes, i+idx)
		i += idx + len(substr)
	}
	return indexes
}
