package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1824/problem/B2
// https://codeforces.com/problemset/status/1824/problem/B2
func TestCF1824B2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 2
1 2
2 3
3 4
outputCopy
666666674
inputCopy
5 5
1 2
2 3
3 4
3 5
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1824B2)
}
