package request

import (
	"fmt"
	"log"
	"net/http"
)

const (
	baseURL       = "https://api.notion.com/v1"
	apiVersion    = "2021-05-13"
)

type Key struct {
	ClientSecret string
	BlockId string
}

func Request(client *http.Client, key *Key) *http.Response {
	req, err := http.NewRequest("GET", baseURL+"/blocks/"+key.BlockId+"/children", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", key.ClientSecret))
	req.Header.Set("Notion-Version", apiVersion)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	//defer resp.Body.Close()

	return resp
}
