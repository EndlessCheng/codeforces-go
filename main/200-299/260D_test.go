package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/260/D
// https://codeforces.com/problemset/status/260/problem/D
func TestCF260D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 3
1 2
0 5
outputCopy
3 1 3
3 2 2
inputCopy
6
1 0
0 3
1 8
0 2
0 3
0 0
outputCopy
2 3 3
5 3 3
4 3 2
1 6 0
2 1 0
inputCopy
6
0 0
1 0
0 0
1 0
0 0
1 0
outputCopy
`
	testutil.AssertEqualCase(t, rawText, 0, CF260D)
}
