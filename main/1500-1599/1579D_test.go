package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1579/D
// https://codeforces.com/problemset/status/1579/problem/D
func TestCF1579D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8
2
2 3
3
1 2 3
4
1 2 3 4
3
0 0 2
2
6 2
3
0 0 2
5
8 2 0 1 1
5
0 1 0 0 6
outputCopy
2
1 2
1 2
3
1 3
2 3
2 3
5
1 3
2 4
2 4
3 4
3 4
0
2
1 2
1 2
0
4
1 2
1 5
1 4
1 2
1
5 2`
	testutil.AssertEqualCase(t, rawText, 0, CF1579D)
}
