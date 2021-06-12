package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1037/E
// https://codeforces.com/problemset/status/1037/problem/E
func TestCF1037E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 4 2
2 3
1 2
1 3
1 4
outputCopy
0
0
3
3
inputCopy
5 8 2
2 1
4 2
5 4
5 2
4 3
5 1
4 1
3 2
outputCopy
0
0
0
3
3
4
4
5
inputCopy
5 7 2
1 5
3 2
2 5
3 4
1 2
5 3
1 3
outputCopy
0
0
0
0
3
4
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1037E)
}
