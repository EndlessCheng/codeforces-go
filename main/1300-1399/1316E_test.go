package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1316/E
// https://codeforces.com/problemset/status/1316/problem/E
func TestCF1316E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 1 2
1 16 10 3
18
19
13
15
outputCopy
44
inputCopy
6 2 3
78 93 9 17 13 78
80 97
30 52
26 17
56 68
60 36
84 55
outputCopy
377
inputCopy
3 2 1
500 498 564
100002 3
422332 2
232323 1
outputCopy
422899`
	testutil.AssertEqualCase(t, rawText, 0, CF1316E)
}
