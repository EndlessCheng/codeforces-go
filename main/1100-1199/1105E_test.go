package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1105/E
// https://codeforces.com/problemset/status/1105/problem/E
func TestCF1105E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 3
1
2 motarack
2 mike
1
2 light
outputCopy
2
inputCopy
4 3
1
2 alice
2 bob
2 tanyaromanova
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1105E)
}
