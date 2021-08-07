package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/490/E
// https://codeforces.com/problemset/status/490/problem/E
func TestCF490E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
?
18
1?
outputCopy
YES
1
18
19
inputCopy
2
??
?
outputCopy
NO
inputCopy
5
12224
12??5
12226
?0000
?00000
outputCopy
YES
12224
12225
12226
20000
100000`
	testutil.AssertEqualCase(t, rawText, 0, CF490E)
}
