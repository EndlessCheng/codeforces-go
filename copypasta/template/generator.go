package main

import (
	"fmt"
	"github.com/levigross/grequests"
	"github.com/skratchdot/open-golang/open"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// 生成比赛模板（需要在创建目录之后）
func GenContestTemplates(contestID string, overwrite bool) error {
	if contestID == "" {
		return nil
	}

	rootPath := fmt.Sprintf("../../dash/%s/", contestID)
	openedOneFile := false

	dirNames := []string{}
	if err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if path == rootPath || !info.IsDir() {
			return nil
		}
		parentName := filepath.Base(path)
		dirNames = append(dirNames, parentName)
		for _, srcFileName := range [...]string{"main.go", "main_test.go"} {
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
		return nil
	}); err != nil {
		return err
	}

	_, err := strconv.Atoi(contestID)
	if isCF := err == nil; isCF {
		tips := fmt.Sprintf("cd %s\n", contestID)
		for _, name := range dirNames {
			tips += fmt.Sprintf("cf submit %s %[2]s %[2]s/%[2]s.go\n", contestID, name)
		}
		if err := ioutil.WriteFile(rootPath+"CF_SUBMIT.txt", []byte(tips), 0644); err != nil {
			return err
		}
	}
	return nil
}

// 生成单道题目的模板（Codeforces）
func GenCodeforcesNormalTemplates(problemURL string, openWebsite bool) error {
	contestID, problemID, isGYM := parseCodeforcesProblemURL(problemURL)
	if _, err := strconv.Atoi(contestID); err != nil {
		return fmt.Errorf("invalid URL: %v", err)
	}

	if openWebsite {
		luoguURL := fmt.Sprintf("https://www.luogu.com.cn/problem/solution/CF%s%s", contestID, problemID)
		open.Run(luoguURL)

		urlObj, err := url.Parse(problemURL)
		if err != nil {
			return err
		}
		var statusURL string
		if isGYM {
			statusURL = fmt.Sprintf("https://%s/gym/%s/status/%s", urlObj.Host, contestID, problemID)
		} else {
			statusURL = fmt.Sprintf("https://%s/problemset/status/%s/problem/%s", urlObj.Host, contestID, problemID)
		}
		open.Run(statusURL)

		if resp, err := grequests.Head(problemURL, nil); err != nil {
			fmt.Println(err)
			// CF 连接失败，打开洛谷的页面
			problemURL = fmt.Sprintf("https://www.luogu.com.cn/problem/CF%s%s", contestID, problemID)
		} else if resp.StatusCode != http.StatusOK {
			fmt.Println(resp.StatusCode, resp)
			// CF 维护中，打开洛谷的页面
			problemURL = fmt.Sprintf("https://www.luogu.com.cn/problem/CF%s%s", contestID, problemID)
		}
		open.Run(problemURL)
	}

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

func main() { CF%[1]s(os.Stdin, os.Stdout) }
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

	const rootPath = "../../main/"
	mainFilePath := rootPath + problemID + ".go"
	if _, err := os.Stat(mainFilePath); !os.IsNotExist(err) {
		open.Run(absPath(mainFilePath))
		return fmt.Errorf("文件已存在！")
	}
	if err := ioutil.WriteFile(mainFilePath, []byte(mainStr), 0644); err != nil {
		return err
	}
	open.Run(absPath(mainFilePath))
	testFilePath := rootPath + problemID + "_test.go"
	if err := ioutil.WriteFile(testFilePath, []byte(mainTestStr), 0644); err != nil {
		return err
	}
	open.Run(absPath(testFilePath))
	return nil
}

// 批量生成模板（非 Codeforces）
// rootPath like "../../nowcoder/2720/"
func GenNormalTemplates(rootPath string, overwrite bool) error {
	for i := 'a'; i <= 'h'; i++ {
		dir := rootPath + string(i) + "/"
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
		for _, fileName := range [...]string{"main.go", "main_test.go"} {
			goFilePath := dir + fileName
			if !overwrite {
				if _, err := os.Stat(goFilePath); !os.IsNotExist(err) {
					continue
				}
			}
			if err := copyFile(goFilePath, fileName); err != nil {
				return err
			}
		}
	}
	return nil
}
