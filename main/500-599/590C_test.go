package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/590/C
// https://codeforces.com/problemset/status/590/problem/C
func TestCF590C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 5
11..2
#..22
#.323
.#333
outputCopy
2
inputCopy
1 5
1#2#3
outputCopy
-1
inputCopy
1 3
123
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF590C)
}
