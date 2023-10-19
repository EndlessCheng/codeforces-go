package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1443/C
// https://codeforces.com/problemset/status/1443/problem/C
func TestCF1443C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
4
3 7 4 5
2 1 2 4
4
1 2 3 4
3 3 3 3
2
1 2
10 10
2
10 10
1 2
outputCopy
5
3
2
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1443C)
}
