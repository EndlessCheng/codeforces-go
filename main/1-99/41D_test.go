package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/41/D
// https://codeforces.com/problemset/status/41/problem/D
func TestCF41D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3 1
123
456
789
outputCopy
16
2
RL
inputCopy
3 3 0
123
456
789
outputCopy
17
3
LR
inputCopy
2 2 10
98
75
outputCopy
-1
inputCopy
3 4 2
8244
4768
4474
outputCopy
18
3
LR
inputCopy
4 3 10
194
707
733
633
outputCopy
22
3
LLR`
	testutil.AssertEqualCase(t, rawText, -1, CF41D)
}
