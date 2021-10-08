package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1515/C
// https://codeforces.com/problemset/status/1515/problem/C
func TestCF1515C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
5 2 3
1 2 3 1 2
4 3 3
1 1 2 3
outputCopy
YES
1 1 1 2 2
YES
1 2 2 3`
	testutil.AssertEqualCase(t, rawText, 0, CF1515C)
}
