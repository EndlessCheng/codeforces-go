package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1675/C
// https://codeforces.com/problemset/status/1675/problem/C
func TestCF1675C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8
0
1
1110000
?????
1?1??0?0
0?0???
??11
??0??
outputCopy
1
1
2
5
4
1
1
3`
	testutil.AssertEqualCase(t, rawText, 0, CF1675C)
}
