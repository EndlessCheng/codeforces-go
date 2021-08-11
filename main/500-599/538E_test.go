package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/538/E
// https://codeforces.com/problemset/status/538/problem/E
func TestCF538E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 2
1 3
2 4
2 5
outputCopy
3 2
inputCopy
6
1 2
1 3
3 4
1 5
5 6
outputCopy
3 3`
	testutil.AssertEqualCase(t, rawText, 0, CF538E)
}
