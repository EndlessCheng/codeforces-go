package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1572/B
// https://codeforces.com/problemset/status/1572/problem/B
func TestCF1572B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
3
0 0 0
5
1 1 1 1 0
4
1 0 0 1
outputCopy
YES
0
YES
2
3 1
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1572B)
}
