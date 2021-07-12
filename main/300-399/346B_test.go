package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/346/B
// https://codeforces.com/problemset/status/346/problem/B
func TestCF346B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
AJKEQSLOBSROFGZ
OVGURWZLWVLUXTH
OZ
outputCopy
ORZ
inputCopy
AA
A
A
outputCopy
0
inputCopy
ABABA
ABABA
AB
outputCopy
BBA`
	testutil.AssertEqualCase(t, rawText, -1, CF346B)
}
