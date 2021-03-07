package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1118/C
// https://codeforces.com/problemset/status/1118/problem/C
func TestCF1118C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 8 8 1 2 2 2 2 2 2 2 2 1 8 8 1
outputCopy
YES
1 2 2 1
8 2 2 8
8 2 2 8
1 2 2 1
inputCopy
3
1 1 1 1 1 3 3 3 3
outputCopy
YES
1 3 1
3 1 3
1 3 1
inputCopy
4
1 2 1 9 8 4 3 8 8 3 4 8 9 2 1 1
outputCopy
NO
inputCopy
1
10
outputCopy
YES
10 `
	testutil.AssertEqualCase(t, rawText, 0, CF1118C)
}
