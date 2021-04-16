package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/670/C
// https://codeforces.com/problemset/status/670/problem/C
func TestCF670C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2 3 2
2
3 2
2 3
outputCopy
2
inputCopy
6
6 3 1 1 3 7
5
1 2 3 4 5
2 3 4 5 1
outputCopy
1
inputCopy
10
7 6 1 2 7 3 9 7 7 9
10
2 9 6 5 9 3 10 3 1 6
4 6 7 9 7 4 1 9 2 5
outputCopy
5`
	testutil.AssertEqualCase(t, rawText, -1, CF670C)
}
