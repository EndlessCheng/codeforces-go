
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/276/D
// https://codeforces.com/problemset/status/276/problem/D
func TestCF276D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 2
outputCopy
3
inputCopy
8 16
outputCopy
31
inputCopy
1 1
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF276D)
}
