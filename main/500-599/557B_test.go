package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/557/B
// https://codeforces.com/problemset/status/557/problem/B
func TestCF557B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 4
1 1 1 1
outputCopy
3
inputCopy
3 18
4 4 4 2 2 2
outputCopy
18
inputCopy
1 5
2 3
outputCopy
4.5`
	testutil.AssertEqualCase(t, rawText, 0, CF557B)
}
