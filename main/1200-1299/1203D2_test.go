package main

import (
	"bufio"
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"testing"
)

// https://codeforces.com/problemset/problem/1203/D2
// https://codeforces.com/problemset/status/1203/problem/D2
func TestCF1203D2(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
bbaba
bb
outputCopy
3
inputCopy
baaba
ab
outputCopy
2
inputCopy
abcde
abcde
outputCopy
0
inputCopy
asdfasdf
fasd
outputCopy
3
inputCopy
zywmerhahxlqsjekpqsdqxnjiduyjrytswiweohctztgpiorwimhjmdfofqynyggcrtzslbyvkuvqrsgwyacyvcuathplliwwshusluiqwuhutnzwvuchfedhwwfdzizltdxibtsaocpnqezstblgkfdcvfsjjyzwalkksumsaljqljmmkcyejwwdkolmcgmodoiclte
zywmehahxlqjekqsdqjidytswwztgiowimhmffqygctzslbykurwacyvcuaplwshsluiqwuhutnwchfewwfdizttcpnqestgkfvsjylkksumaljmmkcjwwdkolmcgodcle
outputCopy
5
inputCopy
abaab
aaab
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, -1, CF1203D2)
}

func TestCompareCF1203D2(t *testing.T) {
	//return
	testutil.DebugTLE = 0

	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		s := rg.Str(2, 5, 'a', 'b')
		rg.NewLine()
		for i := 0; i < len(s); {
			p := rg.IntOnly(i, len(s)-1)
			rg.Byte(s[p])
			i = p + 1
		}
		return rg.String()
	}


	testutil.AssertEqualRunResultsInf(t, inputGenerator, CF1203D2AC, CF1203D2)
}

func CF1203D2AC(in io.Reader, out io.Writer) {
	var s, t string
	Fscan(bufio.NewReader(in), &s, &t)
	n, m := len(s), len(t)
	suf := make([]int, n+1)
	for i, j := n-1, m-1; i > 0; i-- {
		if j >= 0 && s[i] == t[j] {
			j--
		}
		suf[i] = m - 1 - j
	}
	ans := 0
	for l, r, j := 0, 1, 0; r <= n; r++ {
		for j+suf[r] < m {
			if s[l] == t[j] {
				j++
			}
			l++
		}
		if r-l > ans {
			ans = r - l
		}
	}
	Fprint(out, ans)
}