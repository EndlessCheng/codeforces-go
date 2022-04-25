package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1659/B
// https://codeforces.com/problemset/status/1659/problem/B
func TestCF1659B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
6 3
100001
6 4
100011
6 0
000000
6 1
111001
6 11
101100
6 12
001110
outputCopy
111110
1 0 0 2 0 0 
111110
0 1 1 1 0 1 
000000
0 0 0 0 0 0 
100110
1 0 0 0 0 0 
111111
1 2 1 3 0 4 
111110
1 1 4 2 0 4`
	testutil.AssertEqualCase(t, rawText, 0, CF1659B)
}
