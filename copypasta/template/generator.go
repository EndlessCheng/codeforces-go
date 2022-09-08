package main

import (
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	CmdCodeforces = "cf"  // https://github.com/xalanq/cf-tool
	CmdAtcoder    = "atc" // https://github.com/sempr/cf-tool rename
)

// 生成 CF 比赛模板（需要先 cf race，以确认题目数量）
func GenCodeforcesContestTemplates(cmdName, rootPath, contestID string, overwrite bool) error {
	if contestID == "" {
		fmt.Println("contest ID is empty")
		return nil
	}

	openedOneFile := false

	return filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if path == rootPath || !info.IsDir() {
			return nil
		}

		parentName := filepath.Base(path)
		for _, srcFileName := range []string{"main.go", "main_test.go"} {
			// 为了便于区分，把 main 替换成所在目录的名字
			dstFileName := strings.Replace(srcFileName, "main", parentName, 1)
			dstFilePath := filepath.Join(path, dstFileName)
			if !overwrite {
				if _, err := os.Stat(dstFilePath); !os.IsNotExist(err) {
					continue
				}
			}
			if err := copyFile(dstFilePath, srcFileName); err != nil {
				return err
			}
			if !openedOneFile {
				openedOneFile = true
				open.Run(absPath(dstFilePath))
			}
		}
		cmd := fmt.Sprintf("%s submit contest %s %s -f %s.go", cmdName, contestID, parentName, parentName)
		if err := ioutil.WriteFile(filepath.Join(path, parentName+".bat"), []byte(cmd), 0644); err != nil {
			return err
		}
		return nil
	})
}

// 生成单道题目的模板（Codeforces）
func GenCodeforcesProblemTemplates(problemURL string, openWebsite bool) error {
	urlObj, err := url.Parse(problemURL)
	if err != nil {
		return err
	}

	contestID, problemID, isGYM := parseCodeforcesProblemURL(problemURL)
	if _, err := strconv.Atoi(contestID); err != nil {
		return fmt.Errorf("invalid URL: %v", err)
	}

	var statusURL string
	if isGYM {
		statusURL = fmt.Sprintf("https://%s/gym/%s/status/%s", urlObj.Host, contestID, problemID)
	} else {
		statusURL = fmt.Sprintf("https://%s/problemset/status/%s/problem/%s", urlObj.Host, contestID, problemID)
	}

	if openWebsite {
		luoguURL := fmt.Sprintf("https://www.luogu.com.cn/problem/CF%s%s", contestID, problemID)
		open.Run(luoguURL)
		open.Run(statusURL)
		open.Run(problemURL)
	}

	if !isGYM {
		problemID = contestID + problemID
	}
	mainStr := fmt.Sprintf(`package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func CF%[1]s(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	
}

func main() { CF%[1]s(os.Stdin, os.Stdout) }
`, problemID)
	mainTestStr := fmt.Sprintf(`package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// %s
// %s
func TestCF%[3]s(t *testing.T) {
	// just copy from website
	rawText := `+"`\n`"+`
	testutil.AssertEqualCase(t, rawText, 0, CF%[3]s)
}
`, problemURL, statusURL, problemID)

	var dir string
	if isGYM {
		dir = "../../main/gym/" + contestID + "/"
	} else {
		dir = "../../main/" + genDirName(contestID) + "/"
	}
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	mainFilePath := dir + problemID + ".go"
	if _, err := os.Stat(mainFilePath); !os.IsNotExist(err) {
		open.Run(absPath(mainFilePath))
		return fmt.Errorf("文件已存在！")
	}
	if err := ioutil.WriteFile(mainFilePath, []byte(mainStr), 0644); err != nil {
		return err
	}
	open.Run(absPath(mainFilePath))
	testFilePath := dir + problemID + "_test.go"
	if err := ioutil.WriteFile(testFilePath, []byte(mainTestStr), 0644); err != nil {
		return err
	}
	open.Run(absPath(testFilePath))
	return nil
}

// 在某一路径下批量生成模板
func GenTemplates(problemNum int, rootPath string, overwrite bool) error {
	for i := 'a'; i < 'a'+int32(problemNum); i++ {
		dir := rootPath + string(i) + "/"
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
		for j, fileName := range []string{"main.go", "main_test.go"} {
			goFilePath := dir + strings.Replace(fileName, "main", string(i), 1)
			if !overwrite {
				if _, err := os.Stat(goFilePath); !os.IsNotExist(err) {
					continue
				}
			}
			if err := copyFile(goFilePath, fileName); err != nil {
				return err
			}
			if i == 'a' && j == 0 {
				open.Run(absPath(goFilePath))
			}
		}
	}
	return nil
}
