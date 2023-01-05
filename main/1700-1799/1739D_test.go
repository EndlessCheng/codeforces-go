package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1739/D
// https://codeforces.com/problemset/status/1739/problem/D
func TestCF1739D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
5 1
1 1 2 2
5 2
1 1 2 2
6 0
1 2 3 4 5
6 1
1 2 3 4 5
4 3
1 1 1
outputCopy
2
1
5
3
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1739D)
}
