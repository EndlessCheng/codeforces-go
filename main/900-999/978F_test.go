package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF978F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 2
10 4 10 15
1 2
4 3
outputCopy
0 0 1 2 
inputCopy
10 4
5 4 1 5 4 3 7 1 2 5
4 6
2 1
10 8
3 5
outputCopy
5 4 0 5 3 3 9 0 2 5
inputCopy
10 35
322022227 751269818 629795150 369443545 344607287 250044294 476897672 184054549 986884572 917181121
6 3
7 3
1 9
7 9
10 7
3 4
8 6
7 4
6 10
7 2
3 5
6 9
3 10
8 7
6 5
8 1
8 5
1 7
8 10
8 2
1 5
10 4
6 7
4 6
2 6
5 4
9 10
9 2
4 8
5 9
4 1
3 2
2 1
4 2
9 8
outputCopy
1 1 2 0 0 0 1 0 2 3`
	testutil.AssertEqualCase(t, rawText, -1, CF978F)
}
