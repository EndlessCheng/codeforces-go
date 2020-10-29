package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1401D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
4
1 2
2 3
3 4
2
2 2
4
3 4
1 3
3 2
2
3 2
7
6 1
2 3
4 6
7 3
5 1
3 6
4
7 5 13 3
outputCopy
17
18
286
inputCopy
1
16
5 10
16 1
14 1
7 5
13 2
16 11
1 7
1 4
3 14
8 16
1 6
4 9
4 12
5 13
1 15
18
45893 9901 51229 15511 46559 28433 4231 30241 29837 34421 12953 6577 12143 52711 40493 89 34819 28571
outputCopy
667739716`
	testutil.AssertEqualCase(t, rawText, 0, CF1401D)
}
