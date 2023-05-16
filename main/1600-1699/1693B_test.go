package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1693/B
// https://codeforces.com/problemset/status/1693/problem/B
func TestCF1693B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
2
1
1 5
2 9
3
1 1
4 5
2 4
6 10
4
1 2 1
6 9
5 6
4 5
2 4
5
1 2 3 4
5 5
4 4
3 3
2 2
1 1
outputCopy
1
2
2
5`
	testutil.AssertEqualCase(t, rawText, 0, CF1693B)
}
