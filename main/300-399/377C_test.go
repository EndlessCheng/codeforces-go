package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/377/C
// https://codeforces.com/problemset/status/377/problem/C
func TestCF377C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
2 1
2
p 1
p 2
outputCopy
1
inputCopy
6
6 4 5 4 5 5
4
b 2
p 1
b 1
p 2
outputCopy
0
inputCopy
4
1 2 3 4
4
p 2
b 2
p 1
b 1
outputCopy
-2`
	testutil.AssertEqualCase(t, rawText, 0, CF377C)
}
