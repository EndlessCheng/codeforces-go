package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/476/D
// https://codeforces.com/problemset/status/476/problem/D
func TestCF476D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 1
outputCopy
5
1 2 3 5
inputCopy
2 2
outputCopy
22
2 4 6 22
14 18 10 16`
	testutil.AssertEqualCase(t, rawText, 0, CF476D)
}
