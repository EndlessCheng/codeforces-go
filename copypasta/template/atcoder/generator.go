package atcoder

import (
	"fmt"
	"github.com/levigross/grequests"
	"github.com/skratchdot/open-golang/open"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func login(username, password string) (session *grequests.Session, err error) {
	const ua = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.106 Safari/537.36"
	session = grequests.NewSession(&grequests.RequestOptions{
		UserAgent:    ua,
		UseCookieJar: true,
	})

	// "touch" home page to get CSRF token
	var resp *grequests.Response
	home := "https://atcoder.jp"
	for {
		resp, err = session.Get(home, nil)
		if err != nil {
			// sometimes we got timeout error
			fmt.Println(err)
		} else {
			break
		}
	}
	if !resp.Ok {
		err = fmt.Errorf("GET %s return code %d", home, resp.StatusCode)
		return
	}
	var csrfToken string
	root, err := html.Parse(resp)
	if err != nil {
		return
	}
	var f func(*html.Node)
	f = func(o *html.Node) {
		if o.Type == html.TextNode {
			const s = `var csrfToken = "`
			if idx := strings.Index(o.Data, s); idx != -1 {
				csrfToken = strings.TrimSpace(o.Data[idx+len(s):])
				csrfToken = csrfToken[:len(csrfToken)-1] // remove last "
				return
			}
		}
		for c := o.FirstChild; c != nil; c = c.NextSibling {
			f(c)
			if csrfToken != "" {
				return
			}
		}
	}
	f(root)

	loginURL := "https://atcoder.jp/login"
	resp, err = session.Post(loginURL, &grequests.RequestOptions{
		Data: map[string]string{
			"username":   username,
			"password":   password,
			"csrf_token": csrfToken,
		},
		//Headers: map[string]string{
		//	"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		//	"accept-encoding":           "gzip, deflate, br",
		//	"accept-language":           "zh-CN,zh;q=0.9,en;q=0.8",
		//	"cache-control":             "max-age=0",
		//	"content-type":              "application/x-www-form-urlencoded",
		//	"origin":                    "https://atcoder.jp",
		//	"referer":                   "https://atcoder.jp/login",
		//	"sec-fetch-dest":            "document",
		//	"sec-fetch-mode":            "navigate",
		//	"sec-fetch-site":            "same-origin",
		//	"sec-fetch-user":            "?1",
		//	"upgrade-insecure-requests": "1",
		//},
	})
	if err != nil {
		return
	}
	if !resp.Ok {
		return nil, fmt.Errorf("GET %s return code %d", loginURL, resp.StatusCode)
	} // resp.StatusCode != http.StatusFound
	return
}

func parseTask(session *grequests.Session, problemURL string) (sampleIns, sampleOuts []string, err error) {
	resp, err := session.Get(problemURL, nil)
	if err != nil {
		return
	}
	if !resp.Ok {
		err = fmt.Errorf("GET %s return code %d", problemURL, resp.StatusCode)
		return
	}

	root, err := html.Parse(resp)
	if err != nil {
		return
	}

	// 解析样例输入输出
	const (
		tokenInputJP  = "入力例"
		tokenOutputJP = "出力例"

		tokenInputEN  = "Sample Input "
		tokenOutputEN = "Sample Output "
	)

	var f func(*html.Node)
	f = func(o *html.Node) {
		if o.Type == html.TextNode {
			if strings.Contains(o.Data, tokenInputEN) {
				raw := o.Parent.NextSibling.FirstChild.Data
				raw = strings.TrimSpace(raw)
				sampleIns = append(sampleIns, raw)
			} else if strings.Contains(o.Data, tokenOutputEN) {
				if o.Parent.NextSibling.FirstChild == nil {
					// 样例输出为空，例如 https://atcoder.jp/contests/abc150/tasks/abc150_f
					sampleOuts = append(sampleOuts, "")
				} else {
					raw := o.Parent.NextSibling.FirstChild.Data
					raw = strings.TrimSpace(raw)
					sampleOuts = append(sampleOuts, raw)
				}
			}
			return
		}
		for c := o.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(root)

	if len(sampleIns) != len(sampleOuts) {
		err = fmt.Errorf("len(sampleIns) != len(sampleOuts) : %d != %d", len(sampleIns), len(sampleOuts))
		return
	}

	return
}

func genTemplates(session *grequests.Session, problemURL string, isContest bool) error {
	// 解析样例输入输出
	ins, outs, err := parseTask(session, problemURL)
	if err != nil {
		return err
	}

	problemName := filepath.Base(problemURL)
	sp := strings.Split(problemName, "_")
	if len(sp) != 2 {
		return fmt.Errorf("invlaid url %s", problemURL)
	}

	contestID, taskID := sp[0], sp[1]

	// 生成目录
	dirPath := filepath.Join(contestDir, contestID, taskID) + "/"
	if err := os.MkdirAll(dirPath, 0644); err != nil {
		return err
	}

	if !isContest {
		languageID := 4026
		if contestID <= "abc161" { // todo arc agc
			languageID = 3013
		}
		statusURL := filepath.Dir(filepath.Dir(problemURL)) + fmt.Sprintf("/submissions?f.Language=%d&f.Status=AC&f.Task=%s&orderBy=source_length", languageID, problemName)
		open.Start(statusURL)
	}

	// 创建 x.go
	mainFileContent := `package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)

}

func main() { run(os.Stdin, os.Stdout) }
`
	mainFilePath := dirPath + taskID + ".go"
	if !isContest || taskID == "a" {
		// 比赛时，在 IDE 中打开 A 题
		open.Start(absPath(mainFilePath))
	}
	if err := ioutil.WriteFile(mainFilePath, []byte(mainFileContent), 0644); err != nil {
		return err
	}

	// 创建 x_test.go
	examples := ""
	for i, in := range ins {
		out := outs[i]
		examples += "\n\t\t{\n"
		examples += "\t\t\t`" + in + "`,\n"
		examples += "\t\t\t`" + out + "`,\n"
		examples += "\t\t},"
	}
	submitURL := fmt.Sprintf("https://atcoder.jp/contests/%s/submit?taskScreenName=%s", contestID, problemName)
	testFileContent := fmt.Sprintf(`// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [%s]")
	testCases := [][2]string{%s
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
	//testutil.AssertEqualRunResults(t, testCases, 0, runAC, run)
}
// %s
// %s
`, taskID, examples, problemURL, submitURL)
	testFilePath := dirPath + taskID + "_test.go"
	if err := ioutil.WriteFile(testFilePath, []byte(testFileContent), 0644); err != nil {
		return err
	}

	return nil
}

func genAtCoderContestTemplates(contestID string, retryTimes int) error {
	if retryTimes == 0 {
		submitURL := fmt.Sprintf("https://atcoder.jp/contests/%s/submit", contestID)
		tasksPrintURL := fmt.Sprintf("https://atcoder.jp/contests/%s/tasks_print", contestID)
		open.Start(submitURL)
		open.Start(tasksPrintURL)
	}

	const ua = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.106 Safari/537.36"
	session := grequests.NewSession(&grequests.RequestOptions{
		UserAgent:    ua,
		UseCookieJar: true,
	})

	tasksHome := fmt.Sprintf("https://atcoder.jp/contests/%s/tasks", contestID)
	revelSession, err := ioutil.ReadFile("revel_session.txt")
	if err != nil {
		return err
	}
	resp, err := session.Get(tasksHome, &grequests.RequestOptions{
		Cookies: []*http.Cookie{
			{
				Name:  "REVEL_SESSION",
				Value: string(revelSession),
			},
		},
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return genAtCoderContestTemplates(contestID, retryTimes+1)
	}
	if !resp.Ok {
		return fmt.Errorf("未找到比赛或比赛尚未开始")
	}

	fmt.Println("开始解析样例输入输出")
	wg := &sync.WaitGroup{}
	defer wg.Wait()
	for taskID := byte('a'); taskID <= 'f'; taskID++ { // 默认六道题目
		wg.Add(1)
		// we don't want spent too much time on waiting responses one by one, so we use goroutine!
		go func(id byte) {
			defer wg.Done()
			problemURL := fmt.Sprintf("https://atcoder.jp/contests/%[1]s/tasks/%[1]s_%[2]c", contestID, id)
			if err := genTemplates(session, problemURL, true); err != nil {
				fmt.Fprintln(os.Stderr, string(id), err)
				return
			}
			fmt.Println("[ok]", string(id), problemURL)
		}(taskID)
	}

	return nil
}

func GenAtCoderContestTemplates(contestID string) error {
	return genAtCoderContestTemplates(contestID, 0)
}

func GenAtCoderProblemTemplate(problemURL string) error {
	const ua = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.106 Safari/537.36"
	session := grequests.NewSession(&grequests.RequestOptions{
		UserAgent:    ua,
		UseCookieJar: true,
	})
	return genTemplates(session, problemURL, false)
}
