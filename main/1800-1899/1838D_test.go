package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1838/problem/D
// https://codeforces.com/problemset/status/1838/problem/D
func TestCF1838D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
10 9
(())()()))
9
7
2
6
3
6
7
4
8
outputCopy
YES
YES
NO
NO
YES
NO
YES
NO
NO
inputCopy
3 2
(()
2
3
outputCopy
NO
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1838D)
}
