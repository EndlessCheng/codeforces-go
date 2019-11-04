package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol706C(t *testing.T) {
	// just copy from website
	rawText := `
2
1 2
ba
ac
outputCopy
1
inputCopy
3
1 3 1
aa
ba
ac
outputCopy
1
inputCopy
2
5 5
bbb
aaa
outputCopy
-1
inputCopy
2
3 3
aaa
aa
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, Sol706C)
}
