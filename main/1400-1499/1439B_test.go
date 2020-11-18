package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1439B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
5 9 4
1 2
1 3
1 4
1 5
2 3
2 4
2 5
3 4
3 5
10 15 3
1 2
2 3
3 4
4 5
5 1
1 7
2 8
3 9
4 10
5 6
7 10
10 8
8 6
6 9
9 7
4 5 4
1 2
2 3
3 4
4 1
1 3
outputCopy
2
4 1 2 3 
1 10
1 2 3 4 5 6 7 8 9 10 
-1`
	testutil.AssertEqualCase(t, rawText, 0, CF1439B)
}
