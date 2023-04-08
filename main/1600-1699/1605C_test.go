package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1605/C
// https://codeforces.com/problemset/status/1605/problem/C
func TestCF1605C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2
aa
5
cbabb
8
cacabccc
outputCopy
2
-1
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1605C)
}
