package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"math/rand"
	"testing"
)

func Test(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
aabaaabaaabaaabaab
outputCopy
4
inputCopy
aabaabb
outputCopy
2
inputCopy
aaaa
outputCopy
4
inputCopy
ppppplppp
outputCopy
5
inputCopy
nn
outputCopy
2
inputCopy
n
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, 0, run)
}

func Test2(t *testing.T) {
	inputs := []string{}
	//rand.Seed(time.Now().UnixNano())
	for i := 0; i < 0; i++ {
		s := []byte{}
		for j := 0; j < 1000; j++ {
			x := byte(rand.Intn(2))
			s = append(s, 'a'+x)
		}
		inputs = append(inputs, string(s))
	}

	runAC := func(in io.Reader, out io.Writer) {
		var s string
		Fscan(in, &s)
		n := len(s)
		ans := 1
		for i := range s {
			for j := i + 1; j <= n; j++ {
				t := s[i:j]
				cnt := 1
				for k := j; k+len(t) <= n; k += len(t) {
					if s[k:k+len(t)] != t {
						break
					}
					cnt++
				}
				if cnt > ans {
					ans = cnt
				}
			}
		}
		Fprint(out, ans)
	}

	gen := func() string {
		s := []byte{}
		mx := rand.Intn(50)+3
		for j := 0; j < mx; j++ {
			x := byte(rand.Intn(5))
			s = append(s, 'a'+x)
		}
		return string(s)
	}

	testutil.AssertEqualRunResultsInf(t, gen, runAC, run)
}
