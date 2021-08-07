package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1012/B
// https://codeforces.com/problemset/status/1012/problem/B
func TestCF1012B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 2 3
1 2
2 2
2 1
outputCopy
0
inputCopy
1 5 3
1 3
1 1
1 5
outputCopy
2
inputCopy
4 3 6
1 2
1 3
2 2
2 3
3 1
3 3
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1012B)
}
