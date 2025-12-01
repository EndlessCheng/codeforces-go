package atcoder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	"os"
)

const problemsetJsonFolderPath = "./atcoder/resources/"
const problemsetJsonFilePath = problemsetJsonFolderPath + "problem-models.json"

func downloadJson() (err error) {
	resp, err := grequests.Get("https://kenkoooo.com/atcoder/resources/problem-models.json", &grequests.RequestOptions{UserAgent: ua})
	if err != nil {
		return
	}
	if !resp.Ok {
		return fmt.Errorf("downloadJson %d", resp.StatusCode)
	}
	if err := os.MkdirAll(problemsetJsonFolderPath, os.ModePerm); err != nil {
		return err
	}
	return resp.DownloadToFile(problemsetJsonFilePath)
}

// 获取题目的难度分
func getRating(problemName string) (rating int, err error) {
	get := func() (int, error) {
		jsonBS, err := os.ReadFile(problemsetJsonFilePath)
		if err != nil {
			return 0, err
		}
		mp := map[string]struct {
			Difficulty int `json:"difficulty"`
		}{}
		if err := json.NewDecoder(bytes.NewReader(jsonBS)).Decode(&mp); err != nil {
			return 0, err
		}
		if res, ok := mp[problemName]; ok {
			return max(res.Difficulty, 1), nil // res.Difficulty 可能 <= 0，所以和 1 取 max
		}
		return 0, nil
	}

	downloaded := false
	if _, er := os.Stat(problemsetJsonFilePath); os.IsNotExist(er) {
		if err = downloadJson(); err != nil {
			return
		}
		downloaded = true
	}
	rating, err = get()
	if err != nil {
		return
	}

	if rating == 0 && !downloaded {
		if err = downloadJson(); err != nil {
			return
		}
		rating, err = get()
		if err != nil {
			return
		}
	}

	return
}
