package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/959/C
// https://codeforces.com/problemset/status/959/problem/C
func TestCF959C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
outputCopy
-1
1 2
inputCopy
8
outputCopy
1 2
1 3
2 4
2 5
3 6
4 7
4 8
1 2
1 3
2 4
2 5
2 6
3 7
6 8`
	testutil.AssertEqualCase(t, rawText, 0, CF959C)
}
