package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/56/E
// https://codeforces.com/problemset/status/56/problem/E
func TestCF56E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
16 5
20 5
10 10
18 2
outputCopy
3 1 4 1 
inputCopy
4
0 10
1 5
9 10
15 10
outputCopy
4 1 2 1 `
	testutil.AssertEqualCase(t, rawText, 0, CF56E)
}
