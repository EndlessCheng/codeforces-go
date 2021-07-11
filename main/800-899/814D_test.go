package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/814/D
// https://codeforces.com/problemset/status/814/problem/D
func TestCF814D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
2 1 6
0 4 1
2 -1 3
1 -2 1
4 -1 1
outputCopy
138.23007676
inputCopy
8
0 0 1
0 0 2
0 0 3
0 0 4
0 0 5
0 0 6
0 0 7
0 0 8
outputCopy
289.02652413`
	testutil.AssertEqualCase(t, rawText, 0, CF814D)
}
