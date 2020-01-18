package main

import (
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

// 生成比赛模板（需要先创建目录）
func TestGenContestTemplates(t *testing.T) {
	const contestID = "1284"
	const overwrite = false
	rootPath := fmt.Sprintf("../../dash/%s/", contestID)
	if err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if path == rootPath || !info.IsDir() {
			return nil
		}
		for _, fileName := range [...]string{"main.go", "main_test.go"} {
			goFilePath := filepath.Join(path, fileName)
			if !overwrite {
				if _, err := os.Stat(goFilePath); !os.IsNotExist(err) {
					continue
				}
			}
			if err := copyFile(goFilePath, fileName); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		t.Fatal(err)
	}

	_, err := strconv.Atoi(contestID)
	if isCF := err == nil; isCF {
		tips := fmt.Sprintf("cd %[1]s\ncf submit %[1]s a a/main.go\n", contestID)
		if err := ioutil.WriteFile(rootPath+"tips.txt", []byte(tips), 0644); err != nil {
			t.Fatal(err)
		}
	}
}

// 生成单道题目的模板（Codeforces）
func TestGenCodeforcesNormalTemplates(t *testing.T) {
	const problemURL = "https://codeforces.com/problemset/problem/690/D2"
	// https://codeforces.com/problemset/status/690/problem/D2
	// https://codeforces.com/gym/102253/problem/C
	contestID, problemID := parseProblemURL(problemURL)
	statusURL := fmt.Sprintf("https://codeforces.com/problemset/status/%s/problem/%s", contestID, problemID)
	defer open.Run(problemURL)
	defer open.Run(statusURL)

	problemID = contestID + problemID
	mainStr := fmt.Sprintf(`package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func CF%[1]s(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	
}

func main() {
	CF%[1]s(os.Stdin, os.Stdout)
}
`, problemID)
	mainTestStr := fmt.Sprintf(`package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF%[1]s(t *testing.T) {
	// just copy from website
	rawText := `+"`\n`"+`
	testutil.AssertEqualCase(t, rawText, 0, CF%[1]s)
}
`, problemID)
	rootPath := "../../main/"
	if err := ioutil.WriteFile(rootPath+problemID+".go", []byte(mainStr), 0644); err != nil {
		t.Fatal(err)
	}
	if err := ioutil.WriteFile(rootPath+problemID+"_test.go", []byte(mainTestStr), 0644); err != nil {
		t.Fatal(err)
	}
}

// 生成单道题目的模板（非 Codeforces）
func TestGenNormalTemplates(t *testing.T) {
	const rootPath = "../../nowcoder/2720/"
	const overwrite = false
	for i := 'a'; i <= 'h'; i++ {
		dir := rootPath + string(i) + "/"
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			t.Fatal(err)
		}
		for _, fileName := range [...]string{"main.go", "main_test.go"} {
			goFilePath := dir + fileName
			if !overwrite {
				if _, err := os.Stat(goFilePath); !os.IsNotExist(err) {
					continue
				}
			}
			if err := copyFile(goFilePath, fileName); err != nil {
				t.Fatal(err)
			}
		}
	}
}
