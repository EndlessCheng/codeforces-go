package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1215/C
// https://codeforces.com/problemset/status/1215/problem/C
func TestCF1215C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
abab
aabb
outputCopy
2
3 3
3 2
inputCopy
1
a
b
outputCopy
-1
inputCopy
8
babbaabb
abababaa
outputCopy
3
2 6
1 3
7 8`
	testutil.AssertEqualCase(t, rawText, 0, CF1215C)
}
