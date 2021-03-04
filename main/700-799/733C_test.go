package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/733/C
// https://codeforces.com/problemset/status/733/problem/C
func TestCF733C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
6
1 2 2 2 1 2
2
5 5
outputCopy
YES
2 L
1 R
4 L
3 L
inputCopy
5
1 2 3 4 5
1
15
outputCopy
YES
5 L
4 L
3 L
2 L
inputCopy
5
1 1 1 3 3
3
2 1 6
outputCopy
NO
inputCopy
5
1 2 3 4 5
1
10
outputCopy
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF733C)
}
