package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/863/E
// https://codeforces.com/problemset/status/863/problem/E
func TestCF863E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 3
4 6
1 7
outputCopy
1
inputCopy
2
0 10
0 10
outputCopy
1
inputCopy
3
1 2
3 4
6 8
outputCopy
-1
inputCopy
3
1 2
2 3
3 4
outputCopy
2
inputCopy
3
1 4
2 100
4 5
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF863E)
}
