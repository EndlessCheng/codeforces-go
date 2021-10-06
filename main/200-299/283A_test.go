package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/283/A
// https://codeforces.com/problemset/status/283/problem/A
func TestCF283A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
2 1
3
2 3
2 1
3
outputCopy
0.500000
0.000000
1.500000
1.333333
1.500000
inputCopy
6
2 1
1 2 20
2 2
1 2 -3
3
3
outputCopy
0.500000
20.500000
14.333333
12.333333
17.500000
17.000000`
	testutil.AssertEqualCase(t, rawText, 0, CF283A)
}
