package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1506/F
// https://codeforces.com/problemset/status/1506/problem/F
func TestCF1506F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
3
1 4 2
1 3 1
2
2 4
2 3
2
1 1000000000
1 1000000000
4
3 10 5 8
2 5 2 4
outputCopy
0
1
999999999
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1506F)
}
