package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1175/E
// https://codeforces.com/problemset/status/1175/problem/E
func TestCF1175E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 3
1 3
2 4
1 3
1 4
3 4
outputCopy
1
2
1
inputCopy
3 4
1 3
1 3
4 5
1 2
1 3
1 4
1 5
outputCopy
1
1
-1
-1`
	testutil.AssertEqualCase(t, rawText, 1, CF1175E)
}
