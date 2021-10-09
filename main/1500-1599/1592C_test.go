package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1592/C
// https://codeforces.com/problemset/status/1592/problem/C
func TestCF1592C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
2 2
1 3
1 2
5 5
3 3 3 3 3
1 2
2 3
1 4
4 5
5 2
1 7 2 3 5
1 2
2 3
1 4
4 5
5 3
1 6 4 1 2
1 2
2 3
1 4
4 5
3 3
1 7 4
1 2
2 3
outputCopy
NO
YES
NO
YES
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1592C)
}
