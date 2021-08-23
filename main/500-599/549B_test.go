package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/549/B
// https://codeforces.com/problemset/status/549/problem/B
func TestCF549B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
101
010
001
0 1 2
outputCopy
1
1 
inputCopy
1
1
1
outputCopy
0

inputCopy
4
1111
0101
1110
0001
1 0 1 0
outputCopy
4
1 2 3 4 `
	testutil.AssertEqualCase(t, rawText, 0, CF549B)
}
