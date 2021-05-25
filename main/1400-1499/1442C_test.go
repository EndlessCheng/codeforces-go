package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1442/C
// https://codeforces.com/problemset/status/1442/problem/C
func TestCF1442C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
4 4
1 2
2 3
3 4
4 1
outputCopy
2
inputCopy
4 3
2 1
2 3
4 3
outputCopy
10
inputCopy
10 20
2 1
7 9
10 2
4 9
3 1
6 4
3 6
2 9
5 2
3 9
6 8
8 7
10 4
7 4
8 5
3 4
6 7
2 6
10 6
3 8
outputCopy
3
inputCopy
50 49
1 3
6 46
47 25
11 49
47 10
26 10
12 38
45 38
24 39
34 22
36 3
21 16
43 44
45 23
2 31
26 13
28 42
43 30
12 27
32 44
24 25
28 20
15 19
6 48
41 7
15 17
8 9
2 48
33 5
33 23
4 19
40 31
11 9
40 39
35 27
14 37
32 50
41 20
21 13
14 42
18 30
35 22
36 5
18 7
4 49
29 16
29 17
8 37
34 46
outputCopy
`
	testutil.AssertEqualCase(t, rawText, -1, CF1442C)
}
