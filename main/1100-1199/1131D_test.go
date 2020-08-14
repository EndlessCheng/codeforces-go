package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1131D(t *testing.T) {
	// just copy from website
	rawText := `
3 3
<>>
<<>
<<<
outputCopy
Yes
5 3 1 
6 4 2 
inputCopy
3 3
<<<
<<=
<<=
outputCopy
Yes
1 2 2 
3 3 2 
inputCopy
3 4
>>>>
>>>>
>>>>
outputCopy
Yes
2 2 2 
1 1 1 1 
inputCopy
3 3
>>>
<<<
>>>
outputCopy
Yes
3 1 3 
2 2 2 
inputCopy
3 2
==
=<
==
outputCopy
No`
	testutil.AssertEqualCase(t, rawText, -1, Sol1131D)
}
