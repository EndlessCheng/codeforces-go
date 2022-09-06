package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/729/problem/D
// https://codeforces.com/problemset/status/729/problem/D
func TestCF729D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 1 2 1
00100
outputCopy
2
4 2
inputCopy
13 3 2 3
1000000010001
outputCopy
2
7 11`
	testutil.AssertEqualCase(t, rawText, 0, CF729D)
}
