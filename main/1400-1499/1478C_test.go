package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1478/C
// https://codeforces.com/problemset/status/1478/problem/C
func TestCF1478C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
2
8 12 8 12
2
7 7 9 11
2
7 11 7 11
1
1 1
4
40 56 48 40 80 56 80 48
6
240 154 210 162 174 154 186 240 174 186 162 210
outputCopy
YES
NO
NO
NO
NO
YES
inputCopy
1
2
4 4 4 4
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, -1, CF1478C)
}
