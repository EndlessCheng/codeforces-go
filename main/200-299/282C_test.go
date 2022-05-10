package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/282/C
// https://codeforces.com/problemset/status/282/problem/C
func TestCF282C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
11
10
outputCopy
YES
inputCopy
1
01
outputCopy
NO
inputCopy
000
101
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF282C)
}
