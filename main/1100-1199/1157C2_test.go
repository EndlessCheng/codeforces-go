package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1157/C2
// https://codeforces.com/problemset/status/1157/problem/C2
func TestCF1157C2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 2 4 3 2
outputCopy
4
LRRR
inputCopy
7
1 3 5 6 5 4 2
outputCopy
6
LRLRRR
inputCopy
3
2 2 2
outputCopy
1
R
inputCopy
4
1 2 4 3
outputCopy
4
LLRR`
	testutil.AssertEqualCase(t, rawText, 0, CF1157C2)
}
