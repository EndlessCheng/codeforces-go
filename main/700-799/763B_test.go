package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/763/B
// https://codeforces.com/problemset/status/763/problem/B
func TestCF763B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8
0 0 5 3
2 -1 5 0
-3 -4 2 -1
-1 -1 2 0
-3 0 0 5
5 2 10 3
7 -3 10 2
4 -2 7 -1
outputCopy
YES
1
2
2
3
2
2
4
1`
	testutil.AssertEqualCase(t, rawText, 0, CF763B)
}
