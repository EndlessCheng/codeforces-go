package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1515/E
// https://codeforces.com/problemset/status/1515/problem/E
func TestCF1515E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 100000007
outputCopy
1
inputCopy
2 100000007
outputCopy
2
inputCopy
3 100000007
outputCopy
6
inputCopy
4 100000007
outputCopy
20
inputCopy
5 100000007
outputCopy
78
inputCopy
400 234567899
outputCopy
20914007`
	testutil.AssertEqualCase(t, rawText, 0, CF1515E)
}
