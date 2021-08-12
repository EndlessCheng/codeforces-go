package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 1227F2

// https://codeforces.com/problemset/problem/1261/D2
// https://codeforces.com/problemset/status/1261/problem/D2
func TestCF1261D2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3
1 3 1
outputCopy
9
inputCopy
5 5
1 1 4 2 2
outputCopy
1000
inputCopy
6 2
1 1 2 2 1 1
outputCopy
16`
	testutil.AssertEqualCase(t, rawText, 0, CF1261D2)
}
