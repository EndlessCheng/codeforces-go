package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/641/C
// https://codeforces.com/problemset/status/641/problem/C
func TestCF641C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 3
1 2
2
1 2
outputCopy
4 3 6 5 2 1
inputCopy
2 3
1 1
2
1 -2
outputCopy
1 2
inputCopy
4 2
2
1 3
outputCopy
1 4 3 2`
	testutil.AssertEqualCase(t, rawText, 0, CF641C)
}
