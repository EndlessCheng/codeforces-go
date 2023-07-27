package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/796/C
// https://codeforces.com/problemset/status/796/problem/C
func TestCF796C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 2 3 4 5
1 2
2 3
3 4
4 5
outputCopy
5
inputCopy
7
38 -29 87 93 39 28 -55
1 2
2 5
3 2
2 4
1 7
7 6
outputCopy
93
inputCopy
5
1 2 7 6 7
1 5
5 3
3 4
2 4
outputCopy
8
inputCopy
3
2 2 2
3 2
1 2
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, -1, CF796C)
}
