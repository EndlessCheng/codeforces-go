package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/814/B
// https://codeforces.com/problemset/status/814/problem/B
func TestCF814B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 2 3 4 3
1 2 5 4 5
outputCopy
1 2 5 4 3
inputCopy
5
4 4 2 3 1
5 4 5 3 1
outputCopy
5 4 2 3 1
inputCopy
4
1 1 3 4
1 4 3 4
outputCopy
1 2 3 4`
	testutil.AssertEqualCase(t, rawText, 0, CF814B)
}
