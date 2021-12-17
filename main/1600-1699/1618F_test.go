package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1618/F
// https://codeforces.com/problemset/status/1618/problem/F
func TestCF1618F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3
outputCopy
YES
inputCopy
7 4
outputCopy
NO
inputCopy
2 8
outputCopy
NO
inputCopy
34 69
outputCopy
YES
inputCopy
8935891487501725 71487131900013807
outputCopy
YES
inputCopy
3165137368662540 34690334760256012
outputCopy
NO
inputCopy
26 47
outputCopy
YES`
	testutil.AssertEqualCase(t, rawText, -1, CF1618F)
}
