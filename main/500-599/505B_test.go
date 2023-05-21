package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/505/B
// https://codeforces.com/problemset/status/505/problem/B
func TestCF505B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 5
1 2 1
1 2 2
2 3 1
2 3 3
2 4 3
3
1 2
3 4
1 4
outputCopy
2
1
0
inputCopy
5 7
1 5 1
2 5 1
3 5 1
4 5 1
1 2 2
2 3 2
3 4 2
5
1 5
5 1
2 5
1 5
1 4
outputCopy
1
1
1
1
2`
	testutil.AssertEqualCase(t, rawText, 0, CF505B)
}
