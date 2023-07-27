package leetcode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	"github.com/skratchdot/open-golang/open"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"io"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

const host = "leetcode.cn"
const graphqlURL = "https://" + host + "/graphql"

// 使用用户名和密码登录
func login(username, password string) (session *grequests.Session, err error) {
	const ua = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36"
	session = grequests.NewSession(&grequests.RequestOptions{
		UserAgent:    ua,
		UseCookieJar: true,
	})

	// "touch" csrfToken
	resp, err := session.Post(graphqlURL, &grequests.RequestOptions{
		JSON: map[string]interface{}{
			"operationName": "globalData",
			"query":         "query globalData {\n  feature {\n    questionTranslation\n    subscription\n    signUp\n    discuss\n    mockInterview\n    contest\n    store\n    book\n    chinaProblemDiscuss\n    socialProviders\n    studentFooter\n    cnJobs\n    __typename\n  }\n  userStatus {\n    isSignedIn\n    isAdmin\n    isStaff\n    isSuperuser\n    isTranslator\n    isPremium\n    isVerified\n    isPhoneVerified\n    isWechatVerified\n    checkedInToday\n    username\n    realName\n    userSlug\n    groups\n    jobsCompany {\n      nameSlug\n      logo\n      description\n      name\n      legalName\n      isVerified\n      permissions {\n        canInviteUsers\n        canInviteAllSite\n        leftInviteTimes\n        maxVisibleExploredUser\n        __typename\n      }\n      __typename\n    }\n    avatar\n    optedIn\n    requestRegion\n    region\n    activeSessionId\n    permissions\n    notificationStatus {\n      lastModified\n      numUnread\n      __typename\n    }\n    completedFeatureGuides\n    useTranslation\n    __typename\n  }\n  siteRegion\n  chinaHost\n  websocketUrl\n}\n",
		},
	})
	if err != nil {
		// maybe timeout
		fmt.Println("访问失败，重试", err)
		time.Sleep(time.Second)
		return login(username, password)
	}
	if !resp.Ok {
		return nil, fmt.Errorf("POST %s return code %d", graphqlURL, resp.StatusCode)
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
	resp, err = session.Post(loginURL, &grequests.RequestOptions{
		Data: map[string]string{
			"csrfmiddlewaretoken": csrfToken,
			"login":               username,
			"password":            password,
			"next":                "/",
		},
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

	u, _ := url.Parse(loginURL)
	for _, cookie := range session.HTTPClient.Jar.Cookies(u) {
		if cookie.Name == "LEETCODE_SESSION" {
			return
		}
	}

	return nil, fmt.Errorf("登录失败：账号或密码错误")
}

// 获取题目信息（含题目链接）
// contestTag 如 "weekly-contest-200"，可以从比赛链接中获取
func fetchProblemURLs(session *grequests.Session, contestTag string) (problems []*problem, err error) {
	contestInfoURL := fmt.Sprintf("https://%s/contest/api/info/%s/", host, contestTag)
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
		Registered bool `json:"registered"` // 是否报名
		UserNum    int  `json:"user_num"`   // 参赛人数
	}{}
	if err = resp.JSON(&d); err != nil {
		return
	}
	if d.Contest.StartTime == 0 {
		return nil, fmt.Errorf("未找到比赛或比赛尚未开始: %s", contestTag)
	}

	//fmt.Println("当前报名人数", d.UserNum)

	if sleepTime := time.Until(time.Unix(d.Contest.StartTime, 0)); sleepTime > 0 {
		if !d.Registered {
			fmt.Printf("该账号尚未报名%s\n", d.Contest.Title)
			return
		}

		sleepTime += 500 * time.Millisecond // 消除误差
		fmt.Printf("%s尚未开始，等待中……\n%v\n", d.Contest.Title, sleepTime)
		time.Sleep(sleepTime)
		return fetchProblemURLs(session, contestTag)
	}

	if len(d.Questions) == 0 {
		return nil, fmt.Errorf("题目链接为空: %s", contestTag)
	}

	fmt.Println("难度 标题")
	for _, q := range d.Questions {
		fmt.Printf("%3d %s\n", q.Credit, q.Title)
	}

	problems = make([]*problem, len(d.Questions))
	for i, q := range d.Questions {
		problems[i] = &problem{
			id:  string(byte('a' + i)),
			url: fmt.Sprintf("https://%s/contest/%s/problems/%s/", host, contestTag, q.TitleSlug),

			isFuncProblem: true,
		}
	}
	return
}

// 获取力扣杯的题目链接
// slug 如 "2020-fall", "2021-spring"
func fetchSeasonProblemURLs(session *grequests.Session, slug string, isSolo bool) (problems []*problem, err error) {
	resp, err := session.Post(graphqlURL, &grequests.RequestOptions{
		JSON: map[string]interface{}{
			"operationName": "contestGroup",
			"query":         "query contestGroup($slug: String!) {\n  contestGroup(slug: $slug) {\n    title\n    titleCn\n    contestCount\n    contests {\n      title\n      titleCn\n      titleSlug\n      startTime\n      duration\n      registered\n      questions {\n        title\n        titleCn\n        titleSlug\n        credit\n        questionId\n        __typename\n      }\n      teamSettings {\n        maxTeamSize\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n",
			"variables": map[string]string{
				"slug": slug,
			},
		},
	})
	if err != nil {
		return
	}
	if !resp.Ok {
		return nil, fmt.Errorf("POST %s return code %d", graphqlURL, resp.StatusCode)
	}

	type contestInfo struct {
		TitleSlug string `json:"titleSlug"` // 2020-fall-solo
		StartTime int64  `json:"startTime"`
		Questions []struct {
			Credit    int    `json:"credit"`    // 得分/难度
			Title     string `json:"title"`     // 题目标题
			TitleSlug string `json:"titleSlug"` // 题目链接
		} `json:"questions"`
	}

	d := struct {
		Data struct {
			ContestGroup struct {
				Contests []contestInfo `json:"contests"`
			} `json:"contestGroup"`
		} `json:"data"`
	}{}
	if err = resp.JSON(&d); err != nil {
		return
	}

	var contest contestInfo
	for _, c := range d.Data.ContestGroup.Contests {
		if isSolo {
			if strings.Contains(c.TitleSlug, "solo") {
				contest = c
				break
			}
		} else {
			if strings.Contains(c.TitleSlug, "team") {
				contest = c
				break
			}
		}
	}

	if contest.TitleSlug == "" {
		return nil, fmt.Errorf("未找到比赛 %s！", slug)
	}

	if sleepTime := time.Until(time.Unix(contest.StartTime, 0)); sleepTime > 0 {
		sleepTime += 2 * time.Second // 消除误差
		fmt.Printf("%s尚未开始，等待中……\n%v\n", contest.TitleSlug, sleepTime)
		time.Sleep(sleepTime)
		return fetchSeasonProblemURLs(session, slug, isSolo)
	}

	if len(contest.Questions) == 0 {
		return nil, fmt.Errorf("题目链接为空: %s", contest.TitleSlug)
	}

	fmt.Println("难度 标题")
	for _, q := range contest.Questions {
		fmt.Printf("%3d %s\n", q.Credit, q.Title)
	}

	problems = make([]*problem, len(contest.Questions))
	for i, q := range contest.Questions {
		problems[i] = &problem{
			id:  string(byte('a' + i)),
			url: fmt.Sprintf("https://%s/contest/season/%s/problems/%s/", host, slug, q.TitleSlug),
		}
	}
	return
}

//

type problem struct {
	id      string
	url     string
	openURL bool

	defaultCode   string
	funcName      string
	isFuncProblem bool
	needMod       bool
	funcLos       []int
	customComment string

	sampleIns  [][]string
	sampleOuts [][]string

	contestDir string
}

// 解析一个样例输入或输出
func (p *problem) parseSampleText(text string, parseArgs bool) []string {
	text = strings.ReplaceAll(text, " ", " ") // 替换 NBSP 为正常空格

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

// 获取题目样例和代码
func (p *problem) parseHTML(session *grequests.Session) (err error) {
	defer func() {
		// visit htmlNode may cause panic
		if er := recover(); er != nil {
			err = fmt.Errorf("need fix: %v", er)
		}
	}()

	resp, err := session.Get(p.url, nil)
	if err != nil {
		return
	}
	if !resp.Ok {
		return fmt.Errorf("GET %s return code %d", p.url, resp.StatusCode)
	}

	htmlText, _ := io.ReadAll(resp)
	p.needMod = bytes.Contains(htmlText, []byte("取余")) || bytes.Contains(htmlText, []byte("取模")) || bytes.Contains(htmlText, []byte("答案可能很大"))

	rootNode, err := html.Parse(bytes.NewReader(htmlText))
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

	// 下面这段是旧的解析方案测试过程中发现的 bug
	// 新版解析方案改用英文 + strong tag 解析，没有这些 bug
	//     提取并解析每个 <pre> 块内的文本
	//     需要判断 <pre> 的下一个子元素是否为 tag
	//         https://leetcode-cn.com/contest/weekly-contest-190/problems/max-dot-product-of-two-subsequences/
	//         https://leetcode-cn.com/contest/weekly-contest-212/problems/arithmetic-subarrays/
	//     有 tag 也不一定为 <strong>
	//         <img> https://leetcode-cn.com/contest/weekly-contest-103/problems/snakes-and-ladders/
	//         <b> https://leetcode-cn.com/contest/weekly-contest-210/problems/split-two-strings-to-make-palindrome/
	//         <code> https://leetcode-cn.com/contest/weekly-contest-163/problems/shift-2d-grid/
	//         https://leetcode.cn/contest/weekly-contest-345/problems/maximum-number-of-moves-in-a-grid/
	//         https://leetcode.cn/contest/biweekly-contest-36/problems/find-valid-matrix-given-row-and-column-sums/
	//     找到第一个文本，这样写是因为可能有额外的嵌套 tag https://leetcode-cn.com/contest/weekly-contest-163/problems/shift-2d-grid/

	var parseNode func(*html.Node)
	parseNode = func(o *html.Node) {
		// 寻找 <strong>Input:</strong> 和 <strong>Output:</strong>
		// 不要用中文的，国服偶尔会破坏这个规则
		const inputToken = "Input"
		const outputToken = "Output"
		const explanationToken = "Explanation"
		const explanationToken2 = "Explaination"
		// 设计题末尾没有 ':'
		tidy := func(data string) string {
			data = strings.TrimSpace(data)
			if data != "" && data[len(data)-1] == ':' {
				data = data[:len(data)-1]
			}
			return data
		}
		isInput := func(data string) bool { return tidy(data) == inputToken }
		isOutput := func(data string) bool { return tidy(data) == outputToken }
		isExplanation := func(data string) bool { return tidy(data) == explanationToken || tidy(data) == explanationToken2 }
		if o.DataAtom == atom.Strong && o.FirstChild != nil && (isInput(o.FirstChild.Data) || isOutput(o.FirstChild.Data)) {
			curNode := o.FirstChild
			// 提取输入输出信息
			rawData := &strings.Builder{}
			var parseTextAfterStrong func(*html.Node) bool
			parseTextAfterStrong = func(o *html.Node) bool {
				if o != curNode && o.Type == html.TextNode {
					// 不再继续寻找
					if isOutput(o.Data) || isExplanation(o.Data) {
						return true
					}
					rawData.WriteString(o.Data)
				}
				for c := o.FirstChild; c != nil; c = c.NextSibling {
					if parseTextAfterStrong(c) {
						return true
					}
				}
				return false
			}
			for c := o; c != nil; c = c.NextSibling {
				if parseTextAfterStrong(c) {
					break
				}
			}

			if isInput(o.FirstChild.Data) {
				p.sampleIns = append(p.sampleIns, p.parseSampleText(rawData.String(), true))
			} else if isOutput(o.FirstChild.Data) {
				p.sampleOuts = append(p.sampleOuts, p.parseSampleText(rawData.String(), true))
			} else {
				panic("这不可能。代码有误！")
			}
		}
		for c := o.FirstChild; c != nil; c = c.NextSibling {
			parseNode(c)
		}
	}
	parseNode(bodyNode)

	if len(p.sampleIns) == 0 {
		// 没找到 <pre>，国服特殊比赛（春秋赛等）
		parseNode = func(o *html.Node) {
			if o.DataAtom == atom.Div && o.FirstChild != nil && strings.Contains(o.FirstChild.Data, "示例") {
				raw := o.FirstChild.Data
				sp := strings.Split(raw, "`")
				for i, s := range sp {
					if strings.Contains(s, ">输入") || strings.Contains(s, "> 输入") {
						text := sp[i+1]
						if !p.isFuncProblem {
							// https://leetcode-cn.com/contest/season/2020-fall/problems/IQvJ9i/
							text += "\n" + sp[i+3] // 跳过 sp[i+2]
						}
						p.sampleIns = append(p.sampleIns, p.parseSampleText(text, true))
					} else if strings.Contains(s, ">输出") || strings.Contains(s, "> 输出") {
						p.sampleOuts = append(p.sampleOuts, p.parseSampleText(sp[i+1], true))
					}
				}
			}
			for c := o.FirstChild; c != nil; c = c.NextSibling {
				parseNode(c)
			}
		}
		parseNode(bodyNode)
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
	return os.MkdirAll(p.contestDir+p.id, os.ModePerm)
}

func (p *problem) writeMainFile() error {
	imports := ""
	if strings.Contains(p.defaultCode, "Definition for") {
		imports = `
import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
`
	}
	p.defaultCode = strings.TrimSpace(p.defaultCode)
	fileContent := fmt.Sprintf(`package main
%s
%s%s
`, imports, p.customComment, p.defaultCode)

	filePath := p.contestDir + fmt.Sprintf("%[1]s/%[1]s.go", p.id)
	if p.id == "a" {
		defer open.Run(absPath(filePath)) // 打开第一道题的文件
	}
	return os.WriteFile(filePath, []byte(fileContent), 0644)
}

func (p *problem) writeTestFile() error {
	logInfo := ""
	testUtilFunc := "testutil.RunLeetCodeFuncWithFile"
	if !p.isFuncProblem {
		logInfo += "\n\t" + `t.Log("记得初始化所有全局变量")`
		testUtilFunc = "testutil.RunLeetCodeClassWithFile"
	}

	u := p.url
	i := strings.Index(u, "contest/")
	j := strings.Index(u, "problems/")
	problemURL := u[:i] + u[j:]

	testStr := fmt.Sprintf(`// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test_%s(t *testing.T) {%s
	targetCaseNum := 0 // -1
	if err := %s(t, %s, "%s.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// %s
// %s
`, p.id, logInfo, testUtilFunc, p.funcName, p.id, p.url, problemURL)

	filePath := p.contestDir + fmt.Sprintf("%[1]s/%[1]s_test.go", p.id)
	return os.WriteFile(filePath, []byte(testStr), 0644)
}

func (p *problem) writeTestDataFile() error {
	lines := []string{}
	for i, inArgs := range p.sampleIns {
		lines = append(lines, inArgs...)
		if i < len(p.sampleOuts) {
			lines = append(lines, p.sampleOuts[i]...)
		}
		lines = append(lines, "") // empty line for clarity
	}
	lines = append(lines, "", "", "")
	testDataStr := strings.Join(lines, "\n")

	filePath := p.contestDir + fmt.Sprintf("%[1]s/%[1]s.txt", p.id)
	return os.WriteFile(filePath, []byte(testDataStr), 0644)
}

func handleProblems(session *grequests.Session, problems []*problem) error {
	wg := &sync.WaitGroup{}
	wg.Add(1 + len(problems))

	go func() {
		defer wg.Done()
		for _, p := range problems {
			if p.openURL {
				if err := open.Run(p.url); err != nil {
					fmt.Println("open err:", p.url, err)
				}
			}
		}
	}()

	for _, p := range problems {
		fmt.Println(p.id, p.url)

		go func(p *problem) {
			defer wg.Done()

			if err := p.parseHTML(session); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}

			customFuncContent := "\t\n" // 空换行
			if p.needMod {
				customFuncContent += "\t\n\t\n\tans = (ans%mod + mod) % mod\n"
			}
			customFuncContent += "\treturn" // 补上 return

			p.defaultCode = modifyDefaultCode(p.defaultCode, p.funcLos, []modifyLineFunc{
				toGolangReceiverName,
				lowerArgsFirstChar,
				renameInputArgs,
				namedReturnFunc("ans"),
			}, customFuncContent)

			if p.needMod {
				p.defaultCode = "const mod = 1_000_000_007\n\n" + p.defaultCode
			}

			if err := p.createDir(); err != nil {
				fmt.Println("createDir err:", p.url, err)
				return
			}
			if err := p.writeMainFile(); err != nil {
				fmt.Println("writeMainFile err:", p.url, err)
			}
			if err := p.writeTestFile(); err != nil {
				fmt.Println("writeTestFile err:", p.url, err)
			}
			if err := p.writeTestDataFile(); err != nil {
				fmt.Println("writeTestFile err:", p.url, err)
			}
		}(p)
	}

	wg.Wait()
	return nil
}

func updateComment(cmt string) string {
	if cmt != "" {
		if !strings.HasPrefix(cmt, "//") {
			cmt = "// " + cmt
		}
		cmt += "\n"
	}
	return cmt
}

// 获取题目信息（含题目链接）
// contestTag 如 "weekly-contest-200"，可以从比赛链接中获取
func GenLeetCodeTests(username, password, contestTag string, openWebPage bool, contestDir, customComment string) error {
	session, err := login(username, password)
	if err != nil {
		return err
	}
	fmt.Println("登录成功")
	//fmt.Println("登录成功", host, username)

	var problems []*problem
	for {
		problems, err = fetchProblemURLs(session, contestTag)
		if err == nil {
			break
		}
		fmt.Println(err)
		time.Sleep(500 * time.Millisecond)
	}

	if len(problems) == 0 {
		return nil
	}

	customComment = updateComment(customComment)

	for _, p := range problems {
		p.openURL = openWebPage
		p.customComment = customComment
		p.contestDir = contestDir
	}

	fmt.Println("题目链接获取成功，开始解析")
	return handleProblems(session, problems)
}

const (
	SeasonSpring = "spring"
	SeasonFall   = "fall"
)

// 获取力扣杯题目信息
// slug 如 "2020-fall", "2021-spring"
func GenLeetCodeSeasonTests(username, password, slug string, isSolo, openWebPage bool, contestDir, customComment string) error {
	session, err := login(username, password)
	if err != nil {
		return err
	}
	fmt.Println("登录成功")

	var problems []*problem
	for {
		problems, err = fetchSeasonProblemURLs(session, slug, isSolo)
		if err == nil {
			break
		}
		fmt.Println(err)
		time.Sleep(time.Second)
	}

	customComment = updateComment(customComment)

	for _, p := range problems {
		p.openURL = openWebPage
		p.customComment = customComment
		p.contestDir = contestDir
	}

	fmt.Println("题目链接获取成功，开始解析")
	return handleProblems(session, problems)
}

func GenLeetCodeSpecialTests(username, password string, urlsZHs []string, openWebPage bool, contestDir, customComment string) error {
	session, err := login(username, password)
	if err != nil {
		return err
	}
	fmt.Println(host, "登录成功")

	customComment = updateComment(customComment)

	problems := make([]*problem, len(urlsZHs))
	for i, url := range urlsZHs {
		problems[i] = &problem{
			id:            string(byte('a' + i)),
			url:           url,
			openURL:       openWebPage,
			customComment: customComment,
			contestDir:    contestDir,
		}
	}

	return handleProblems(session, problems)
}
