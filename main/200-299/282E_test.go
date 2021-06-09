package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/282/E
// https://codeforces.com/problemset/status/282/problem/E
func TestCF282E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2
1 2
outputCopy
3
inputCopy
3
1 2 3
outputCopy
3
inputCopy
2
1000 1000
outputCopy
1000
inputCopy
6
13 21 3 61 1 2
outputCopy
62`
	testutil.AssertEqualCase(t, rawText, 0, CF282E)
}
