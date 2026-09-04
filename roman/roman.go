// Package roman converts between integers and roman numerals.
//
// Every exported function is pure: same input always produces the same
// output (or the same error), and none of them touch a filesystem, a
// clock, or global state. That makes them trivial to table-test and safe
// to call concurrently.
package roman

import (
	"fmt"
	"strings"
)

// Min and Max are the bounds of the classic roman numeral system. There is
// no standard symbol for zero or for numbers past 3999 without adding
// notation (vinculum, apostrophus, ...) this package does not implement.
const (
	Min = 1
	Max = 3999
)

type symbol struct {
	value int
	glyph string
}

// table is ordered from largest to smallest so ToRoman can greedily peel
// off the biggest symbol that still fits. The four subtractive pairs
// (CM, CD, XC, XL, IX, IV) are listed explicitly rather than derived,
// because deriving them is more code than just writing them down.
var table = []symbol{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// ToRoman converts n into its canonical uppercase roman numeral form.
// It returns an error if n falls outside [Min, Max].
func ToRoman(n int) (string, error) {
	if n < Min || n > Max {
		return "", fmt.Errorf("roman: %d is out of range [%d, %d]", n, Min, Max)
	}

	var b strings.Builder
	remaining := n
	for _, s := range table {
		for remaining >= s.value {
			b.WriteString(s.glyph)
			remaining -= s.value
		}
	}
	return b.String(), nil
}

// FromRoman parses a roman numeral back into an integer. Input is matched
// case-insensitively, but only the single canonical spelling of each value
// is accepted: non-canonical strings such as "IIII", "VX", or "IC" are
// rejected even though a naive left-to-right sum would give them a value.
//
// The check is done by re-encoding the parsed value with ToRoman and
// comparing it against the (uppercased) input, rather than by writing a
// second set of validation rules that could drift out of sync with table.
func FromRoman(s string) (int, error) {
	if s == "" {
		return 0, fmt.Errorf("roman: empty string is not a valid numeral")
	}

	upper := strings.ToUpper(s)
	values := make([]int, len(upper))
	for i, r := range upper {
		v, ok := digitValue(r)
		if !ok {
			return 0, fmt.Errorf("roman: invalid character %q in %q", r, s)
		}
		values[i] = v
	}

	total := 0
	for i, v := range values {
		if i+1 < len(values) && v < values[i+1] {
			total -= v
		} else {
			total += v
		}
	}

	if total < Min || total > Max {
		return 0, fmt.Errorf("roman: %q evaluates to %d, out of range [%d, %d]", s, total, Min, Max)
	}

	canonical, err := ToRoman(total)
	if err != nil {
		return 0, err
	}
	if canonical != upper {
		return 0, fmt.Errorf("roman: %q is not canonical, expected %q", s, canonical)
	}

	return total, nil
}

func digitValue(r rune) (int, bool) {
	switch r {
	case 'I':
		return 1, true
	case 'V':
		return 5, true
	case 'X':
		return 10, true
	case 'L':
		return 50, true
	case 'C':
		return 100, true
	case 'D':
		return 500, true
	case 'M':
		return 1000, true
	default:
		return 0, false
	}
}
