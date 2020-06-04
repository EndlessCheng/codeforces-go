package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF500D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
2 3 5
1 3 3
5
1 4
2 2
1 2
2 1
1 1
outputCopy
14.0000000000
12.0000000000
8.0000000000
6.0000000000
4.0000000000
inputCopy
6
1 5 3
5 3 2
6 1 7
1 4 4
5 2 3
5
1 2
2 1
3 5
4 1
5 2
outputCopy
19.6000000000
18.6000000000
16.6000000000
13.6000000000
12.6000000000
inputCopy
5
1 2 11
2 3 10
3 4 10
4 5 10
1
1 10
outputCopy
14.0000000000`
	testutil.AssertEqualCase(t, rawText, 0, CF500D)
}
