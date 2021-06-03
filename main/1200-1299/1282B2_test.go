package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1282/B2
// https://codeforces.com/problemset/status/1282/problem/B2
func TestCF1282B2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8
5 6 2
2 4 3 5 7
5 11 2
2 4 3 5 7
3 2 3
4 2 6
5 2 3
10 1 3 9 2
2 10000 2
10000 10000
2 9999 2
10000 10000
4 6 4
3 2 3 2
5 5 3
1 2 2 1 2
outputCopy
3
4
1
1
2
0
4
5`
	testutil.AssertEqualCase(t, rawText, 0, CF1282B2)
}
