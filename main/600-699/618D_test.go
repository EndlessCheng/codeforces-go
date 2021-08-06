package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/618/D
// https://codeforces.com/problemset/status/618/problem/D
func TestCF618D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 2 3
1 2
1 3
3 4
5 3
outputCopy
9
inputCopy
5 3 2
1 2
1 3
3 4
5 3
outputCopy
8`
	testutil.AssertEqualCase(t, rawText, 0, CF618D)
}
