package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/878/A
// https://codeforces.com/problemset/status/878/problem/A
func TestCF878A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
| 3
^ 2
| 1
outputCopy
2
| 3
^ 2
inputCopy
3
& 1
& 3
& 5
outputCopy
1
& 1
inputCopy
3
^ 1
^ 2
^ 3
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, 0, CF878A)
}
