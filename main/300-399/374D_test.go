package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/374/D
// https://codeforces.com/problemset/status/374/problem/D
func TestCF374D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
10 3
1 3 6
-1
1
1
0
0
-1
0
1
-1
1
outputCopy
011
inputCopy
2 1
1
1
-1
outputCopy
Poor stack!`
	testutil.AssertEqualCase(t, rawText, 0, CF374D)
}
