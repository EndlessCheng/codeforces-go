package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1196/D2
// https://codeforces.com/problemset/status/1196/problem/D2
func TestCF1196D2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
5 2
BGGGG
5 3
RBRGR
5 5
BBBRR
outputCopy
1
0
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1196D2)
}
