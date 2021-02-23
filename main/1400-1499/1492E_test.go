package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1492/E
// https://codeforces.com/problemset/status/1492/problem/E
func TestCF1492E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 4
1 10 10 100
1 1 1 100
10 100 1 100
outputCopy
Yes
1 10 1 100
inputCopy
10 7
1 1 1 1 1 1 1
1 1 1 1 1 1 2
1 1 1 1 1 2 2
1 1 1 1 2 2 1
1 1 1 2 2 1 1
1 1 2 2 1 1 1
1 2 2 1 1 1 1
2 2 1 1 1 1 1
2 1 1 1 1 1 1
1 1 1 1 1 1 1
outputCopy
Yes
1 1 1 1 1 1 1
inputCopy
2 5
2 2 1 1 1
1 1 2 2 2
outputCopy
No`
	testutil.AssertEqualCase(t, rawText, 0, CF1492E)
}
