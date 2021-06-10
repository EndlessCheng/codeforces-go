package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1220/problem/E
// https://codeforces.com/problemset/status/1220/problem/E
func TestCF1220E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 7
2 2 8 6 9
1 2
1 3
2 4
3 2
4 5
2 5
1 5
2
outputCopy
27
inputCopy
10 12
1 7 1 9 3 3 6 30 1 10
1 2
1 3
3 5
5 7
2 3
5 4
6 9
4 6
3 7
6 8
9 4
9 10
6
outputCopy
61
inputCopy
1 0
1000000000
1
outputCopy
1000000000
inputCopy
8 9
1 7 1 9 3 3 6 30
1 2
1 3
3 5
5 7
2 3
5 4
4 6
3 7
6 8
6
outputCopy
60`
	testutil.AssertEqualCase(t, rawText, 0, CF1220E)
}
