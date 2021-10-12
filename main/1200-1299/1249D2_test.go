package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1249/D2
// https://codeforces.com/problemset/status/1249/problem/D2
func TestCF1249D2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7 2
11 11
9 11
7 8
8 9
7 8
9 11
7 9
outputCopy
3
4 6 7 
inputCopy
5 1
29 30
30 30
29 29
28 30
30 30
outputCopy
3
1 4 5 
inputCopy
6 1
2 3
3 3
2 3
2 2
2 3
2 3
outputCopy
4
1 3 5 6 `
	testutil.AssertEqualCase(t, rawText, 0, CF1249D2)
}
