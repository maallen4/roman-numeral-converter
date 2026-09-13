package roman_test

import (
	"fmt"

	"romanconv/roman"
)

func ExampleToRoman() {
	s, err := roman.ToRoman(1994)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
	// Output: MCMXCIV
}

func ExampleToRoman_outOfRange() {
	_, err := roman.ToRoman(4000)
	fmt.Println(err)
	// Output: roman: 4000 is out of range [1, 3999]
}

func ExampleFromRoman() {
	n, err := roman.FromRoman("XL")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)
	// Output: 40
}

func ExampleFromRoman_nonCanonical() {
	_, err := roman.FromRoman("IIII")
	fmt.Println(err)
	// Output: roman: "IIII" is not canonical, expected "IV"
}
