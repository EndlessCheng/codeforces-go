package leetcode

import (
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	"github.com/skratchdot/open-golang/open"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

// TODO: 根据构造器传入的参数来预定义 var()
// https://leetcode-cn.com/contest/biweekly-contest-20/problems/apply-discount-every-n-orders/

func login(username, password string) (session *grequests.Session, err error) {
	const ua = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36"
	session = grequests.NewSession(&grequests.RequestOptions{
		UserAgent:    ua,
		UseCookieJar: true,
	})

	// "touch" csrfToken
	csrfTokenURL := fmt.Sprintf("https://%s/graphql/", host)
	var csrfJson interface{}
	if host == hostZH {
		csrfJson = map[string]interface{}{
			"operationName": "globalData",
			"query":         "query globalData {\n  feature {\n    questionTranslation\n    subscription\n    signUp\n    discuss\n    mockInterview\n    contest\n    store\n    book\n    chinaProblemDiscuss\n    socialProviders\n    studentFooter\n    cnJobs\n    __typename\n  }\n  userStatus {\n    isSignedIn\n    isAdmin\n    isStaff\n    isSuperuser\n    isTranslator\n    isPremium\n    isVerified\n    isPhoneVerified\n    isWechatVerified\n    checkedInToday\n    username\n    realName\n    userSlug\n    groups\n    jobsCompany {\n      nameSlug\n      logo\n      description\n      name\n      legalName\n      isVerified\n      permissions {\n        canInviteUsers\n        canInviteAllSite\n        leftInviteTimes\n        maxVisibleExploredUser\n        __typename\n      }\n      __typename\n    }\n    avatar\n    optedIn\n    requestRegion\n    region\n    activeSessionId\n    permissions\n    notificationStatus {\n      lastModified\n      numUnread\n      __typename\n    }\n    completedFeatureGuides\n    useTranslation\n    __typename\n  }\n  siteRegion\n  chinaHost\n  websocketUrl\n}\n",
		}
	} else {
		csrfJson = map[string]interface{}{
			"operationName": "fetchAllLeetcodeTemplates",
			"query":         "query fetchAllLeetcodeTemplates {\n  allLeetcodePlaygroundTemplates {\n    templateId\n    name\n    nameSlug\n    __typename\n  }\n}\n",
		}
	}
	resp, err := session.Post(csrfTokenURL, &grequests.RequestOptions{JSON: csrfJson})
	if err != nil {
		// maybe timeout
		fmt.Println("访问失败，重试", err)
		time.Sleep(time.Second)
		return login(username, password)
	}
	if !resp.Ok {
		return nil, fmt.Errorf("POST %s return code %d", csrfTokenURL, resp.StatusCode)
	}
	var csrfToken string
	for _, c := range resp.RawResponse.Cookies() {
		if c.Name == "csrftoken" {
			csrfToken = c.Value
			break
		}
	}
	if csrfToken == "" {
		return nil, fmt.Errorf("csrftoken not found in response")
	}

	// log in
	loginURL := fmt.Sprintf("https://%s/accounts/login/", host)
	loginData := map[string]string{
		"csrfmiddlewaretoken": csrfToken,
		"login":               username,
		"password":            password,
		"next":                "/",
	}
	if host == hostEN {
		// todo
		loginData["recaptcha_token"] = ""
	}
	resp, err = session.Post(loginURL, &grequests.RequestOptions{
		Data: loginData,
		Headers: map[string]string{
			"origin":  "https://" + host,
			"referer": "https://" + host + "/",
		},
	})
	if err != nil {
		fmt.Println("访问失败，重试", err)
		time.Sleep(time.Second)
		return login(username, password)
	}
	if !resp.Ok {
		return nil, fmt.Errorf("POST %s return code %d", loginURL, resp.StatusCode)
	}
	return
}

func fetchProblemURLs(session *grequests.Session) (problems []*problem, err error) {
	contestInfoURL := fmt.Sprintf("https://%s/contest/api/info/%s%d/", host, contestPrefix, contestID)
	resp, err := session.Get(contestInfoURL, nil)
	if err != nil {
		return
	}
	if !resp.Ok {
		return nil, fmt.Errorf("GET %s return code %d", contestInfoURL, resp.StatusCode)
	}

	d := struct {
		Contest struct {
			ID              int    `json:"id"`
			OriginStartTime int64  `json:"origin_start_time"`
			StartTime       int64  `json:"start_time"`
			Title           string `json:"title"`
		} `json:"contest"`
		Questions []struct {
			Credit    int    `json:"credit"`     // 得分/难度
			Title     string `json:"title"`      // 题目标题
			TitleSlug string `json:"title_slug"` // 题目链接
		} `json:"questions"`
		UserNum int `json:"user_num"` // 参赛人数
	}{}
	if err = resp.JSON(&d); err != nil {
		return
	}
	if d.Contest.StartTime == 0 {
		return nil, fmt.Errorf("未找到比赛或比赛尚未开始: %s%d", contestPrefix, contestID)
	}

	if sleepTime := time.Until(time.Unix(d.Contest.StartTime, 0)); sleepTime > 0 {
		sleepTime += 2 * time.Second // 消除误差
		fmt.Printf("%s尚未开始，等待中……\n%v\n", d.Contest.Title, sleepTime)
		time.Sleep(sleepTime)
		return fetchProblemURLs(session)
	}

	if len(d.Questions) == 0 {
		return nil, fmt.Errorf("题目链接为空: %s%d", contestPrefix, contestID)
	}

	fmt.Println("难度 标题")
	for _, q := range d.Questions {
		fmt.Printf("%3d %s\n", q.Credit, q.Title)
	}

	problems = make([]*problem, len(d.Questions))
	for i, q := range d.Questions {
		problems[i] = &problem{
			id:    string(byte('a' + i)),
			urlZH: fmt.Sprintf("https://%s/contest/%s%d/problems/%s/", hostZH, contestPrefix, contestID, q.TitleSlug),
			urlEN: fmt.Sprintf("https://%s/contest/%s%d/problems/%s/", hostEN, contestPrefix, contestID, q.TitleSlug),
		}
	}
	return
}

type problem struct {
	id            string
	urlZH         string
	urlEN         string
	defaultCode   string
	funcName      string
	isFuncProblem bool
	funcLos       []int
	sampleIns     [][]string
	sampleOuts    [][]string
	customComment string
}

// 解析一个样例输入或输出
func (p *problem) parseSampleText(text string, parseArgs bool) []string {
	text = strings.TrimSpace(text)
	if text == "" {
		return nil
	}

	lines := strings.Split(text, "\n")
	for i, s := range lines {
		lines[i] = strings.TrimSpace(s)
	}

	// 由于新版的样例不是这种格式了，这种特殊情况就不处理了
	// 见 https://leetcode-cn.com/contest/weekly-contest-121/problems/time-based-key-value-store/
	if !p.isFuncProblem {
		return lines
	}

	text = strings.Join(lines, "")

	// 包含中文的话，说明原始数据有误，截断首个中文字符之后的字符
	if idx := findNonASCII(text); idx != -1 {
		fmt.Println("[warn] 样例数据含有非 ASCII 字符，截断，原文为", text)
		text = text[:idx]
	}

	// 不含等号，说明只有一个参数
	if !parseArgs || !strings.Contains(text, "=") {
		return []string{text}
	}

	// TODO: 处理参数本身含有 = 的情况
	splits := strings.Split(text, "=")
	sample := make([]string, 0, len(splits)-1)
	for _, s := range splits[1 : len(splits)-1] {
		end := strings.LastIndexByte(s, ',')
		sample = append(sample, strings.TrimSpace(s[:end]))
	}
	sample = append(sample, strings.TrimSpace(splits[len(splits)-1]))
	if !p.isFuncProblem {
		sample = []string{strings.Join(sample, "\n") + "\n"}
	}
	return sample
}

func (p *problem) parsePossibleSampleTexts(texts []string, parseArgs bool) []string {
	for _, text := range texts {
		if sample := p.parseSampleText(text, parseArgs); len(sample) > 0 {
			return sample
		}
	}
	return nil
}

func (p *problem) parseHTML(session *grequests.Session) (err error) {
	defer func() {
		// visit htmlNode may cause panic
		if er := recover(); er != nil {
			err = fmt.Errorf("need fix: %v", er)
		}
	}()

	resp, err := session.Get(p.urlZH, nil)
	if err != nil {
		return
	}
	if !resp.Ok {
		return fmt.Errorf("GET %s return code %d", p.urlZH, resp.StatusCode)
	}

	rootNode, err := html.Parse(resp)
	if err != nil {
		return err
	}

	htmlNode := rootNode.FirstChild.NextSibling
	var bodyNode *html.Node
	for o := htmlNode.FirstChild; o != nil; o = o.NextSibling {
		if o.DataAtom == atom.Body {
			bodyNode = o
			break
		}
	}

	// parse defaultCode
	for o := bodyNode.FirstChild; o != nil; o = o.NextSibling {
		if o.DataAtom == atom.Script && o.FirstChild != nil {
			jsText := o.FirstChild.Data
			if start := strings.Index(jsText, "codeDefinition:"); start != -1 {
				end := strings.Index(jsText, "enableTestMode")
				jsonText := jsText[start+len("codeDefinition:") : end]
				jsonText = strings.TrimSpace(jsonText)
				jsonText = jsonText[:len(jsonText)-3] + "]" // remove , at end
				jsonText = strings.Replace(jsonText, `'`, `"`, -1)

				data := []struct {
					Value       string `json:"value"`
					DefaultCode string `json:"defaultCode"`
				}{}
				if err := json.Unmarshal([]byte(jsonText), &data); err != nil {
					return err
				}

				for _, template := range data {
					if template.Value == "golang" {
						p.defaultCode = strings.TrimSpace(template.DefaultCode)
						// 下面解析样例需要知道 p.isFuncProblem
						p.funcName, p.isFuncProblem, p.funcLos = parseCode(p.defaultCode)
						break
					}
				}
				break
			}
		}
	}
	if p.defaultCode == "" {
		fmt.Println("解析失败，未找到 Go 代码模板！")
	}

	// parse sample inputs and sample outputs
	// 注：官方描述可能会打错字（比如「输入」写成「输出」），这里只匹配第一个字
	parseData := func(nodes []*html.Node) (data string) {
		for _, o := range nodes {
			s := o.Data
			if o.DataAtom != 0 { // 如果有 tag（如 <code>），则取 tag 内的第一个元素
				s = o.FirstChild.Data
			}
			s = strings.TrimSpace(s)
			data += s
		}
		return
	}
	var f func(*html.Node)
	f = func(o *html.Node) {
		// 解析每个 <pre> 块内的文本（以中文为基准解析）
		// 需要判断 <pre> 的下一个子元素是否为 tag
		//     https://leetcode-cn.com/contest/weekly-contest-190/problems/max-dot-product-of-two-subsequences/
		//     https://leetcode-cn.com/contest/weekly-contest-212/problems/arithmetic-subarrays/
		// 有 tag 也不一定为 <strong> https://leetcode-cn.com/contest/weekly-contest-210/problems/split-two-strings-to-make-palindrome/
		if o.DataAtom == atom.Pre && o.FirstChild.DataAtom != 0 { // atom.Strong or atom.B
			if strings.HasPrefix(strings.TrimSpace(o.FirstChild.FirstChild.Data), "输") { // 输入 输出
				var inputNodes, outputNodes []*html.Node
				c := o.FirstChild.NextSibling
				for ; ; c = c.NextSibling {
					if c.DataAtom != 0 {
						s := strings.TrimSpace(c.FirstChild.Data)
						if strings.HasPrefix(s, "输") || strings.HasPrefix(s, "解") { // 输入 输出 解释
							break
						}
					}
					inputNodes = append(inputNodes, c)
				}
				rawInput := parseData(inputNodes)
				p.sampleIns = append(p.sampleIns, p.parseSampleText(rawInput, true))

				for c = c.NextSibling; c != nil; c = c.NextSibling { // 有时候没有解释，那么会通过 c != nil 来跳出循环
					if c.DataAtom != 0 {
						s := strings.TrimSpace(c.FirstChild.Data)
						if strings.HasPrefix(s, "输") || strings.HasPrefix(s, "解") { // 输入 输出 解释
							break
						}
					}
					outputNodes = append(outputNodes, c)
				}
				rawOutput := parseData(outputNodes)
				p.sampleOuts = append(p.sampleOuts, p.parseSampleText(rawOutput, true))
				return
			}
		}
		for c := o.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(bodyNode)
	if len(p.sampleIns) == 0 {
		// 没找到 <pre>，国服特殊比赛（春秋赛等）
		f = func(o *html.Node) {
			if o.DataAtom == atom.Div && o.FirstChild != nil && strings.Contains(o.FirstChild.Data, "示例") {
				raw := o.FirstChild.Data
				sp := strings.Split(raw, "`")
				for i, s := range sp {
					if strings.Contains(s, ">输入") {
						text := sp[i+1]
						if !p.isFuncProblem {
							// https://leetcode-cn.com/contest/season/2020-fall/problems/IQvJ9i/
							text += "\n" + sp[i+3] // 跳过 sp[i+2]
						}
						p.sampleIns = append(p.sampleIns, p.parseSampleText(text, true))
					} else if strings.Contains(s, ">输出") {
						p.sampleOuts = append(p.sampleOuts, p.parseSampleText(sp[i+1], true))
					}
				}
			}
			for c := o.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
		}
		f(bodyNode)
	}

	if len(p.sampleIns) != len(p.sampleOuts) {
		return fmt.Errorf("len(sampleIns) != len(sampleOuts) : %d != %d", len(p.sampleIns), len(p.sampleOuts))
	}
	if len(p.sampleIns) == 0 {
		return fmt.Errorf("解析失败，未找到样例输入输出！")
	}
	return nil
}

func (p *problem) createDir() error {
	return os.MkdirAll(contestDir+p.id, os.ModePerm)
}

var _firstMainFileOpened bool

func (p *problem) writeMainFile() error {
	imports := ""
	if strings.Contains(p.defaultCode, "Definition for") {
		// add imports
		imports = `
import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
`
	}
	p.defaultCode = strings.TrimSpace(p.defaultCode)
	fileContent := fmt.Sprintf(`package main
%s
%s%s
`, imports, p.customComment, p.defaultCode)

	filePath := contestDir + fmt.Sprintf("%[1]s/%[1]s.go", p.id)
	if !_firstMainFileOpened {
		_firstMainFileOpened = true
		defer open.Run(absPath(filePath))
	}
	return ioutil.WriteFile(filePath, []byte(fileContent), 0644)
}

func (p *problem) writeTestFile() error {
	exampleType := "[][]string"
	testUtilFunc := "testutil.RunLeetCodeFuncWithExamples"
	if !p.isFuncProblem {
		exampleType = "[][3]string"
		testUtilFunc = "testutil.RunLeetCodeClassWithExamples"
	}
	examples := ""
	for i, inArgs := range p.sampleIns {
		examples += "\n\t\t{\n\t\t\t"
		for _, arg := range inArgs {
			examples += "`" + arg + "`,"
			if p.isFuncProblem {
				examples += " "
			} else {
				examples += "\n\t\t\t"
			}
		}
		if p.isFuncProblem {
			examples += "\n\t\t\t"
		}
		if i < len(p.sampleOuts) {
			for _, arg := range p.sampleOuts[i] {
				examples += "`" + arg + "`,"
			}
		}
		examples += "\n\t\t},"
	}
	testStr := fmt.Sprintf(`// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [%s]")
	examples := %s{%s
		
	}
	targetCaseNum := 0
	if err := %s(t, %s, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// %s
`, p.id, exampleType, examples, testUtilFunc, p.funcName, p.urlZH)
	filePath := contestDir + fmt.Sprintf("%[1]s/%[1]s_test.go", p.id)
	return ioutil.WriteFile(filePath, []byte(testStr), 0644)
}

func handleProblems(session *grequests.Session, problems []*problem) error {
	if openWebPageZH {
		for _, p := range problems {
			if err := open.Run(p.urlZH); err != nil {
				fmt.Println("open err:", p.urlZH, err)
			}
		}
	}
	if openWebPageEN {
		for _, p := range problems {
			if err := open.Run(p.urlEN); err != nil {
				fmt.Println("open err:", p.urlEN, err)
			}
		}
	}

	for _, p := range problems {
		//if p.id != "a" {
		//	continue
		//}
		fmt.Println(p.id, p.urlZH)
		if err := p.parseHTML(session); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		p.defaultCode = modifyDefaultCode(p.defaultCode, p.funcLos, []modifyLineFunc{
			toGolangReceiverName,
			lowerArgsFirstChar,
			renameInputArgs,
			namedReturnFunc("ans"),
		}, "\t\n\treturn")

		if err := p.createDir(); err != nil {
			return err // IO
		}
		if err := p.writeMainFile(); err != nil {
			return err // IO
		}
		if err := p.writeTestFile(); err != nil {
			return err // IO
		}
	}
	return nil
}

func GenLeetCodeTests(username, password, customComment string) error {
	session, err := login(username, password)
	if err != nil {
		return err
	}
	fmt.Println("登录成功")
	//fmt.Println("登录成功", host, username)

	var problems []*problem
	for {
		problems, err = fetchProblemURLs(session)
		if err == nil {
			break
		}
		fmt.Println(err)
		time.Sleep(time.Second)
	}

	if customComment != "" {
		customComment += "\n"
	}
	for _, p := range problems {
		p.customComment = customComment
	}

	fmt.Println("题目链接获取成功，开始解析")
	return handleProblems(session, problems)
}

func GenLeetCodeSpecialTests(username, password, customComment string, urlsZHs []string) error {
	session, err := login(username, password)
	if err != nil {
		return err
	}
	fmt.Println(host, "登录成功")

	if customComment != "" {
		customComment += "\n"
	}
	problems := make([]*problem, len(urlsZHs))
	for i, url := range urlsZHs {
		problems[i] = &problem{
			id:            string(byte('a' + i)),
			urlZH:         url,
			customComment: customComment,
		}
	}

	return handleProblems(session, problems)
}
