package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1567/D
// https://codeforces.com/problemset/status/1567/problem/D
func TestCF1567D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
97 2
17 1
111 4
100 2
10 9
999999 3
outputCopy
70 27 
17 
3 4 100 4
10 90
1 1 2 1 1 1 1 1 1 
999900 90 9`
	testutil.AssertEqualCase(t, rawText, 0, CF1567D)
}
