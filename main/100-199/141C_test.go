package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/141/C
// https://codeforces.com/problemset/status/141/problem/C
func TestCF141C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
a 0
b 2
c 0
d 0
outputCopy
a 150
c 170
d 180
b 160
inputCopy
4
vasya 0
petya 1
manya 3
dunay 3
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF141C)
}
