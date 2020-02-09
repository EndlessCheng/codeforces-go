package atcoder

import (
	"fmt"
	"github.com/levigross/grequests"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"testing"
)

const contestID = "abc154"

func newSession(username, password string) (session *grequests.Session, err error) {
	const ua = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36"
	session = grequests.NewSession(&grequests.RequestOptions{
		UserAgent:    ua,
		UseCookieJar: true,
	})

	var resp *grequests.Response
	home := "https://atcoder.jp"
	for {
		resp, err = session.Get(home, nil)
		if err != nil {
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
		}
	}
	f(root)

	apiLogin := "https://atcoder.jp/login?continue=https%3A%2F%2Fatcoder.jp%2Fhome"
	resp, err = session.Post(apiLogin, &grequests.RequestOptions{
		Data: map[string]string{
			"username":   os.Getenv("ATCODER_USERNAME"),
			"password":   os.Getenv("ATCODER_PASSWORD"),
			"csrf_token": csrfToken,
		},
	})
	if err != nil {
		return
	}
	if !resp.Ok {
		err = fmt.Errorf("GET %s return code %d", apiLogin, resp.StatusCode)
		return
	}

	fmt.Println("登录成功")
	return
}

var contestDir = fmt.Sprintf("../../../dash/%s/", contestID)

func createDir(taskID byte) error {
	dirPath := contestDir + string(taskID)
	return os.MkdirAll(dirPath, os.ModePerm)
}

func parseTask(session *grequests.Session, htmlURL string) (sampleIns, sampleOuts []string, err error) {
	resp, err := session.Get(htmlURL, nil)
	if err != nil {
		return
	}
	if !resp.Ok {
		err = fmt.Errorf("GET %s return code %d", htmlURL, resp.StatusCode)
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

		tokenInputEN  = "Sample Input"
		tokenOutputEN = "Sample Output"
	)

	var f func(*html.Node)
	f = func(o *html.Node) {
		if o.Type == html.TextNode {
			if strings.Contains(o.Data, tokenInputJP) {
				raw := o.Parent.NextSibling.FirstChild.Data
				sampleIns = append(sampleIns, raw)
			} else if strings.Contains(o.Data, tokenOutputJP) {
				raw := o.Parent.NextSibling.FirstChild.Data
				sampleOuts = append(sampleOuts, raw)
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

func TestGenAtCoderTests(t *testing.T) {
	username := os.Getenv("ATCODER_USERNAME")
	password := os.Getenv("ATCODER_PASSWORD")
	session, err := newSession(username, password)
	if err != nil {
		t.Fatal(err)
	}

	//home := "https://atcoder.jp/home"
	//resp, err := session.Get(home, &grequests.RequestOptions{
	//	Headers: map[string]string{
	//		"upgrade-insecure-requests": "1",
	//	},
	//})
	//if err != nil {
	//	t.Fatal(err)
	//}
	//if !resp.Ok {
	//	t.Fatal("未找到比赛或比赛尚未开始")
	//}

	tasksHome := fmt.Sprintf("https://atcoder.jp/contests/%s/tasks", contestID)
	resp, err := session.Get(tasksHome, &grequests.RequestOptions{
		Cookies: []*http.Cookie{
			{
				Name: "REVEL_SESSION",
				// TODO
				Value: "",
			},
		},
		//Headers: map[string]string{
		//	"upgrade-insecure-requests": "1",
		//},
	})
	if err != nil {
		t.Fatal(err)
	}
	if !resp.Ok {
		t.Fatal("未找到比赛或比赛尚未开始")
	}

	fmt.Println("开始解析样例输入输出")
	wg := &sync.WaitGroup{}
	defer wg.Wait()
	for taskID := byte('a'); taskID <= 'f'; taskID++ {
		wg.Add(1)
		// we don't want spent too much time on waiting responses one by one, so we use goroutine !
		go func(id byte) {
			defer wg.Done()

			if err := createDir(id); err != nil {
				t.Error(err)
				return
			}

			// https://atcoder.jp/contests/abc154/tasks
			htmlURL := fmt.Sprintf("https://atcoder.jp/contests/%[1]s/tasks/%[1]s_%[2]c", contestID, id)
			ins, outs, err := parseTask(session, htmlURL)
			if err != nil {
				if strings.Contains(err.Error(), "404") {
					t.Fatal("未找到比赛或比赛尚未开始")
				}
				t.Error(err)
				return
			}
			if len(ins) == 0 {
				t.Error("未找到比赛或比赛尚未开始")
				return
			}

			for i, in := range ins {
				out := outs[i]
				if err := ioutil.WriteFile(fmt.Sprintf("%s%c/in%d.txt", contestDir, id, i+1), []byte(in), 0644); err != nil {
					t.Error(err)
				}
				if err := ioutil.WriteFile(fmt.Sprintf("%s%c/ans%d.txt", contestDir, id, i+1), []byte(out), 0644); err != nil {
					t.Error(err)
				}
			}

			fmt.Println("[ok]", string(id), htmlURL)
		}(taskID)
	}
}
