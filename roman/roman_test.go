package roman

import "testing"

func TestToRoman(t *testing.T) {
	cases := []struct {
		n    int
		want string
	}{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{9, "IX"},
		{14, "XIV"},
		{40, "XL"},
		{49, "XLIX"},
		{50, "L"},
		{90, "XC"},
		{99, "XCIX"},
		{444, "CDXLIV"},
		{500, "D"},
		{900, "CM"},
		{944, "CMXLIV"},
		{1994, "MCMXCIV"},
		{2024, "MMXXIV"},
		{3000, "MMM"},
		{3888, "MMMDCCCLXXXVIII"},
		{Min, "I"},
		{Max, "MMMCMXCIX"},
	}

	for _, c := range cases {
		got, err := ToRoman(c.n)
		if err != nil {
			t.Errorf("ToRoman(%d) returned error: %v", c.n, err)
			continue
		}
		if got != c.want {
			t.Errorf("ToRoman(%d) = %q, want %q", c.n, got, c.want)
		}
	}
}

func TestToRomanOutOfRange(t *testing.T) {
	for _, n := range []int{0, -1, Max + 1, 10000} {
		if _, err := ToRoman(n); err == nil {
			t.Errorf("ToRoman(%d) = nil error, want error", n)
		}
	}
}

func TestFromRoman(t *testing.T) {
	cases := []struct {
		s    string
		want int
	}{
		{"I", 1},
		{"II", 2},
		{"III", 3},
		{"IV", 4},
		{"V", 5},
		{"IX", 9},
		{"XIV", 14},
		{"XL", 40},
		{"XLIX", 49},
		{"L", 50},
		{"XC", 90},
		{"XCIX", 99},
		{"CDXLIV", 444},
		{"D", 500},
		{"CM", 900},
		{"CMXLIV", 944},
		{"MCMXCIV", 1994},
		{"MMXXIV", 2024},
		{"MMM", 3000},
		{"MMMDCCCLXXXVIII", 3888},
		{"MMMCMXCIX", Max},
		{"mcmxciv", 1994},
		{"XiV", 14},
	}

	for _, c := range cases {
		got, err := FromRoman(c.s)
		if err != nil {
			t.Errorf("FromRoman(%q) returned error: %v", c.s, err)
			continue
		}
		if got != c.want {
			t.Errorf("FromRoman(%q) = %d, want %d", c.s, got, c.want)
		}
	}
}

func TestFromRomanRejectsNonCanonical(t *testing.T) {
	inputs := []string{
		"IIII",
		"VX",
		"IC",
		"IL",
		"XCC",
		"VV",
		"LL",
		"DD",
		"IIX",
		"MCMXCIVI",
	}

	for _, s := range inputs {
		if _, err := FromRoman(s); err == nil {
			t.Errorf("FromRoman(%q) = nil error, want error for non-canonical input", s)
		}
	}
}

func TestFromRomanRejectsInvalid(t *testing.T) {
	inputs := []string{
		"",
		"ABC",
		"MMMM",
		"XYZ",
		"12",
		" I",
		"I ",
	}

	for _, s := range inputs {
		if _, err := FromRoman(s); err == nil {
			t.Errorf("FromRoman(%q) = nil error, want error", s)
		}
	}
}

func TestRoundTrip(t *testing.T) {
	for n := Min; n <= Max; n++ {
		s, err := ToRoman(n)
		if err != nil {
			t.Fatalf("ToRoman(%d) returned error: %v", n, err)
		}
		got, err := FromRoman(s)
		if err != nil {
			t.Fatalf("FromRoman(%q) (from %d) returned error: %v", s, n, err)
		}
		if got != n {
			t.Fatalf("round trip for %d produced %q -> %d", n, s, got)
		}
	}
}
