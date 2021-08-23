package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1540/B
// https://codeforces.com/problemset/status/1540/problem/B
func TestCF1540B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 2
1 3
outputCopy
166666669
inputCopy
6
2 1
2 3
6 1
1 4
2 5
outputCopy
500000009
inputCopy
5
1 2
1 3
1 4
2 5
outputCopy
500000007`
	testutil.AssertEqualCase(t, rawText, 0, CF1540B)
}
