package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1066/E
// https://codeforces.com/problemset/status/1066/problem/E
func TestCF1066E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 4
1010
1101
outputCopy
12
inputCopy
4 5
1001
10101
outputCopy
11`
	testutil.AssertEqualCase(t, rawText, 0, CF1066E)
}
