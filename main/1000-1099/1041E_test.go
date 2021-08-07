package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1041/E
// https://codeforces.com/problemset/status/1041/problem/E
func TestCF1041E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
3 4
1 4
3 4
outputCopy
YES
1 3
3 2
2 4
inputCopy
3
1 3
1 3
outputCopy
NO
inputCopy
3
1 2
2 3
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1041E)
}
