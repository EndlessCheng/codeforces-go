package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1555/D
// https://codeforces.com/problemset/status/1555/problem/D
func TestCF1555D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 4
baacb
1 3
1 5
4 5
2 3
outputCopy
1
2
0
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1555D)
}
