package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	baseURL       = "https://api.notion.com/v1"
	apiVersion    = "2021-05-13"
)

//func (c *Client) newRequest(ctx context.Context, method, url string, body io.Reader) (*http.Request, error) {
//	req, err := http.NewRequestWithContext(ctx, method, baseURL+url, body)
//	if err != nil {
//		return nil, err
//	}
//
//	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", c.apiKey))
//	req.Header.Set("Notion-Version", apiVersion)
//	//req.Header.Set("User-Agent", "go-notion/"+clientVersion)
//
//	if body != nil {
//		req.Header.Set("Content-Type", "application/json")
//	}
//
//	return req, nil
//}

var (
	clientSecret string
	blockId string
)

func main()  {
	clientSecret = os.Getenv("NOTION_KEY")
	blockId = os.Getenv("BLOCK_ID")
	client := &http.Client{}
	get(client)
}

func get(client *http.Client) {
	req, err := http.NewRequest("GET", baseURL+"/blocks/"+blockId+"/children", nil)
	if err != nil {
		println("error")
		return
	}

	// User-Agentを設定
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", clientSecret))

	resp, err := client.Do(req)
	if err != nil {
		println("error")
		return
	}
	defer resp.Body.Close()

	execute(resp)
}

func execute(resp *http.Response) {
	// response bodyを文字列で取得するサンプル
	// ioutil.ReadAllを使う
	b, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		println(string(b))
	}
}
