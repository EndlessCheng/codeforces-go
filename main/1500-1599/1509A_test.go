package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1509/A
// https://codeforces.com/problemset/status/1509/problem/A
func TestCF1509A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
3
1 1 2
3
1 1 1
8
10 9 13 15 3 16 9 13
2
18 9
outputCopy
1 1 2 
1 1 1 
13 9 13 15 3 9 16 10 
9 18 `
	testutil.AssertEqualCase(t, rawText, 0, CF1509A)
}
