package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"github.com/stretchr/testify/assert"
	"io"
	"path/filepath"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	// TODO: 测试参数的下界和上界！
	customTestCases := [][2]string{
		{
			`1 2 3`,
			`2`,
		},
		{
			`1 5 3`,
			`2`,
		},
		{
			`0 0 0`,
			`0`,
		},
		{
			`1000000000000 1000000000000 1000000000000`,
			`1732050`,
		},
	}
	if len(customTestCases) > 0 && strings.TrimSpace(customTestCases[0][0]) != "" {
		testutil.AssertEqualStringCase(t, customTestCases, 0, run)
		//testutil.AssertEqualRunResults(t, customTestCases, 0, runAC, run)
		t.Log("======= custom =======")
	}

	dir, _ := filepath.Abs(".")
	testutil.AssertEqualFileCaseWithName(t, dir, "in*.txt", "ans*.txt", 0, run)
	//testutil.AssertEqualFileCaseWithName(t, dir, "*.in", "*.out", 0, run)
	t.Logf("Current problem is [%s]", filepath.Base(dir))
}

// 无尽对拍
func TestCompare(t *testing.T) {
	return
	//rand.Seed(time.Now().UnixNano())
	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		n := rg.Int(1, 10)
		rg.NewLine()
		rg.IntSlice(n, 1, n)
		//Println(rg.String())
		return rg.String()
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		//return
		var n int
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}

		ans := 0

		Fprint(out, ans)
	}

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF, run)
}

// 无尽检查输出是否正确
func TestCheck(t *testing.T) {
	return
	assert := assert.New(t)

	//rand.Seed(time.Now().UnixNano())
	inputGenerator := func() (string, testutil.OutputChecker) {
		//return ``
		rg := testutil.NewRandGenerator()
		n := rg.Int(1, 10)
		rg.NewLine()
		a := rg.IntSlice(n, 1, n)
		//Println(rg.String())
		return rg.String(), func(output string) (_b bool) {
			in := strings.NewReader(output)
			var outN int
			Fscan(in, &outN)
			if !assert.Equal(n, outN) {
				return
			}

			outA := make([]int, outN)
			for i := range outA {
				Fscan(in, &outA[i])
			}
			if !assert.EqualValues(a, outA) {
				return
			}

			return true
		}
	}

	testutil.CheckRunResultsInf(t, inputGenerator, run)
}
