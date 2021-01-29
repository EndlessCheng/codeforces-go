package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1478/B
// https://codeforces.com/problemset/status/1478/problem/B
func TestCF1478B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
3 7
24 25 27
10 7
51 52 53 54 55 56 57 58 59 60
outputCopy
YES
NO
YES
YES
YES
NO
YES
YES
YES
YES
YES
YES
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1478B)
}
