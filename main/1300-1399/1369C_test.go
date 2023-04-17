package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1369/C
// https://codeforces.com/problemset/status/1369/problem/C
func TestCF1369C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
4 2
1 13 7 17
1 3
6 2
10 10 10 10 11 11
3 3
4 4
1000000000 1000000000 1000000000 1000000000
1 1 1 1
outputCopy
48
42
8000000000`
	testutil.AssertEqualCase(t, rawText, 0, CF1369C)
}
