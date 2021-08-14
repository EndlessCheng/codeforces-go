package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1542/B
// https://codeforces.com/problemset/status/1542/problem/B
func TestCF1542B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
24 3 5
10 3 6
2345 1 4
19260817 394 485
19260817 233 264
outputCopy
Yes
No
Yes
No
Yes`
	testutil.AssertEqualCase(t, rawText, 0, CF1542B)
}
