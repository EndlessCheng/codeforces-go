package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1100/E
// https://codeforces.com/problemset/status/1100/problem/E
func TestCF1100E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 6
2 1 1
5 2 6
2 3 2
3 4 3
4 5 5
1 5 4
outputCopy
2 2
1 3 
inputCopy
5 7
2 1 5
3 2 3
1 3 3
2 4 1
4 3 5
5 4 1
1 5 3
outputCopy
3 3
3 4 7 `
	testutil.AssertEqualCase(t, rawText, 0, CF1100E)
}
