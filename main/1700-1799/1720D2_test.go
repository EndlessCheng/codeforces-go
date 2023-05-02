package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1720/D2
// https://codeforces.com/problemset/status/1720/problem/D2
func TestCF1720D2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2
1 2
5
5 2 4 3 1
10
3 8 8 2 9 1 6 2 8 3
outputCopy
2
3
6`
	testutil.AssertEqualCase(t, rawText, 0, CF1720D2)
}
