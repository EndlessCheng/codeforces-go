package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"testing"
)

func TestCF1067B(t *testing.T) {
	// just copy from website
	rawText := `
14 2
1 4
2 4
3 4
4 13
10 5
11 5
12 5
14 5
5 13
6 7
8 6
13 6
9 6
outputCopy
Yes
inputCopy
3 1
1 3
2 3
outputCopy
No
inputCopy
8 1
8 2
2 5
5 1
7 2
2 4
3 5
5 6
outputCopy
No`
	testutil.AssertEqualCase(t, rawText, 0, func(r io.Reader, w io.Writer) { CF1067B(r, w) })
}
