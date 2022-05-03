package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1674/E
// https://codeforces.com/problemset/status/1674/problem/E
func TestCF1674E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
20 10 30 10 20
outputCopy
10
inputCopy
3
1 8 1
outputCopy
1
inputCopy
6
7 6 6 8 5 8
outputCopy
4
inputCopy
6
14 3 8 10 15 4
outputCopy
4
inputCopy
4
1 100 100 1
outputCopy
2
inputCopy
3
40 10 10
outputCopy
7
inputCopy
3
1 100 3
outputCopy
2`
	testutil.AssertEqualCase(t, rawText, 0, CF1674E)
}
