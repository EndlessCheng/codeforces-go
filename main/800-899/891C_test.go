package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/891/C
// https://codeforces.com/problemset/status/891/problem/C
func TestCF891C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5 7
1 2 2
1 3 2
2 3 1
2 4 1
3 4 1
3 5 2
4 5 2
4
2 3 4
3 3 4 5
2 1 7
2 1 2
outputCopy
YES
NO
YES
NO
inputCopy
12 29
2 1 1
3 1 2
4 1 5
5 1 3
6 3 5
7 3 3
8 7 4
9 4 2
10 2 4
11 1 4
12 2 5
5 9 2
6 9 3
12 5 1
2 10 1
10 11 5
5 7 4
2 12 5
10 11 3
5 9 3
2 12 2
11 12 1
1 6 4
2 10 4
12 7 1
6 12 4
7 10 1
4 7 3
12 4 5
10
10 12 6 3 8 7 5 11 10 9 1
9 12 9 10 7 4 6 1 8 3
2 5 3
7 7 11 10 12 6 5 9
9 3 8 10 2 9 12 6 11 7
2 6 2
8 4 1 7 10 12 2 9 11
10 5 9 3 1 4 7 6 2 10 8
9 10 8 6 11 1 4 7 3 2
3 7 6 11
outputCopy
NO
NO
NO
NO
NO
NO
NO
NO
NO
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF891C)
}
