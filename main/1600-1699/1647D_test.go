package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1647/D
// https://codeforces.com/problemset/status/1647/problem/D
func TestCF1647D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8
6 2
12 2
36 2
8 2
1000 10
2376 6
128 4
16384 4
outputCopy
NO
NO
YES
NO
YES
YES
NO
YES`
	testutil.AssertEqualCase(t, rawText, 0, CF1647D)
}
