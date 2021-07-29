package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1223/E
// https://codeforces.com/problemset/status/1223/problem/E
func TestCF1223E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
4 1
1 2 5
3 1 2
3 4 3
7 2
1 2 5
1 3 4
1 4 2
2 5 1
2 6 2
4 7 3
outputCopy
8
14
inputCopy
1
7 2
1 2 5
1 3 4
1 4 2
2 5 1
2 6 2
4 7 3
outputCopy
14`
	testutil.AssertEqualCase(t, rawText, -1, CF1223E)
}
