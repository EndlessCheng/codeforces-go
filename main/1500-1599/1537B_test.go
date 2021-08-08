package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/1537/B
// https://codeforces.com/problemset/status/1537/problem/B
func TestCF1537B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
2 3 1 1
4 4 1 2
3 5 2 2
5 1 2 1
3 1 3 1
1 1 1 1
1000000000 1000000000 1000000000 50
outputCopy
1 2 2 3
4 1 4 4
3 1 1 5
5 1 1 1
1 1 2 1
1 1 1 1
50 1 1 1000000000`
	testutil.AssertEqualCase(t, rawText, 0, CF1537B)
}
