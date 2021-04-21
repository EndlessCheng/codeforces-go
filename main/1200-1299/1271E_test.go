package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1271/E
// https://codeforces.com/problemset/status/1271/problem/E
func TestCF1271E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
11 3
outputCopy
5
inputCopy
11 6
outputCopy
4
inputCopy
20 20
outputCopy
1
inputCopy
14 5
outputCopy
6
inputCopy
1000000 100
outputCopy
31248`
	testutil.AssertEqualCase(t, rawText, 0, CF1271E)
}
