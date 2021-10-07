package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1579/C
// https://codeforces.com/problemset/status/1579/problem/C
func TestCF1579C(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
8
2 3 1
*.*
...
4 9 2
*.*.*...*
.*.*...*.
..*.*.*..
.....*...
4 4 1
*.*.
****
.**.
....
5 5 1
.....
*...*
.*.*.
..*.*
...*.
5 5 2
.....
*...*
.*.*.
..*.*
...*.
4 7 1
*.....*
.....*.
..*.*..
...*...
3 3 1
***
***
***
3 5 1
*...*
.***.
.**..
outputCopy
NO
YES
YES
YES
NO
NO
NO
NO`
	testutil.AssertEqualCase(t, rawText, 0, CF1579C)
}
