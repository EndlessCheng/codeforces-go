package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1276/C
// https://codeforces.com/problemset/status/1276/problem/C
func TestCF1276C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
12
3 1 4 1 5 9 2 6 5 3 5 8
outputCopy
12
3 4
1 2 3 5
3 1 5 4
5 6 8 9
inputCopy
5
1 1 1 1 1
outputCopy
1
1 1
1
inputCopy
4
1 2 3 3
outputCopy
4
2 2
3 1
2 3
inputCopy
15
6 7 11 8 13 11 4 20 17 12 9 15 18 13 9
outputCopy
15
3 5
4 11 17 8 13
13 6 11 18 9
9 15 7 12 20`
	testutil.AssertEqualCase(t, rawText, 0, CF1276C)
}
