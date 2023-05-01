package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1817/problem/B
// https://codeforces.com/problemset/status/1817/problem/B
func TestCF1817B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
7 8
1 2
2 3
3 4
4 1
4 5
4 6
4 2
6 7
7 7
6 7
1 2
2 3
3 4
4 1
1 3
3 5
4 4
1 3
3 4
4 1
1 2
outputCopy
YES
6
5 4
6 4
4 3
1 4
2 1
3 2
YES
5
5 3
2 3
3 1
4 3
1 4
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1817B)
}
