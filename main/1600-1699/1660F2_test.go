package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1660/F2
// https://codeforces.com/problemset/status/1660/problem/F2
func TestCF1660F2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
3
+-+
5
-+---
4
----
7
--+---+
6
+++---
outputCopy
2
4
2
7
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1660F2)
}
