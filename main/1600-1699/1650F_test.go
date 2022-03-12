package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/contest/1650/problem/F
// https://codeforces.com/problemset/status/1650/problem/F
func TestCF1650F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
3 5
5 7 8
1 1 30
2 3 50
2 3 100
1 1 80
3 3 100
1 5
51
1 36 91
1 8 40
1 42 83
1 3 45
1 13 40
2 9
9 20
2 8 64
2 7 64
1 20 56
2 8 76
2 20 48
1 2 89
1 3 38
2 18 66
1 7 51
3 2
7 18 33
1 5 80
3 4 37
2 5
569452312 703565975
1 928391659 66
1 915310 82
2 87017081 92
1 415310 54
2 567745964 82
outputCopy
4
1 4 3 5 
3
2 4 5 
4
6 7 1 2 
-1
4
2 4 3 5 
inputCopy
3
3 9
20 31 40
1 9 64
3 17 100
3 9 59
3 18 57
3 20 49
2 20 82
2 14 95
1 8 75
2 16 67
2 6
20 36
2 2 66
2 20 93
1 3 46
1 10 64
2 8 49
2 18 40
1 1
1000000000
1 1000000000 100
outputCopy
-1
4
3 4 1 5 
1
1 `
	testutil.AssertEqualCase(t, rawText, 0, CF1650F)
}
