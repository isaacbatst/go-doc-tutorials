package main

import (
	"testing"
	"unicode/utf8"
)

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, orig string) {
		rev, revErr := Reverse(orig)

		if revErr != nil {
			return
		}

		doubleRev, dRevErr := Reverse(rev)

		if dRevErr != nil {
			return
		}

		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, rev)
		}

		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTTF-8 string %q", rev)
		}
	})
}
