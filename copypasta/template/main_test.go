package main

import (
	"bytes"
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"math/rand"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	// TODO: 测试参数的下界和上界！
	customTestCases := [][2]string{
		{
			``,
			``,
		},
	}
	if len(customTestCases) > 0 && strings.TrimSpace(customTestCases[0][0]) != "" {
		testutil.AssertEqualStringCase(t, customTestCases, 0, run)
		//testutil.AssertEqualRunResults(t, customTestCases, 0, runAC, run)
		t.Log("======= custom =======")
	}

	dir, _ := filepath.Abs(".")
	testutil.AssertEqualFileCase(t, dir, 0, run)
	t.Logf("Current problem is [%s]", filepath.Base(dir))
}

// 无尽对拍
func Test2(t *testing.T) {
	return
	//rand.Seed(time.Now().UnixNano())
	inputGenerator := func() string {
		buf := &bytes.Buffer{}
		n := 10
		buf.WriteString(strconv.Itoa(n) + "\n")
		for i := 0; i < n; i++ {
			buf.WriteString(strconv.Itoa(rand.Intn(n)+1) + " ")
		}
		buf.WriteByte('\n')
		return buf.String()
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		var s string
		Fscan(in, &s)
		ans := 0

		Fprint(out, ans)
	}

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF, run)
}
