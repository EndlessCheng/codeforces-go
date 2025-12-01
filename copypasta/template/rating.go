package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	"os"
	"strconv"
)

const problemsetJsonFilePath = "problemset/problemset.json"

func downloadJson() (err error) {
	resp, err := grequests.Get("https://codeforces.com/api/problemset.problems?lang=en", &grequests.RequestOptions{UserAgent: ua})
	if err != nil {
		return
	}
	if !resp.Ok {
		return fmt.Errorf("downloadJson %d", resp.StatusCode)
	}
	return resp.DownloadToFile(problemsetJsonFilePath)
}

// 获取题目的难度分
func getRating(contestID, problemID string) (rating int, err error) {
	get := func() (int, error) {
		jsonBS, err := os.ReadFile(problemsetJsonFilePath)
		if err != nil {
			return 0, err
		}
		d := struct {
			Result struct {
				Problems []struct {
					ContestId int    `json:"contestId"`
					Index     string `json:"index"`
					Rating    int    `json:"rating"`
				} `json:"problems"`
			} `json:"result"`
		}{}
		if err := json.NewDecoder(bytes.NewReader(jsonBS)).Decode(&d); err != nil {
			return 0, err
		}
		for _, p := range d.Result.Problems {
			if strconv.Itoa(p.ContestId) == contestID && p.Index == problemID {
				return p.Rating, nil
			}
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
