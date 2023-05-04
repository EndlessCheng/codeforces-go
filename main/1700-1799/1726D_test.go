package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1726/D
// https://codeforces.com/problemset/status/1726/problem/D
func TestCF1726D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
5 7
1 2
2 3
3 4
4 5
5 1
1 3
3 5
4 4
1 2
2 3
1 4
3 4
6 7
1 2
1 3
3 4
4 5
1 4
5 6
6 2
2 1
1 2
outputCopy
0111010
1001
0001111
0`
	testutil.AssertEqualCase(t, rawText, 0, CF1726D)
}
