package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/555/B
// https://codeforces.com/problemset/status/555/problem/B
func TestCF555B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 4
1 4
7 8
9 10
12 14
4 5 3 8
outputCopy
Yes
2 3 1 
inputCopy
2 2
11 14
17 18
2 9
outputCopy
No
inputCopy
2 1
1 1
1000000000000000000 1000000000000000000
999999999999999999
outputCopy
Yes
1 
inputCopy
2 1
1 2
5 6
1
outputCopy
No`
	testutil.AssertEqualCase(t, rawText, 1, CF555B)
}
