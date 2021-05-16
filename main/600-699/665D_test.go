package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/665/D
// https://codeforces.com/problemset/status/665/problem/D
func TestCF665D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
2 3
outputCopy
2
3 2
inputCopy
2
2 2
outputCopy
1
2
inputCopy
3
2 1 1
outputCopy
3
1 1 2
inputCopy
2
83 14
outputCopy
2
14 83
inputCopy
10
10 10 1 2 3 3 1 2 1 5
outputCopy
4
1 1 10 1`
	testutil.AssertEqualCase(t, rawText, 0, CF665D)
}
