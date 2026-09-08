// Command romanconv is a thin CLI shell around the roman package. All the
// actual conversion logic lives there and is unit-tested directly; this
// file just handles argv, stdin, and stdout/stderr.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"romanconv/roman"
)

func main() {
	os.Exit(run(os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
}

// run processes one value per argument, or one value per line of stdin if
// no arguments are given, so the same binary works for both a quick
// one-off lookup and a piped batch of values. It returns the process exit
// code: 0 if every value converted cleanly, 1 if any failed.
func run(args []string, stdin *os.File, stdout, stderr *os.File) int {
	if len(args) > 0 {
		return convertAll(args, stdout, stderr)
	}

	var values []string
	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		values = append(values, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}

	if len(values) == 0 {
		fmt.Fprintf(stderr, "usage: %s <integer|roman-numeral> ...\n       or pipe values on stdin, one per line\n", os.Args[0])
		return 2
	}

	return convertAll(values, stdout, stderr)
}

func convertAll(values []string, stdout, stderr *os.File) int {
	exitCode := 0
	for _, v := range values {
		out, err := convert(v)
		if err != nil {
			fmt.Fprintln(stderr, err)
			exitCode = 1
			continue
		}
		fmt.Fprintln(stdout, out)
	}
	return exitCode
}

// convert dispatches to whichever direction of the roman package applies:
// integers go to numerals, everything else is treated as a numeral to
// parse back into an integer.
func convert(arg string) (string, error) {
	if n, err := strconv.Atoi(arg); err == nil {
		return roman.ToRoman(n)
	}

	n, err := roman.FromRoman(arg)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(n), nil
}
