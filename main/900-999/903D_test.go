package main

import (
	"bytes"
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"math/rand"
	"strconv"
	"testing"
)

func TestCF903D(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5
1 2 3 1 3
outputCopy
4
inputCopy
4
6 6 5 5
outputCopy
0
inputCopy
4
6 6 4 4
outputCopy
-8`
	testutil.AssertEqualCase(t, rawText, 0, CF903D)
}

// 无尽对拍
func TestCF903DInf(t *testing.T) {
	//rand.Seed(time.Now().UnixNano())
	inputGenerator := func() string {
		buf := &bytes.Buffer{}
		n := 20
		buf.WriteString(strconv.Itoa(n) + "\n")
		for i := 0; i < n; i++ {
			buf.WriteString(strconv.Itoa(rand.Intn(100)+1) + " ")
		}
		buf.WriteByte('\n')
		return buf.String()
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		var n int
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		abs := func(x int) int {
			if x < 0 {
				return -x
			}
			return x
		}
		d := func(x,y int) int {
			if abs(x-y) > 1 {
				return y-x
			}
			return 0
		}
		s := 0
		for i := range a {
			for j := i; j < n; j++ {
				s+=d(a[i],a[j])
			}
		}
		Fprint(out, s)
	}

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF, CF903D)
}
