package roman_numerals

import "strings"

func Roman(n int) string {
	if n >= 50 {
		return "L" + Roman(n-50)
	}
	if n >= 40 && n < 50 {
		return "XL" + Roman(n-40)
	}
	if n >= 10 {
		return "X" + Roman(n-10)
	}
	if n > 5 && n < 10 && (n-5 > 3) {
		return Roman(10-n) + Roman(10)
	}
	if n >= 5 {
		return "V" + Roman(n-5)
	}
	if n > 3 {
		return Roman(5-n) + Roman(5)
	}
	return strings.Repeat("I", n)
}
