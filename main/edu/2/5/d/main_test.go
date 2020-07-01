package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"math/rand"
	"testing"
	"time"
)

func Test(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
ababb
outputCopy
20
inputCopy
xxju
outputCopy
11
inputCopy
hhhuhhh
outputCopy
50
inputCopy
m
outputCopy
1
inputCopy
aa
outputCopy
4
inputCopy
aaa
outputCopy
10
inputCopy
aaab
outputCopy
14
inputCopy
aaaa
outputCopy
20
inputCopy
baab
outputCopy
12
inputCopy
aaabad
outputCopy
28

`
	testutil.AssertEqualCase(t, rawText, -1, run)
}

func Test2(t *testing.T) {
	testCases := [][2]string{
		{
			`aaabad`,
			``,
		},
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		s := []byte{}
		for j := 0; j < 6; j++ {
			x := byte(rand.Intn(4))
			s = append(s, 'a'+x)
		}
		testCases = append(testCases, [2]string{string(s)})
	}

	runAC := func(in io.Reader, out io.Writer) {
		var s string
		Fscan(in, &s)
		n := len(s)
		ans := 0
		for i := range s {
			for j := i + 1; j <= n; j++ {
				t := s[i:j]
				for k := range t {
					if t[:k] == t[len(t)-k:] {
						ans++
					}
				}
			}
		}
		Fprint(out, ans)
	}

	testutil.AssertEqualRunResults(t, testCases, 1, runAC, run)
}
