package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1051/C
// https://codeforces.com/problemset/status/1051/problem/C
func TestCF1051C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
3 5 7 1
outputCopy
YES
BABA
inputCopy
3
3 5 1
outputCopy
NO
inputCopy
100
7 1 5 50 49 14 30 42 44 46 10 32 17 21 36 26 46 21 48 9 5 24 33 34 33 4 4 39 3 50 39 41 28 18 29 20 13 2 44 9 31 35 8 7 13 28 50 10 11 30 38 17 11 15 40 8 26 18 19 23 42 20 24 31 43 48 29 12 19 12 36 45 16 40 37 47 6 2 45 22 16 27 37 27 34 15 43 22 38 35 23 14 47 6 25 41 1 32 25 3
outputCopy
YES
AAABAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA`
	testutil.AssertEqualCase(t, rawText, -1, CF1051C)
}
