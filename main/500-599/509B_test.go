package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/509/B
// https://codeforces.com/problemset/status/509/problem/B
func TestCF509B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 4
1 2 3 4
outputCopy
YES
1
1 4
1 2 4
1 2 3 4
inputCopy
5 2
3 2 4 1 3
outputCopy
NO
inputCopy
5 4
3 2 4 3 5
outputCopy
YES
1 2 3
1 3
1 2 3 4
1 3 4
1 1 2 3 4`
	testutil.AssertEqualCase(t, rawText, 0, CF509B)
}
