package atcoder

import (
	"fmt"
	"github.com/levigross/grequests"
	"golang.org/x/net/html"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

const contestID = "abc146"

const (
	ua = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36"
)

var contestDir = fmt.Sprintf("../../../dash/%s/", contestID)

func createDir(taskID byte) error {
	dirPath := contestDir + string(taskID)
	return os.MkdirAll(dirPath, os.ModePerm)
}

func parseTask(htmlURL string) (sampleIns, sampleOuts []string, err error) {
	resp, err := grequests.Get(htmlURL, &grequests.RequestOptions{
		UserAgent: ua,
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
	for taskID := byte('a'); taskID <= 'f'; taskID++ {
		if err := createDir(taskID); err != nil {
			t.Fatal(err)
		}
		htmlURL := fmt.Sprintf("https://atcoder.jp/contests/%[1]s/tasks/%[1]s_%[2]c", contestID, taskID)
		fmt.Println(string(taskID), htmlURL)
		ins, outs, err := parseTask(htmlURL)
		if err != nil {
			t.Error(htmlURL, err)
			continue
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
