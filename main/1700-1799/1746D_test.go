package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1746/D
// https://codeforces.com/problemset/status/1746/problem/D
func TestCF1746D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
5 4
1 2 1 3
6 2 1 5 7
5 3
1 2 1 3
6 6 1 4 10
outputCopy
54
56`
	testutil.AssertEqualCase(t, rawText, 0, CF1746D)
}
