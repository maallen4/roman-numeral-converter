# romanconv

A converter between integers and roman numerals, written as a small pure
Go package plus a CLI wrapper.

## The problem

Encoding an integer as a roman numeral is easy — everyone's seen the
greedy-symbol-subtraction algorithm. Decoding is where implementations get
sloppy: a naive left-to-right sum will happily accept garbage like `IIII`,
`VX`, or `IC` and produce a number for them, because summing digit values
doesn't know those aren't real roman numerals. `MCMXCIV` is 1994; `IC` is
not "99", it's not a numeral at all.

This package treats "parse" and "validate" as the same operation: `FromRoman`
computes a candidate value, re-encodes it with `ToRoman`, and only accepts
the input if that round trip reproduces the original string exactly. That
one check replaces a whole second copy of the grammar rules (repetition
limits, which pairs may subtract, ordering) that would otherwise have to be
maintained separately from the encoder and kept in sync with it by hand.

## Usage

As a library:

```go
import "romanconv/roman"

s, err := roman.ToRoman(1994)   // "MCMXCIV", nil
n, err := roman.FromRoman("XL") // 40, nil

_, err = roman.FromRoman("IIII") // error: not canonical
_, err = roman.ToRoman(4000)     // error: out of range
```

From the command line:

```
$ go run ./cmd/romanconv 1994
MCMXCIV

$ go run ./cmd/romanconv MCMXCIV
1994

$ go run ./cmd/romanconv IIII
roman: "IIII" is not canonical, expected "IV"
```

## Design

Both `roman.ToRoman` and `roman.FromRoman` are pure functions: no I/O, no
package-level mutable state, deterministic output for a given input. That
was the one thing this project set out to get right, since a converter with
side effects or hidden state is far more annoying to write tests against
than one without.

Valid range is 1 through 3999 — the numerals the classical symbol set
(I, V, X, L, C, D, M) can represent without inventing extra notation.

## Status

Early skeleton. Encoding and decoding both work and are covered by
table-driven tests, a full round-trip check over the entire valid range,
and fuzz tests (`go test -fuzz=FuzzToRoman ./roman` /
`-fuzz=FuzzFromRoman`) for inputs the tables don't enumerate.
