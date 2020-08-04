package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF830A(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
2 4 50
20 100
60 10 40 80
outputCopy
50
inputCopy
1 2 10
11
15 7
outputCopy
7
inputCopy
40 45 1000
6 55 34 32 20 76 2 84 47 68 31 60 14 70 99 72 21 61 81 79 26 51 96 86 10 1 43 69 87 78 13 11 80 67 50 52 9 29 94 12
1974 1232 234 28 1456 626 408 1086 1525 1209 1096 940 795 1867 548 1774 1993 1199 1112 1087 1923 1156 876 1715 1815 1027 1658 955 398 910 620 1164 749 996 113 109 500 328 800 826 766 518 1474 1038 1029
outputCopy
2449`
	testutil.AssertEqualCase(t, rawText, 0, CF830A)
}
