package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1136/D
// https://codeforces.com/problemset/status/1136/problem/D
func TestCF1136D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 1
1 2
1 2
outputCopy
1
inputCopy
3 3
3 1 2
1 2
3 1
3 2
outputCopy
2
inputCopy
5 2
3 1 5 4 2
5 2
5 4
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, -1, CF1136D)
}
