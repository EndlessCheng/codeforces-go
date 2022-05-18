package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/815/A
// https://codeforces.com/problemset/status/815/problem/A
func TestCF815A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 5
2 2 2 3 2
0 0 0 1 0
1 1 1 2 1
outputCopy
4
row 1
row 1
col 4
row 3
inputCopy
3 3
0 0 0
0 1 0
0 0 0
outputCopy
-1
inputCopy
3 3
1 1 1
1 1 1
1 1 1
outputCopy
3
row 1
row 2
row 3
inputCopy
3 2
1 1
2 2
2 2
outputCopy
`
	testutil.AssertEqualCase(t, rawText, 0, CF815A)
}
