package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1398/C
// https://codeforces.com/problemset/status/1398/problem/C
func TestCF1398C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
3
120
5
11011
6
600005
outputCopy
3
6
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1398C)
}
