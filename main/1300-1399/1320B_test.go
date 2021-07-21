package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1320/B
// https://codeforces.com/problemset/status/1320/problem/B
func TestCF1320B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 9
1 5
5 4
1 2
2 3
3 4
4 1
2 6
6 4
4 2
4
1 2 3 4
outputCopy
1 2
inputCopy
7 7
1 2
2 3
3 4
4 5
5 6
6 7
7 1
7
1 2 3 4 5 6 7
outputCopy
0 0
inputCopy
8 13
8 7
8 6
7 5
7 4
6 5
6 4
5 3
5 2
4 3
4 2
3 1
2 1
1 8
5
8 7 5 2 1
outputCopy
0 3`
	testutil.AssertEqualCase(t, rawText, 0, CF1320B)
}
