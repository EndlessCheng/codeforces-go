package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/475/D
// https://codeforces.com/problemset/status/475/problem/D
func TestCF475D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2 6 3
5
1
2
3
4
6
outputCopy
1
2
2
0
1
inputCopy
7
10 20 3 15 1000 60 16
10
1
2
3
4
5
6
10
20
60
1000
outputCopy
14
0
2
2
2
0
2
2
1
1`
	testutil.AssertEqualCase(t, rawText, 0, CF475D)
}
