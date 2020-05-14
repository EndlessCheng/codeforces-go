package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1353F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
3 4
1 2 3 4
5 6 7 8
9 10 11 12
5 5
2 5 4 8 3
9 10 11 5 1
12 8 4 2 5
2 2 5 4 1
6 8 2 4 2
2 2
100 10
10 1
1 2
123456789876543 987654321234567
1 1
42
outputCopy
9
49
111
864197531358023
0
inputCopy
2
2 2
1 1
1 1
1 1
1000000000000000
outputCopy
3
0`
	testutil.AssertEqualCase(t, rawText, 2, CF1353F)
}
