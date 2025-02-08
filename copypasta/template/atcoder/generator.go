package atcoder

import (
	"fmt"
	"github.com/levigross/grequests"
	"github.com/skratchdot/open-golang/open"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"
)

const ua = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.106 Safari/537.36"

func fetchStartTime(contestID string) (t time.Time, err error) {
	contestHome := "https://atcoder.jp/contests/" + contestID
	resp, err := grequests.Get(contestHome, &grequests.RequestOptions{UserAgent: ua})
	if err != nil {
		return
	}
	if !resp.Ok {
		return time.Time{}, fmt.Errorf("%d", resp.StatusCode)
	}

	htmlStr := resp.String()
	const token = `fixtime-full'>`
	i := strings.Index(htmlStr, token)
	if i == -1 {
		return time.Time{}, fmt.Errorf("invalid html content %s", htmlStr)
	}

	datetimeStr := htmlStr[i+len(token) : i+len(token)+24] // 2020-08-02 21:00:00+0900
	return time.Parse("2006-01-02 15:04:05-0700", datetimeStr)
}

func fetchTaskNum(contestID string) (taskNum int, err error) {
	contestHome := "https://atcoder.jp/contests/" + contestID
	resp, err := grequests.Get(contestHome, &grequests.RequestOptions{UserAgent: ua})
	if err != nil {
		return
	}
	if !resp.Ok {
		return -1, fmt.Errorf("%d", resp.StatusCode)
	}

	root, err := html.Parse(resp)
	if err != nil {
		return
	}
	var f func(*html.Node) bool
	f = func(o *html.Node) bool {
		if o.DataAtom == atom.Tbody {
			for c := o.FirstChild; c != nil; c = c.NextSibling {
				if c.FirstChild != nil {
					taskNum++
				}
			}
			return true
		}
		for c := o.FirstChild; c != nil; c = c.NextSibling {
			if f(c) {
				return true
			}
		}
		return false
	}

	if !f(root) || taskNum == 0 {
		return -1, fmt.Errorf("题目数获取失败")
	}

	return
}

func login(username, password string) (session *grequests.Session, err error) {
	session = grequests.NewSession(&grequests.RequestOptions{
		UserAgent: ua,
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
		return nil, fmt.Errorf("POST %s return code %d", loginURL, resp.StatusCode)
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
		tokenInputJP  = "^入力例"
		tokenOutputJP = "^出力例"

		tokenInputEN  = "^Sample Input "
		tokenOutputEN = "^Sample Output "
	)

	inputRegex := regexp.MustCompile(tokenInputEN + `\d+`)
	outputRegex := regexp.MustCompile(tokenOutputEN + `\d+`)

	var f func(*html.Node)
	f = func(o *html.Node) {
		if o.Type == html.TextNode {
			if inputRegex.MatchString(o.Data) {
				raw := o.Parent.NextSibling.FirstChild.Data
				raw = strings.TrimSpace(raw)
				sampleIns = append(sampleIns, raw)
			} else if outputRegex.MatchString(o.Data) {
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
	spIdx := strings.LastIndexByte(problemName, '_')
	if spIdx < 0 {
		return fmt.Errorf("invlaid url %s", problemURL)
	}

	dirID, taskID := problemName[:spIdx], problemName[spIdx+1:]
	contestName := strings.ReplaceAll(dirID, "_", "-")

	// 生成目录
	dirPath := filepath.Join(contestDir, dirID, taskID) + "/"
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return err
	}

	submitURL := fmt.Sprintf("https://atcoder.jp/contests/%s/submit?taskScreenName=%s", contestName, problemName)
	statusURL := fmt.Sprintf("https://atcoder.jp/contests/%s/submissions?f.LanguageName=Go&f.Status=AC&f.Task=%s&orderBy=source_length", contestName, problemName)
	shortestURL := fmt.Sprintf("https://atcoder.jp/contests/%s/submissions?f.Status=AC&f.Task=%s&orderBy=source_length", contestName, problemName)

	// 创建 x.go
	mainFileContent := `package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)

}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
`
	mainFilePath := dirPath + taskID + ".go"
	if !isContest || taskID == "a" {
		// 比赛时，在 IDE 中打开 A 题
		defer open.Run(absPath(mainFilePath))
	}
	if _, err := os.Stat(mainFilePath); !os.IsNotExist(err) {
		open.Run(absPath(mainFilePath))
		return fmt.Errorf("文件已存在！")
	}
	if err := os.WriteFile(mainFilePath, []byte(mainFileContent), 0644); err != nil {
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
	testFileContent := fmt.Sprintf(`// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 题目：%s
// 提交：%s
// 对拍：%s
// 最短：%s
func Test_%s(t *testing.T) {
	testCases := [][2]string{%s
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
`, problemURL, submitURL, statusURL, shortestURL, taskID, examples)
	testFilePath := dirPath + taskID + "_test.go"
	if err := os.WriteFile(testFilePath, []byte(testFileContent), 0644); err != nil {
		return err
	}

	return nil
}

func genAtCoderContestTemplates(contestID string, taskNum, retryTimes int) error {
	if retryTimes == 0 {
		//submitURL := fmt.Sprintf("https://atcoder.jp/contests/%s/submit", contestID)
		//open.Run(submitURL)
		tasksPrintURL := fmt.Sprintf("https://atcoder.jp/contests/%s/tasks_print", contestID)
		open.Run(tasksPrintURL)
	}

	const ua = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.106 Safari/537.36"
	session := grequests.NewSession(&grequests.RequestOptions{
		UserAgent: ua,
	})

	tasksHome := fmt.Sprintf("https://atcoder.jp/contests/%s/tasks", contestID)
	revelSession, err := os.ReadFile("revel_session.txt")
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
		return genAtCoderContestTemplates(contestID, taskNum, retryTimes+1)
	}
	if !resp.Ok {
		return fmt.Errorf("未找到比赛或比赛尚未开始")
	}

	fmt.Println("开始解析样例输入输出")
	wg := &sync.WaitGroup{}
	defer wg.Wait()
	for i := 0; i < taskNum; i++ {
		wg.Add(1)
		// we don't want spent too much time on waiting responses one by one, so we use goroutine!
		go func(id byte) {
			defer wg.Done()
			defer func() {
				if err := recover(); err != nil {
					fmt.Println("[error]", string(id), err)
				}
			}()
			problemURL := fmt.Sprintf("https://atcoder.jp/contests/%[1]s/tasks/%[1]s_%[2]c", contestID, id)
			if err := genTemplates(session, problemURL, true); err != nil {
				fmt.Println("[error]", string(id), err)
				return
			}
			fmt.Println("[ok]", string(id), problemURL)
		}('a' + byte(i))
	}

	return nil
}

func GenAtCoderContestTemplates(contestID string) error {
	startTime, err := fetchStartTime(contestID)
	if err != nil {
		return err
	}

	taskNum, err := fetchTaskNum(contestID)
	if err != nil {
		return err
	}
	fmt.Printf("共 %d 道题目\n", taskNum)

	if t := time.Until(startTime); t > 0 {
		t += time.Second
		fmt.Println("sleep", t)
		time.Sleep(t)
	}

	return genAtCoderContestTemplates(contestID, taskNum, 0)
}

func GenAtCoderProblemTemplate(problemURL string) error {
	const ua = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.106 Safari/537.36"
	session := grequests.NewSession(&grequests.RequestOptions{
		UserAgent: ua,
	})
	return genTemplates(session, problemURL, false)
}
