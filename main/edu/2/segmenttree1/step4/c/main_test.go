package main

import (
	"bytes"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"math/rand"
	"strconv"
	"testing"
)

func Test(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7 6
1 2 3 6 5 4 19
1 1 3
1 2 5
1 2 4
2 2 8
1 1 6
1 1 3
outputCopy
0
1
0
7
1`
	testutil.AssertEqualCase(t, rawText, 0, run)
}

func Test2(t *testing.T) {
	//rand.Seed(time.Now().UnixNano())
	inputGenerator := func() string {
		buf := &bytes.Buffer{}
		var n, q int
		n = 1e5
		q = 1e5
		const mx int = 2
		buf.WriteString(strconv.Itoa(n) + " " + strconv.Itoa(q) + "\n")
		for i := 0; i < n; i++ {
			buf.WriteString(strconv.Itoa(rand.Intn(mx)+1) + " ")
		}
		buf.WriteByte('\n')
		for i := 0; i < q; i++ {
			op := rand.Intn(2) + 1
			buf.WriteString(strconv.Itoa(op) + " ")
			if op == 2 {
				i, v := rand.Intn(n), rand.Intn(mx)+1
				buf.WriteString(strconv.Itoa(i+1) + " " + strconv.Itoa(v))
			} else {
				l := rand.Intn(n)
				r := rand.Intn(n-l) + l
				buf.WriteString(strconv.Itoa(l+1) + " " + strconv.Itoa(r+1))
			}
			buf.WriteByte('\n')
		}
		buf.WriteByte('\n')
		//fmt.Println(buf.String())
		return buf.String()
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		return
	}

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF, run)
}
