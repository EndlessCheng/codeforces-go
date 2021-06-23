package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/52/problem/B
// https://codeforces.com/problemset/status/52/problem/B
func TestCF52B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 2
**
*.
outputCopy
1
inputCopy
3 4
*..*
.**.
*.**
outputCopy
9`
	testutil.AssertEqualCase(t, rawText, 0, CF52B)
}
