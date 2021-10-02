package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/489/E
// https://codeforces.com/problemset/status/489/problem/E
func TestCF489E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 9
10 10
20 10
30 1
31 5
40 10
outputCopy
1 2 4 5 
inputCopy
2 7
1 9
5 6
outputCopy
2 `
	testutil.AssertEqualCase(t, rawText, -1, CF489E)
}
