package atcoder

import (
	"fmt"
	"github.com/levigross/grequests"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
)

const contestID = "abc147"

var contestDir = fmt.Sprintf("../../../dash/%s/", contestID)

func login(session *grequests.Session) (err error) {
	home := "https://atcoder.jp"
	resp, err := session.Get(home, nil)
	if err != nil {
		return
	}
	if !resp.Ok {
		err = fmt.Errorf("GET %s return code %d", home, resp.StatusCode)
		return
	}
	//session.RequestOptions.Cookies = resp.RawResponse.Cookies()

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

	apiLogin := "https://atcoder.jp/login"
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
	//session.RequestOptions.Cookies = resp.RawResponse.Cookies()
	return
}

func createDir(taskID byte) error {
	dirPath := contestDir + string(taskID)
	return os.MkdirAll(dirPath, os.ModePerm)
}

func parseTask(htmlURL string) (sampleIns, sampleOuts []string, err error) {
	resp, err := grequests.Get(htmlURL, &grequests.RequestOptions{
		UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36",
		Cookies: []*http.Cookie{{
			Name:  "REVEL_SESSION",
			Value: os.Getenv("REVEL_SESSION"),
		}},
	})
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
	session := grequests.NewSession(&grequests.RequestOptions{
		UserAgent:    "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36",
		UseCookieJar: true,
	})
	if err := login(session); err != nil {
		t.Fatal(err)
	}

	for taskID := byte('a'); taskID <= 'f'; taskID++ {
		if err := createDir(taskID); err != nil {
			t.Fatal(err)
		}

		htmlURL := fmt.Sprintf("https://atcoder.jp/contests/%[1]s/tasks/%[1]s_%[2]c", contestID, taskID)
		fmt.Println(string(taskID), htmlURL)
		ins, outs, err := parseTask(htmlURL)
		if err != nil {
			if strings.Contains(err.Error(), "404") {
				t.Fatal("未找到比赛或比赛尚未开始")
			}
			t.Error(err)
			continue
		}
		if len(ins) == 0 {
			t.Fatal("未找到比赛或比赛尚未开始")
		}

		for i, in := range ins {
			out := outs[i]
			if err := ioutil.WriteFile(fmt.Sprintf("%s%c/in%d.txt", contestDir, taskID, i+1), []byte(in), 0644); err != nil {
				t.Fatal(err)
			}
			if err := ioutil.WriteFile(fmt.Sprintf("%s%c/ans%d.txt", contestDir, taskID, i+1), []byte(out), 0644); err != nil {
				t.Fatal(err)
			}
		}
	}
}
