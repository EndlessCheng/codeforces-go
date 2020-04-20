package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1204C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
0110
0010
0001
1000
4
1 2 3 4
outputCopy
3
1 2 4 
inputCopy
4
0110
0010
1001
1000
20
1 2 3 4 1 2 3 4 1 2 3 4 1 2 3 4 1 2 3 4
outputCopy
11
1 2 4 2 4 2 4 2 4 2 4 
inputCopy
3
011
101
110
7
1 2 3 1 3 2 1
outputCopy
7
1 2 3 1 3 2 1 
inputCopy
4
0110
0001
0001
1000
3
1 2 4
outputCopy
2
1 4 `
	testutil.AssertEqualCase(t, rawText, 0, CF1204C)
}
