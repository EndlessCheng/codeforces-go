package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1801/C
// https://codeforces.com/problemset/status/1801/problem/C
func TestCF1801C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
4
5
4 9 4 6 8
1
7
2
8 6
1
1
4
2
3 4
2
1 8
2
2 8
2
7 9
outputCopy
4
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1801C)
}
