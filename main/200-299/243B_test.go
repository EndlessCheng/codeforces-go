package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF243B(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
9 12 2 3
1 2
2 3
1 3
1 4
2 5
4 5
4 6
6 5
6 7
7 5
8 7
9 1
outputCopy
YES
4 1
5 6 
9 3 2 
inputCopy
7 10 3 3
1 2
2 3
1 3
1 4
2 5
4 5
4 6
6 5
6 7
7 5
outputCopy
NO
inputCopy
39 50 10 2
36 30
22 34
19 34
21 30
23 7
35 11
17 30
1 30
20 37
22 28
34 18
12 30
8 33
28 24
26 36
22 30
28 23
6 24
23 32
19 31
20 28
12 8
12 5
26 28
15 17
28 19
22 26
16 30
13 35
28 14
7 14
27 7
4 38
33 25
25 38
1 18
33 26
38 29
4 3
24 8
33 28
30 28
30 3
31 32
36 37
24 32
28 34
29 13
11 37
28 11
outputCopy
YES
28 24
22 23 20 26 19 14 33 30 34 11 
32 8 `
	testutil.AssertEqualCase(t, rawText, 0, CF243B)
}
