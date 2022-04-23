package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1671/B
// https://codeforces.com/problemset/status/1671/problem/B
func TestCF1671B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
2
1 4
3
1 2 3
4
1 2 3 7
1
1000000
3
2 5 6
outputCopy
YES
YES
NO
YES
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF1671B)
}
