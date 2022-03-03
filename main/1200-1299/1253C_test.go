package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1253/C
// https://codeforces.com/problemset/status/1253/problem/C
func TestCF1253C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
9 2
6 19 3 4 4 2 6 7 8
outputCopy
2 5 11 18 30 43 62 83 121
inputCopy
1 1
7
outputCopy
7`
	testutil.AssertEqualCase(t, rawText, 0, CF1253C)
}
