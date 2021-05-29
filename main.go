package main

import (
	"flag"
	"net/http"
	"os"

	"blog-api/pkg/decode"
	"blog-api/pkg/request"
)

func main()  {
	key := &request.Key{
		ClientSecret: os.Getenv("NOTION_KEY"),
		BlockId:      os.Getenv("BLOCK_ID"),
	}

	flag.Parse()
	args := flag.Args()

	client := &http.Client{}
	resp := request.Request(client, key)
	decode.Decode(resp, args)
}
