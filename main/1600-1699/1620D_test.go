package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/1620/D
// https://codeforces.com/problemset/status/1620/problem/D
func TestCF1620D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1
1337
3
10 8 10
5
1 2 3 4 5
3
7 77 777
outputCopy
446
4
3
260`
	testutil.AssertEqualCase(t, rawText, 0, CF1620D)
}
