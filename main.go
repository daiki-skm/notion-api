package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	baseURL       = "https://api.notion.com/v1"
	apiVersion    = "2021-05-13"
)

type Key struct {
	ClientSecret string
	BlockId string
}

func main()  {
	key := &Key{
		os.Getenv("NOTION_KEY"),
		os.Getenv("BLOCK_ID"),
	}
	client := &http.Client{}
	get(client, key)
}

func get(client *http.Client, key *Key) {
	req, err := http.NewRequest("GET", baseURL+"/blocks/"+key.BlockId+"/children", nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", key.ClientSecret))
	req.Header.Set("Notion-Version", apiVersion)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
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
