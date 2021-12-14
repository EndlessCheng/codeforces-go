package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1607/E
// https://codeforces.com/problemset/status/1607/problem/E
func TestCF1607E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 1
L
1 2
L
3 3
RRDLUU
4 3
LUURRDDLLLUU
outputCopy
1 1
1 2
2 1
3 2`
	testutil.AssertEqualCase(t, rawText, 0, CF1607E)
}
