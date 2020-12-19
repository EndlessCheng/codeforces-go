package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1463E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 2
2 3 0 5 3
1 5
5 4
outputCopy
3 2 1 5 4
inputCopy
5 2
2 3 0 5 3
1 5
5 1
outputCopy
0
inputCopy
5 1
2 3 0 5 3
4 5
outputCopy
0
inputCopy
5 4
2 3 0 5 3
2 1
3 5
5 2
1 4
outputCopy
3 5 2 1 4
inputCopy
10 2
7 9 10 3 10 8 0 5 7 7
7 10
8 9
outputCopy
7 10 1 3 5 4 8 9 2 6 
inputCopy
10 3
4 1 5 0 4 4 3 6 5 4
10 5
2 3
6 8
outputCopy
4 1 6 8 10 5 2 3 9 7 
inputCopy
3 1
0 3 1
1 2
outputCopy
0
inputCopy
2 1
0 1
2 1
outputCopy
0`
	testutil.AssertEqualCase(t, rawText, -1, CF1463E)
}
