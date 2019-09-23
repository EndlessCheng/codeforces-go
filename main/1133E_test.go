package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1133E(t *testing.T) {
	// just copy from website
	rawText := `5 4
1 10 13 20 23
outputCopy
5
inputCopy
6 5
1 10 13 20 23 30
outputCopy
6
inputCopy
50 49
2321 2298 1227 3465 748 4678 4564 4927 3070 1180 4855 1136 3238 1941 4668 4807 1115 1400 4836 1525 4004 3071 3839 1565 3408 912 1824 2191 4670 1821 3623 3881 1015 3677 350 2937 1527 1057 4277 1132 759 3399 4175 4507 3102 1571 3626 2105 3251 257
outputCopy
50
inputCopy
5 2
1 2 15 15 15
outputCopy
5
inputCopy
6 1
36 4 1 25 9 16
outputCopy
2
inputCopy
4 4
1 10 100 1000
outputCopy
4`
	testutil.AssertEqual(t, rawText, Sol1133E)
}

func TestName(t *testing.T) {
	arr := make([][6]int, 10)
	t.Log(arr)
}