package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1537/E2
// https://codeforces.com/problemset/status/1537/problem/E2
func TestCF1537E2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8 16
dbcadabc
outputCopy
dbcadabcdbcadabc
inputCopy
4 5
abcd
outputCopy
aaaaa`
	testutil.AssertEqualCase(t, rawText, 0, CF1537E2)
}
