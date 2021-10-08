package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1539/D
// https://codeforces.com/problemset/status/1539/problem/D
func TestCF1539D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
3 4
1 3
1 5
outputCopy
8
inputCopy
5
2 7
2 8
1 2
2 4
1 8
outputCopy
12`
	testutil.AssertEqualCase(t, rawText, 1, CF1539D)
}
