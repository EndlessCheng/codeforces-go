package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1584/E
// https://codeforces.com/problemset/status/1584/problem/E
func TestCF1584E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
2
2 2
3
1 2 3
4
1 1 1 1
4
1 2 2 1
4
1 2 1 2
8
1 2 1 2 1 2 1 2
outputCopy
1
0
4
2
1
3
inputCopy
9
1
0
1
1000000000
8
1 2 1 2 1 2 1 2
2
1000000000 1000000000
11
1 1 1 1 0 1 0 1 2 1 2
11
1 2 1 2 0 1 0 1 1 1 1
7
1 3 2 2 2 2 2
9
1 5 1 1 2 1 1 7 1
7
1 2 3 4 5 6 3
outputCopy
1
0
3
1
10
9
9
3
1
inputCopy
1
11
1 1 1 1 0 1 0 1 2 1 2
outputCopy
10`
	testutil.AssertEqualCase(t, rawText, -1, CF1584E)
}
