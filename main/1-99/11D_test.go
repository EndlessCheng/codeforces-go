package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/11/D
// https://codeforces.com/problemset/status/11/problem/D
func TestCF11D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 6
1 2
1 3
1 4
2 3
2 4
3 4
outputCopy
7`
	testutil.AssertEqualCase(t, rawText, 0, CF11D)
}
