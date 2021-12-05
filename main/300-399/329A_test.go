package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/329/A
// https://codeforces.com/problemset/status/329/problem/A
func TestCF329A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
.E.
E.E
.E.
outputCopy
1 1
2 2
3 3
inputCopy
3
EEE
E..
E.E
outputCopy
-1
inputCopy
5
EE.EE
E.EE.
E...E
.EE.E
EE.EE
outputCopy
3 3
1 3
2 2
4 4
5 3`
	testutil.AssertEqualCase(t, rawText, 0, CF329A)
}
