package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1230C(t *testing.T) {
	// just copy from website
	rawText := `
7 12
6 3
3 5
7 5
1 5
1 7
7 6
4 1
2 1
1 6
5 6
3 4
4 2
outputCopy
12
inputCopy
7 14
2 7
5 7
3 4
4 2
2 3
4 1
6 5
4 7
6 2
6 1
5 3
5 1
7 6
3 1
outputCopy
13
inputCopy
7 14
7 3
2 4
2 1
2 5
5 3
6 7
4 7
5 4
7 5
4 3
4 1
6 1
6 3
3 1
outputCopy
13
inputCopy
4 4
1 2
2 3
3 4
4 1
outputCopy
4
inputCopy
7 0
outputCopy
0
inputCopy
3 1
1 3
outputCopy
1
inputCopy
7 21
1 2
1 3
1 4
1 5
1 6
1 7
2 3
2 4
2 5
2 6
2 7
3 4
3 5
3 6
3 7
4 5
4 6
4 7
5 6
5 7
6 7
outputCopy
16`
	testutil.AssertEqual(t, rawText, Sol1230C)
}
