package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1541/B
// https://codeforces.com/problemset/status/1541/problem/B
func TestCF1541B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2
3 1
3
6 1 5
5
3 1 5 9 2
outputCopy
1
1
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1541B)
}
