package main

import (
	"fmt"
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	testutil2 "github.com/EndlessCheng/codeforces-go/main/testutil"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

//func init() { rand.Seed(time.Now().UnixNano()) }

func TestCompareInf(t *testing.T) {
	inputGenerator := func() (a []int) {
		rg := testutil2.NewRandGenerator()
		n := rg.Int(1, 5)
		a = rg.IntSlice(n, 1, 5)
		return
	}

	runAC := func(a []int) (ans int) {
		// 若要修改 a，必须先 copy 一份，在 copied 上修改

		return
	}

	// test examples first (or make it global)
	examples := [][]string{

	}
	if err := testutil.RunLeetCodeFuncWithExamples(t, runAC, examples, 0); err != nil {
		t.Fatal(err)
	}

	testutil.CompareInf(t, inputGenerator, runAC, nil /*TODO*/)
}

func TestCheckInf(t *testing.T) {
	var solve func([]int) []int /*TODO*/
	for tc := 1; ; tc++ {
		if tc%1e5 == 0 {
			fmt.Println(tc)
		}

		rg := testutil2.NewRandGenerator()
		n := rg.Int(1, 9)
		a := rg.IntSlice(n, 1, 9)
		myAns := solve(a)
		// check myAns is valid ...
		_ = myAns
	}
}

func Test_transCode(t *testing.T) {
	code := `   

`
	fmt.Println(transCode(code))
}

func Test_checkTodo(t *testing.T) {
	dir := "."
	if err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		code := string(data)
		if !strings.Contains(code, "return") {
			fmt.Println("TODO:", path)
			return nil
		}
		line := strings.Split(code, "\n")
		for i, l := range line {
			l = strings.TrimSpace(l)
			if strings.HasPrefix(l, "func ") && l[len(l)-1] == '{' {
				nextLine := strings.TrimSpace(line[i+1])
				if nextLine == "" || nextLine == "return" {
					fmt.Println("TODO:", path)
				}
				break
			}
		}
		return nil
	}); err != nil {
		t.Fatal(err)
	}
}
