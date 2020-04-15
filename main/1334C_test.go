package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1334C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1
3
7 15
2 14
5 3
outputCopy
6
inputCopy
10
8
7 4
9 10
4 8
3 1
1 7
9 5
1 4
2 5
8
9 8
4 7
7 7
3 6
8 9
2 5
4 9
2 8
7
2 2
2 7
2 10
8 6
7 4
2 6
3 10
4
4 10
7 5
6 10
9 8
7
8 9
8 1
3 4
5 1
3 10
9 3
5 5
5
6 1
7 3
6 5
6 1
1 8
2
7 8
7 10
4
6 8
9 3
8 1
9 7
4
8 7
5 4
10 10
8 8
6
6 5
7 10
7 1
10 4
8 7
2 7
outputCopy
10
5
3
5
11
11
7
15
10
16`
	testutil.AssertEqualCase(t, rawText, 0, CF1334C)
}
