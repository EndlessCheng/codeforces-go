package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1627/C
// https://codeforces.com/problemset/status/1627/problem/C
func TestCF1627C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2
1 2
4
1 3
4 3
2 1
7
1 2
1 3
3 4
3 5
6 2
7 2
outputCopy
17
2 5 11
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF1627C)
}
