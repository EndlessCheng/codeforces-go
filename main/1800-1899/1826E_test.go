package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1826/problem/E
// https://codeforces.com/problemset/status/1826/problem/E
func TestCF1826E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 5
10 10 10 10 10
1 2 3 4 5
1 5 2 3 4
2 3 4 5 1
outputCopy
30
inputCopy
3 5
10 10 10 10 50
1 2 3 4 5
1 5 2 3 4
2 3 4 5 1
outputCopy
50
inputCopy
1 1
1000000000
1
outputCopy
1000000000
inputCopy
5 5
1000000000 1000000000 1000000000 1000000000 1000000000
5 4 3 2 1
5 4 3 2 1
5 4 3 2 1
5 4 3 2 1
5 4 3 2 1
outputCopy
5000000000
inputCopy
1 3
1 2 3
3 3 3
outputCopy
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1826E)
}
