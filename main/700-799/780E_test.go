package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/780/E
// https://codeforces.com/problemset/status/780/problem/E
func TestCF780E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 2 1
2 1
3 1
outputCopy
3 2 1 3
inputCopy
5 4 2
1 2
1 3
1 4
1 5
outputCopy
3 2 1 3
3 4 1 5`
	testutil.AssertEqualCase(t, rawText, 0, CF780E)
}