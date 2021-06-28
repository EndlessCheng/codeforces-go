package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/272/E
// https://codeforces.com/problemset/status/272/problem/E
func TestCF272E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3
1 2
3 2
3 1
outputCopy
100
inputCopy
2 1
2 1
outputCopy
00
inputCopy
10 6
1 2
1 3
1 4
2 3
2 4
3 4
outputCopy
0110000000`
	testutil.AssertEqualCase(t, rawText, 0, CF272E)
}
