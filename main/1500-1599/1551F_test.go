package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1551/F
// https://codeforces.com/problemset/status/1551/problem/F
func TestCF1551F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3

4 2
1 2
2 3
2 4

3 3
1 2
2 3

5 3
1 2
2 3
2 4
4 5
outputCopy
6
0
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1551F)
}
