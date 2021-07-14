package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/689/E
// https://codeforces.com/problemset/status/689/problem/E
func TestCF689E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 2
1 2
1 3
2 3
outputCopy
5
inputCopy
3 3
1 3
1 3
1 3
outputCopy
3
inputCopy
3 1
1 2
2 3
3 4
outputCopy
6`
	testutil.AssertEqualCase(t, rawText, 0, CF689E)
}
