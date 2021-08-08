package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// http://codeforces.com/problemset/problem/1063/A
// https://codeforces.com/problemset/status/1063/problem/A
func TestCF1063A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
oolol
outputCopy
ololo
inputCopy
16
gagadbcgghhchbdf
outputCopy
abccbaghghghgdfd`
	testutil.AssertEqualCase(t, rawText, 0, CF1063A)
}
