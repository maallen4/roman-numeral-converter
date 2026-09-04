// Command romanconv is a thin CLI shell around the roman package. All the
// actual conversion logic lives there and is unit-tested directly; this
// file just handles argv and stdout/stderr.
package main

import (
	"fmt"
	"os"
	"strconv"

	"romanconv/roman"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <integer|roman-numeral>\n", os.Args[0])
		os.Exit(2)
	}

	arg := os.Args[1]

	if n, err := strconv.Atoi(arg); err == nil {
		out, err := roman.ToRoman(n)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Println(out)
		return
	}

	n, err := roman.FromRoman(arg)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(n)
}
