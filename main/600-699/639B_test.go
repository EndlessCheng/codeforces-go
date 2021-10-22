package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/639/B
// https://codeforces.com/problemset/status/639/problem/B
func TestCF639B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 3 2
outputCopy
1 2
1 3
3 4
3 5
inputCopy
8 5 2
outputCopy
-1
inputCopy
8 4 2
outputCopy
4 8
5 7
2 3
8 1
2 1
5 6
1 5
inputCopy
10 3 3
outputCopy
1 2
2 3
3 4
5 2
6 2
7 2
8 2
9 2
10 2
inputCopy
3 1 1
outputCopy
-1
inputCopy
4 2 1
outputCopy
1 2
1 3
4 1`
	testutil.AssertEqualCase(t, rawText, 0, CF639B)
}
