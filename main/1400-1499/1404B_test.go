package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1404B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4
4 3 2 1 2
1 2
1 3
1 4
6 6 1 2 5
1 2
6 5
2 3
3 4
4 5
9 3 9 2 5
1 2
1 6
1 9
1 3
9 5
7 9
4 8
4 3
11 8 11 3 3
1 2
11 9
4 9
6 5
2 10
3 2
5 9
8 3
7 4
7 10
outputCopy
Alice
Bob
Alice
Alice`
	testutil.AssertEqualCase(t, rawText, 0, CF1404B)
}
