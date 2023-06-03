package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/427/C
// https://codeforces.com/problemset/status/427/problem/C
func TestCF427C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
1 2 3
3
1 2
2 3
3 2
outputCopy
3 1
inputCopy
5
2 8 0 6 0
6
1 4
1 3
2 4
3 4
4 5
5 1
outputCopy
8 2
inputCopy
10
1 3 2 2 1 3 1 4 10 10
12
1 2
2 3
3 1
3 4
4 5
5 6
5 7
6 4
7 3
8 9
9 10
10 9
outputCopy
15 6
inputCopy
2
7 91
2
1 2
2 1
outputCopy
7 1`
	testutil.AssertEqualCase(t, rawText, 0, CF427C)
}
