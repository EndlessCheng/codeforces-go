package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF962F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3 3
1 2
2 3
3 1
outputCopy
3
1 2 3 
inputCopy
6 7
2 3
3 4
4 2
1 2
1 5
5 6
6 1
outputCopy
6
1 2 3 5 6 7 
inputCopy
5 6
1 2
2 3
2 4
4 3
2 5
5 3
outputCopy
0

inputCopy
1 0
outputCopy
0

inputCopy
7 9
7 3
7 4
1 2
2 3
3 1
3 4
4 5
5 6
6 4
outputCopy
9
1 2 3 4 5 6 7 8 9 
inputCopy
7 8
1 2
2 3
3 4
4 1
3 5
5 6
6 7
7 3
outputCopy
8
1 2 3 4 5 6 7 8 `
	testutil.AssertEqualCase(t, rawText, 6, CF962F)
}
