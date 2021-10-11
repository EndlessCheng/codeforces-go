package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1114/B
// https://codeforces.com/problemset/status/1114/problem/B
func TestCF1114B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
9 2 3
5 2 5 2 4 1 1 3 2
outputCopy
21
3 5 
inputCopy
6 1 4
4 1 3 2 2 3
outputCopy
12
1 3 5 
inputCopy
2 1 2
-1000000000 1000000000
outputCopy
0
1 `
	testutil.AssertEqualCase(t, rawText, 0, CF1114B)
}
