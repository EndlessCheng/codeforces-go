package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1777/D
// https://codeforces.com/problemset/status/1777/problem/D
func TestCF1777D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1
6
1 2
1 3
3 4
3 5
3 6
outputCopy
288`
	testutil.AssertEqualCase(t, rawText, 0, CF1777D)
}
