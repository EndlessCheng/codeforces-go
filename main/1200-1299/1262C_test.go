package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1262C(t *testing.T) {
	// just copy from website
	rawText := `
4
8 2
()(())()
10 3
))()()()((
2 1
()
2 1
)(
outputCopy
4
3 4
1 1
5 8
2 2
3
4 10
1 4
6 7
0
1
1 2`
	testutil.AssertEqualCase(t, rawText, 0, Sol1262C)
}
