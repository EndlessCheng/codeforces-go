package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/337/D
// https://codeforces.com/problemset/status/337/problem/D
func TestCF337D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 2 3
1 2
1 5
2 3
3 4
4 5
5 6
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF337D)
}
