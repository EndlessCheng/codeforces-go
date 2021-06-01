package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1285/E
// https://codeforces.com/problemset/status/1285/problem/E
func TestCF1285E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
4
1 4
2 3
3 6
5 7
3
5 5
5 5
5 5
6
3 3
1 1
5 5
1 5
2 2
4 4
outputCopy
2
1
5
inputCopy
1
2
3 4
-3 2
outputCopy
1
inputCopy
1
2
1 2
2 3
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1285E)
}
