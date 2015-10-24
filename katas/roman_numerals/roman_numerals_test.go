package roman_numerals

import "testing"

func TestRomans(t *testing.T) {
	var testCases = map[string]int{
		"I":     1,
		"II":    2,
		"III":   3,
		"IV":    4,
		"V":     5,
		"VI":    6,
		"VII":   7,
		"VIII":  8,
		"IX":    9,
		"X":     10,
		"XII":   12,
		"XIV":   14,
		"XV":    15,
		"XVI":   16,
		"XIX":   19,
		"XX":    20,
		"XXXIX": 39,
		"XL":    40,
		"XLVI":  46,
		"L":     50,
		"LXI":   61,
		"LXVI":  66,
		//"MLXVI": 1066,
		//"MCMLXXXIX": 1989,
	}

	for expected_roman, arabic := range testCases {
		roman := Roman(arabic)

		if expected_roman != roman {
			t.Errorf("Expected %#v, but got %#v", expected_roman, roman)
		}
	}
}
