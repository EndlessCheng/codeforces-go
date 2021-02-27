package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/758/D
// https://codeforces.com/problemset/status/758/problem/D
func TestCF758D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
13
12
outputCopy
12
inputCopy
16
11311
outputCopy
475
inputCopy
20
999
outputCopy
3789
inputCopy
17
2016
outputCopy
594
inputCopy
2
10000000000000000000000000
outputCopy
33554432`
	testutil.AssertEqualCase(t, rawText, 0, CF758D)
}
