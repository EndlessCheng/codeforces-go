package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1612/A
// https://codeforces.com/problemset/status/1612/problem/A
func TestCF1612A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
10
49 3
2 50
13 0
0 41
42 0
0 36
13 37
42 16
42 13
0 0
outputCopy
23 3
1 25
-1 -1
-1 -1
21 0
0 18
13 12
25 4
-1 -1
0 0`
	testutil.AssertEqualCase(t, rawText, 0, CF1612A)
}
