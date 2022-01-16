package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/357/B
// https://codeforces.com/problemset/status/357/problem/B
func TestCF357B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7 3
1 2 3
1 4 5
4 6 7
outputCopy
1 2 3 3 2 2 1 
inputCopy
9 3
3 6 9
2 5 8
1 4 7
outputCopy
1 1 1 2 2 2 3 3 3 
inputCopy
5 2
4 1 5
3 1 2
outputCopy
2 3 1 1 3 `
	testutil.AssertEqualCase(t, rawText, 0, CF357B)
}
