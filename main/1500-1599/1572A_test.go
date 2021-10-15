package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1572/A
// https://codeforces.com/problemset/status/1572/problem/A
func TestCF1572A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
4
1 2
0
2 1 4
1 2
5
1 5
1 1
1 2
1 3
1 4
5
0
0
2 1 2
1 2
2 2 1
4
2 2 3
0
0
2 3 2
5
1 2
1 3
1 4
1 5
0
outputCopy
2
-1
1
2
5`
	testutil.AssertEqualCase(t, rawText, 0, CF1572A)
}
