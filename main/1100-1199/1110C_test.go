package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1110/C
// https://codeforces.com/problemset/status/1110/problem/C
func TestCF1110C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2
3
5
outputCopy
3
1
7`
	testutil.AssertEqualCase(t, rawText, 0, CF1110C)
}
