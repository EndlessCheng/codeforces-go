package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/580/B
// https://codeforces.com/problemset/status/580/problem/B
func TestCF580B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 5
75 5
0 100
150 20
75 1
outputCopy
100
inputCopy
5 100
0 7
11 32
99 10
46 8
87 54
outputCopy
111
inputCopy
4 2
10909234 9
10909236 8
10909237 10
10909235 98
outputCopy
107`
	testutil.AssertEqualCase(t, rawText, 0, CF580B)
}
