package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/685/B
// https://codeforces.com/problemset/status/685/problem/B
func TestCF685B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7 4
1 1 3 3 5 3
1
2
3
5
outputCopy
3
2
3
6`
	testutil.AssertEqualCase(t, rawText, 0, CF685B)
}
