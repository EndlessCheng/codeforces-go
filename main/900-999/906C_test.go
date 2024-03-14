package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/906/problem/C
// https://codeforces.com/problemset/status/906/problem/C
func TestCF906C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 6
1 2
1 3
2 3
2 5
3 4
4 5
outputCopy
2
2 3 
inputCopy
4 4
1 2
1 3
1 4
3 4
outputCopy
1
1 `
	testutil.AssertEqualCase(t, rawText, 0, CF906C)
}
