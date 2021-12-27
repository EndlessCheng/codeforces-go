package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/777/E
// https://codeforces.com/problemset/status/777/problem/E
func TestCF777E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 5 1
2 6 2
3 7 3
outputCopy
6
inputCopy
4
1 2 1
1 3 3
4 6 2
5 7 1
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, 0, CF777E)
}
