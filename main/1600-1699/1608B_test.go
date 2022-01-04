package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1608/B
// https://codeforces.com/problemset/status/1608/problem/B
func TestCF1608B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
4 1 1
6 1 2
6 4 0
outputCopy
1 3 2 4
4 2 3 1 5 6
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF1608B)
}
