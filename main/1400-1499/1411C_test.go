package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1411/C
// https://codeforces.com/problemset/status/1411/problem/C
func TestCF1411C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
3 1
2 3
3 2
2 1
1 2
5 3
2 3
3 1
1 2
5 4
4 5
5 1
2 2
3 3
outputCopy
1
3
4
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1411C)
}
