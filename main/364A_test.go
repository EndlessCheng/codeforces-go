package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol364A(t *testing.T) {
	// just copy from website
	rawText := `
0
010
outputCopy
20
inputCopy
0
000
outputCopy
36
inputCopy
0
00
outputCopy
9
inputCopy
0
0
outputCopy
1
inputCopy
9
121
outputCopy
4
inputCopy
9
12
outputCopy
1
inputCopy
10
12345
outputCopy
6
inputCopy
16
439873893693495623498263984765
outputCopy
40`
	testutil.AssertEqualCase(t, rawText, 0, Sol364A)
}
