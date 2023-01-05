package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1744/E2
// https://codeforces.com/problemset/status/1744/problem/E2
func TestCF1744E2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
10
1 1 2 2
3 4 5 7
8 9 15 18
12 21 14 24
36 60 48 66
1024 729 373248 730
1024 729 373247 730
5040 40320 40319 1000000000
999999999 999999999 1000000000 1000000000
268435456 268435456 1000000000 1000000000
outputCopy
2 2
4 6
12 12
-1 -1
-1 -1
373248 730
-1 -1
15120 53760
-1 -1
536870912 536870912`
	testutil.AssertEqualCase(t, rawText, 0, CF1744E2)
}
