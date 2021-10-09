package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1566/E
// https://codeforces.com/problemset/status/1566/problem/E
func TestCF1566E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
7
1 2
1 3
1 4
2 5
2 6
4 7
6
1 2
1 3
2 4
2 5
3 6
2
1 2
7
7 3
1 5
1 3
4 6
4 7
2 1
6
2 1
2 3
4 5
3 4
3 6
outputCopy
2
2
1
2
1`
	testutil.AssertEqualCase(t, rawText, 1, CF1566E)
}
