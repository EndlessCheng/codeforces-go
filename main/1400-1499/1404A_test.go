package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1404A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
9
6 4
100110
3 2
1?1
3 2
1?0
4 4
????
7 4
1?0??1?
10 10
11??11??11
4 2
1??1
4 4
?0?0
6 2
????00
outputCopy
YES
YES
NO
YES
YES
NO
NO
YES
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1404A)
}
