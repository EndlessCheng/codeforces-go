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
)

func login(username, password string) (session *grequests.Session, err error) {
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

	loginURL := "https://atcoder.jp/login?continue=https%3A%2F%2Fatcoder.jp%2Fhome"
	resp, err = session.Post(loginURL, &grequests.RequestOptions{
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
		return nil, fmt.Errorf("GET %s return code %d", loginURL, resp.StatusCode)
	}
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

func createDir(taskID byte) error {
	dirPath := contestDir + string(taskID)
	return os.MkdirAll(dirPath, os.ModePerm)
}

func GenAtCoderTests(username, password string) error {
	session, err := login(username, password)
	if err != nil {
		return err
	}
	fmt.Println("登录成功")

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
	revelSession, _ := ioutil.ReadFile("revel_session.txt")
	resp, err := session.Get(tasksHome, &grequests.RequestOptions{
		Cookies: []*http.Cookie{
			{
				Name:  "REVEL_SESSION",
				Value: string(revelSession),
			},
		},
		//Headers: map[string]string{
		//	"upgrade-insecure-requests": "1",
		//},
	})
	if err != nil {
		return err
	}
	if !resp.Ok {
		return fmt.Errorf("未找到比赛或比赛尚未开始")
	}

	fmt.Println("开始解析样例输入输出")
	wg := &sync.WaitGroup{}
	defer wg.Wait()
	for taskID := byte('a'); taskID <= 'f'; taskID++ {
		wg.Add(1)
		// we don't want spent too much time on waiting responses one by one, so we use goroutine!
		go func(id byte) {
			defer wg.Done()

			problemURL := fmt.Sprintf("https://atcoder.jp/contests/%[1]s/tasks/%[1]s_%[2]c", contestID, id)
			ins, outs, err := parseTask(session, problemURL)
			if err != nil {
				fmt.Fprintln(os.Stderr, string(id), err)
				return
			}

			if err := createDir(id); err != nil {
				panic(err)
			}

			for i, in := range ins {
				out := outs[i]
				if err := ioutil.WriteFile(fmt.Sprintf("%s%c/in%d.txt", contestDir, id, i+1), []byte(in), 0644); err != nil {
					panic(err)
				}
				if err := ioutil.WriteFile(fmt.Sprintf("%s%c/ans%d.txt", contestDir, id, i+1), []byte(out), 0644); err != nil {
					panic(err)
				}
			}

			fmt.Println("[ok]", string(id), problemURL)
		}(taskID)
	}

	return nil
}
