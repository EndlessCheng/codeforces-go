package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF1311E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
3
5 7
10 19
10 18
outputCopy
YES
1 2 1 3 
YES
1 2 3 3 9 9 2 1 6 
NO
inputCopy
178
11 27
11 31
12 32
11 39
12 24
12 55
11 23
5 6
11 51
12 25
11 45
11 49
10 26
10 45
5 9
9 35
11 22
10 36
11 30
9 20
9 16
10 37
8 12
5 10
7 16
11 42
3 3
9 17
12 45
8 18
11 54
9 21
10 35
8 14
11 38
6 11
10 46
9 28
12 54
11 25
12 26
12 50
12 44
12 52
5 5
10 42
12 49
12 51
8 15
9 15
12 47
11 24
12 30
9 19
4 5
12 35
9 18
12 42
12 41
12 28
11 28
11 55
7 13
10 43
11 29
10 40
8 16
10 33
10 29
10 34
12 53
7 20
8 29
9 36
11 40
11 53
outputCopy`
	testutil.AssertEqualCase(t, rawText, -1, CF1311E)
}
