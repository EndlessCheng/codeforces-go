package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1504/problem/A
// https://codeforces.com/problemset/status/1504/problem/A
func TestCF1504A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
cbabc
ab
zza
ba
a
nutforajaroftuna
outputCopy
YES
cbabac
YES
aab
YES
zaza
YES
baa
NO
YES
nutforajarofatuna`
	testutil.AssertEqualCase(t, rawText, 0, CF1504A)
}
