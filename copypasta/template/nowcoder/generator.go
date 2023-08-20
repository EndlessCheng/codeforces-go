package nowcoder

import (
	"fmt"
	"github.com/levigross/grequests"
	"github.com/skratchdot/open-golang/open"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const ua = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36"

func fetchWeeklyContestId() (id int, err error) {
	const home = "https://ac.nowcoder.com/acm/contest/vip-index?rankTypeFilter=-1&onlyCreateFilter=false&topCategoryFilter=13&categoryFilter=19&signUpFilter=&orderType=NO"
	resp, err := grequests.Get(home, &grequests.RequestOptions{UserAgent: ua})
	if err != nil {
		return
	}
	if !resp.Ok {
		return 0, fmt.Errorf("%d", resp.StatusCode)
	}

	root, err := html.Parse(resp)
	if err != nil {
		return
	}
	var f func(*html.Node) bool
	f = func(o *html.Node) bool {
		if strings.HasPrefix(o.Data, fmt.Sprintf("牛客周赛 Round ")) {
			for _, a := range o.Parent.Attr {
				if a.Key == "href" {
					i := strings.LastIndex(a.Val, "/")
					id, err = strconv.Atoi(a.Val[i+1:])
					return true
				}
			}
			panic(-1)
		}
		for c := o.FirstChild; c != nil; c = c.NextSibling {
			if f(c) {
				return true
			}
		}
		return false
	}
	if !f(root) {
		return 0, fmt.Errorf("无法找到 id")
	}
	return
}

func fetchStartTime(contestID int) (t time.Time, contestName string, err error) {
	home := fmt.Sprintf("https://ac.nowcoder.com/acm/contest/%d", contestID)
	resp, err := grequests.Get(home, &grequests.RequestOptions{UserAgent: ua})
	if err != nil {
		return
	}
	if !resp.Ok {
		return time.Time{}, "", fmt.Errorf("%d", resp.StatusCode)
	}

	htmlStr := resp.String()

	const tokenName = `"competitionName_var":"`
	i := strings.Index(htmlStr, tokenName)
	if i == -1 {
		return time.Time{}, "", fmt.Errorf("invalid html content %s", htmlStr)
	}
	j := strings.Index(htmlStr[i+len(tokenName):], "\"")
	contestName = htmlStr[i+len(tokenName) : i+len(tokenName)+j]

	const tokenTime = `"startTime":`
	i = strings.Index(htmlStr, tokenTime)
	if i == -1 {
		return time.Time{}, "", fmt.Errorf("invalid html content %s", htmlStr)
	}

	tsStr := htmlStr[i+len(tokenTime) : i+len(tokenTime)+10]
	ts, err := strconv.Atoi(tsStr)
	if err != nil {
		return
	}
	return time.Unix(int64(ts), 0), contestName, nil
}

func login(emailOrPhone, cipherPwd string) (session *grequests.Session, err error) {
	session = grequests.NewSession(&grequests.RequestOptions{
		UserAgent:    ua,
		UseCookieJar: true,
	})
	resp, err := session.Post("https://www.nowcoder.com/nccommon/login/do", &grequests.RequestOptions{
		Data: map[string]string{
			"email":     emailOrPhone,
			"cipherPwd": cipherPwd,
		},
		Headers: map[string]string{
			"Origin": "https://www.nowcoder.com",
		},
	})
	if err != nil {
		return
	}
	if !resp.Ok {
		return nil, fmt.Errorf("%d", resp.StatusCode)
	}
	d := struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}{}
	if err = resp.JSON(&d); err != nil {
		return
	}
	if d.Code != 0 {
		return nil, fmt.Errorf("login: %s", d.Msg)
	}
	fmt.Println("【登录成功】", emailOrPhone)
	return
}

//func fetchRegisterState(session *grequests.Session, contestID int) (registered bool, err error) {
//	contestUrl := "https://ac.nowcoder.com/acm/contest/" + strconv.Itoa(contestID)
//	resp, err := session.Get(contestUrl, &grequests.RequestOptions{
//		Headers: map[string]string{
//			//"Origin":  "https://ac.nowcoder.com",
//			"Referer": "https://ac.nowcoder.com/acm/contest/" + strconv.Itoa(contestID),
//		},
//	})
//	if err != nil {
//		return
//	}
//	if !resp.Ok {
//		return false, fmt.Errorf("%d", resp.StatusCode)
//	}
//}

func register(session *grequests.Session, contestID int) (err error) {
	const api = "https://ac.nowcoder.com/acm/contest/sign-up-team"
	resp, err := session.Post(api, &grequests.RequestOptions{
		Data: map[string]string{
			"contestId": strconv.Itoa(contestID),
		},
		Headers: map[string]string{
			"Origin":  "https://ac.nowcoder.com",
			"Referer": "https://ac.nowcoder.com/acm/contest/" + strconv.Itoa(contestID),
		},
	})
	if err != nil {
		return
	}
	if !resp.Ok {
		return fmt.Errorf("%d", resp.StatusCode)
	}
	d := struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}{}
	if err = resp.JSON(&d); err != nil {
		return
	}
	if d.Code != 0 && d.Code != 1 {
		return fmt.Errorf("register: %s", d.Msg)
	}
	fmt.Println("register", d.Msg)
	return
}

