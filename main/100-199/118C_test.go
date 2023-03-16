package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/118/C
// https://codeforces.com/problemset/status/118/problem/C
func TestCF118C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6 5
898196
outputCopy
4
888188
inputCopy
3 2
533
outputCopy
0
533
inputCopy
10 6
0001112223
outputCopy
3
0000002223`
	testutil.AssertEqualCase(t, rawText, 0, CF118C)
}
