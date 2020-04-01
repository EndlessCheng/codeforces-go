package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1296F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
1 2
3 2
3 4
2
1 2 5
1 3 3
outputCopy
5 3 5
inputCopy
6
1 2
1 6
3 1
1 5
4 1
4
6 1 3
3 4 1
6 5 2
1 2 5
outputCopy
5 3 1 2 1 
inputCopy
6
1 2
1 6
3 1
1 5
4 1
4
6 1 1
3 4 3
6 5 3
1 2 4
outputCopy
-1
inputCopy
7
1 3
5 1
3 7
1 2
6 3
2 4
3
7 6 1
7 2 4
1 7 1
outputCopy
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF1296F)
}
