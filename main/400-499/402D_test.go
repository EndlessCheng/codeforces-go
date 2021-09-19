package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/402/D
// https://codeforces.com/problemset/status/402/problem/D
func TestCF402D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 2
4 20 34 10 10
2 5
outputCopy
-2
inputCopy
4 5
2 4 8 16
3 5 7 11 17
outputCopy
10`
	testutil.AssertEqualCase(t, rawText, 0, CF402D)
}
