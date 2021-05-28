package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/763/A
// https://codeforces.com/problemset/status/763/problem/A
func TestCF763A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 2
2 3
3 4
1 2 1 1
outputCopy
YES
2
inputCopy
3
1 2
2 3
1 2 3
outputCopy
YES
2
inputCopy
4
1 2
2 3
3 4
1 2 1 2
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, -1, CF763A)
}
