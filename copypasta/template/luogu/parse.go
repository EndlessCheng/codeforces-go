package luogu

import (
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	"net/url"
	"strings"
)

const ua = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36"

func getJson(htmlStr string) (string, error) {
	const token = `<script id="lentille-context" type="application/json">`
	i := strings.Index(htmlStr, token)
	if i == -1 {
		return "", fmt.Errorf("invalid html content %s", htmlStr)
	}
	j := strings.Index(htmlStr[i+len(token):], `</script>`)
	if j == -1 {
		return "", fmt.Errorf("invalid html content %s", htmlStr)
	}
	return url.QueryUnescape(htmlStr[i+len(token) : i+len(token)+j])
}

func ParseExamples(problemURL string) (examples [][]string, err error) {
	resp, err := grequests.Get(problemURL, &grequests.RequestOptions{
		UserAgent:    ua,
		Headers:      map[string]string{"Referer": problemURL},
		UseCookieJar: true,
	})
	if err != nil {
		return
	}
	if !resp.Ok {
		return nil, fmt.Errorf("%d", resp.StatusCode)
	}
	jsonStr, err := getJson(resp.String())
	if err != nil {
		return
	}

	d := struct {
		CurrentData struct {
			Problem struct {
				Samples [][]string `json:"samples"`
			} `json:"problem"`
		} `json:"data"`
	}{}
	if err = json.Unmarshal([]byte(jsonStr), &d); err != nil {
		return
	}
	return d.CurrentData.Problem.Samples, nil
}