func fetchProblemNumber(session *grequests.Session, contestID int) (problemNum int, err error) {
	api := fmt.Sprintf("https://ac.nowcoder.com/acm/contest/problem-list?id=%d", contestID)
	resp, err := session.Get(api, &grequests.RequestOptions{
		Headers: map[string]string{
			//"Origin": "https://ac.nowcoder.com",
			"Referer": "https://ac.nowcoder.com/acm/contest/" + strconv.Itoa(contestID),
		},
	})
	if err != nil {
		return
	}
	if !resp.Ok {
		return 0, fmt.Errorf("%d", resp.StatusCode)
	}

	d := struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			BasicInfo struct {
				ProblemCount int `json:"problemCount"`
			} `json:"basicInfo"`
		} `json:"data"`
	}{}
	if err = resp.JSON(&d); err != nil {
		return
	}
	// {"msg":"contest not ready! ","code":1}
	if d.Code != 0 {
		return 0, fmt.Errorf("fetchProblems: %s", d.Msg)
	}

	// 为 0 则比赛尚未开始
	problemNum = d.Data.BasicInfo.ProblemCount
	return
}

// examples 为输入输出交替
func parseExamples(session *grequests.Session, problemURL string) (examples []string, err error) {
	resp, err := session.Get(problemURL, &grequests.RequestOptions{
		Headers: map[string]string{
			//"Origin": "https://ac.nowcoder.com",
			"Referer": problemURL,
		},
	})
	if err != nil {
		return
	}
	if !resp.Ok {
		return nil, fmt.Errorf("%d", resp.StatusCode)
	}

	root, err := html.Parse(resp)
	if err != nil {
		return
	}

	var f func(o *html.Node)
	f = func(o *html.Node) {
		if o.DataAtom == atom.Textarea {
			for _, attribute := range o.Attr {
				if attribute.Key == "data-clipboard-text-id" {
					examples = append(examples, strings.TrimSpace(o.FirstChild.Data))
					break
				}
			}
		}
		for c := o.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(root)
	return
}

func GenNowCoderTemplates(emailOrPhone, cipherPwd, contestDir string, contestID int, openWebPage bool) (err error) {
	startTime, contestName, err := fetchStartTime(contestID)
	if err != nil {
		return
	}

	session, err := login(emailOrPhone, cipherPwd)
	if err != nil {
		return
	}

	if err = register(session, contestID); err != nil {
		return
	}

	if t := time.Until(startTime); t > 0 {
		t += 1 * time.Second
		fmt.Println(contestName, t)
		time.Sleep(t)
	}

	var problemNum int
	for {
		problemNum, err = fetchProblemNumber(session, contestID)
		if problemNum > 0 {
			break
		}
		fmt.Println("fetchProblems", err)
		time.Sleep(time.Second)
	}
	fmt.Println("本次比赛有", problemNum, "题")

	wg := &sync.WaitGroup{}
	wg.Add(1 + problemNum)

	go func() {
		defer wg.Done()
		if !openWebPage {
			return
		}
		for id := byte('a'); id < 'a'+byte(problemNum); id++ {
			problemURL := fmt.Sprintf("https://ac.nowcoder.com/acm/contest/%d/%c", contestID, id)
			if er := open.Run(problemURL); er != nil {
				fmt.Println("open err:", problemURL, er)
			}
		}
	}()

	for id := byte('a'); id < 'a'+byte(problemNum); id++ {
		go func(id byte) {
			defer wg.Done()

			problemURL := fmt.Sprintf("https://ac.nowcoder.com/acm/contest/%d/%c", contestID, id)
			var examples []string
			for {
				examples, err = parseExamples(session, problemURL)
				if err != nil {
					fmt.Println("[error] parseCodeAndExamples", problemURL, err)
					return
				}
				if len(examples) > 0 {
					break
				}
				fmt.Println("样例为空，重试...")
				time.Sleep(time.Second)
			}

			problemDir := contestDir + string(id) + "/"
			if err = os.MkdirAll(problemDir, os.ModePerm); err != nil {
				fmt.Println(err)
				return
			}

			for j, fileName := range []string{"main.go", "main_test.go"} {
				goFilePath := problemDir + strings.Replace(fileName, "main", string(id), 1)
				if err = copyFile(goFilePath, "../"+fileName); err != nil {
					fmt.Println(err)
					return
				}
				if id == 'a' && j == 0 {
					open.Run(absPath(goFilePath))
				}
			}

			for i := 0; i < len(examples); i += 2 {
				filePath := problemDir + fmt.Sprintf("in%d.txt", i/2+1)
				if err = os.WriteFile(filePath, []byte(examples[i]), 0644); err != nil {
					fmt.Println(err)
					return
				}
				filePath = problemDir + fmt.Sprintf("ans%d.txt", i/2+1)
				if err = os.WriteFile(filePath, []byte(examples[i+1]), 0644); err != nil {
					fmt.Println(err)
					return
				}
			}
		}(id)
	}

	wg.Wait()
	fmt.Println("Done.")
	return
}
