package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1658/problem/C
// https://codeforces.com/problemset/status/1658/problem/C
func TestCF1658C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
1
1
2
1 2
2
2 2
6
1 2 4 6 3 5
6
2 3 1 2 3 4
3
3 2 1
outputCopy
YES
YES
NO
NO
YES
NO
inputCopy
1
5
1 2 3 2 4
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1658C)
}
