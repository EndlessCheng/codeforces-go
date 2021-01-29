package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1477/A
// https://codeforces.com/problemset/status/1477/problem/A
func TestCF1477A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
2 1
1 2
3 0
2 3 7
2 -1
31415926 27182818
2 1000000000000000000
1 1000000000000000000
2 -1000000000000000000
-1000000000000000000 123
6 80
-5 -20 13 -14 -2 -11
outputCopy
YES
YES
NO
YES
YES
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1477A)
}
