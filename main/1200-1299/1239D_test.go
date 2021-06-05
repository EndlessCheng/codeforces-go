package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1239/problem/D
// https://codeforces.com/problemset/status/1239/problem/D
func TestCF1239D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
3 4
1 1
2 2
3 3
1 3

3 7
1 1
1 2
1 3
2 2
3 1
3 2
3 3

1 1
1 1

2 4
1 1
1 2
2 1
2 2
outputCopy
Yes
2 1
1 3 
2 
Yes
1 2
2 
1 3 
No
No`
	testutil.AssertEqualCase(t, rawText, 0, CF1239D)
}
