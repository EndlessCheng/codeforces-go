package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1843/problem/E
// https://codeforces.com/problemset/status/1843/problem/E
func TestCF1843E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
5 5
1 2
4 5
1 5
1 3
2 4
5
5
3
1
2
4
4 2
1 1
4 4
2
2
3
5 2
1 5
1 5
4
2
1
3
4
5 2
1 5
1 3
5
4
1
2
3
5
5 5
1 5
1 5
1 5
1 5
1 4
3
1
4
3
3 2
2 2
1 3
3
2
3
1
outputCopy
3
-1
3
3
3
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1843E)
}
