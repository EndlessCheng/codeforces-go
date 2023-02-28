package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/388/B
// https://codeforces.com/problemset/status/388/problem/B
func TestCF388B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
outputCopy
4
NNYY
NNYY
YYNN
YYNN
inputCopy
9
outputCopy
8
NNYYYNNN
NNNNNYYY
YNNNNYYY
YNNNNYYY
YNNNNYYY
NYYYYNNN
NYYYYNNN
NYYYYNNN
inputCopy
1
outputCopy
2
NY
YN
inputCopy
7
outputCopy
`
	testutil.AssertEqualCase(t, rawText, 0, CF388B)
}
