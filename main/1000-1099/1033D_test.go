package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1033D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
9
15
143
outputCopy
32
inputCopy
1
7400840699802997
outputCopy
4
inputCopy
8 
4606061759128693
4606066102679989
4606069767552943
4606063116488033
4606063930903637
4606064745319241
4606063930904021
4606065559735517
outputCopy
1920
inputCopy
3
4
8
16
outputCopy
10
inputCopy
3
15
15
21
outputCopy
24`
	testutil.AssertEqualCase(t, rawText, 0, CF1033D)
}
