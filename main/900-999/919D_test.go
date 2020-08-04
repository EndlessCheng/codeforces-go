package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF919D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 4
abaca
1 2
1 3
3 4
4 5
outputCopy
3
inputCopy
6 6
xzyabc
1 2
3 1
2 3
5 4
4 3
6 4
outputCopy
-1
inputCopy
10 14
xzyzyzyzqx
1 2
2 4
3 5
4 5
2 6
6 8
6 5
2 10
3 9
10 9
4 6
1 10
2 8
3 7
outputCopy
4`
	testutil.AssertEqualCase(t, rawText, 0, CF919D)
}
