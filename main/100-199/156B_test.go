package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/156/B
// https://codeforces.com/problemset/status/156/problem/B
func TestCF156B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1 1
+1
outputCopy
Truth
inputCopy
3 2
-1
-2
-3
outputCopy
Not defined
Not defined
Not defined
inputCopy
4 1
+2
-3
+4
-1
outputCopy
Lie
Not defined
Lie
Not defined`
	testutil.AssertEqualCase(t, rawText, 0, CF156B)
}
