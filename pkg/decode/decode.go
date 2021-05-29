package decode

import (
	"encoding/json"
	"net/http"
	"time"
)

type NotionAPIStruct struct {
	Object  string `json:"object"`
	Results []struct {
		Object         string    `json:"object"`
		ID             string    `json:"id"`
		CreatedTime    time.Time `json:"created_time"`
		LastEditedTime time.Time `json:"last_edited_time"`
		HasChildren    bool      `json:"has_children"`
		Type           string    `json:"type"`
		Paragraph      struct {
			Text []struct {
				Type string `json:"type"`
				Text struct {
					Content string      `json:"content"`
					Link    interface{} `json:"link"`
				} `json:"text"`
				Annotations struct {
					Bold          bool   `json:"bold"`
					Italic        bool   `json:"italic"`
					Strikethrough bool   `json:"strikethrough"`
					Underline     bool   `json:"underline"`
					Code          bool   `json:"code"`
					Color         string `json:"color"`
				} `json:"annotations"`
				PlainText string      `json:"plain_text"`
				Href      interface{} `json:"href"`
			} `json:"text"`
		} `json:"paragraph,omitempty"`
		ToDo struct {
			Text []struct {
				Type string `json:"type"`
				Text struct {
					Content string      `json:"content"`
					Link    interface{} `json:"link"`
				} `json:"text"`
				Annotations struct {
					Bold          bool   `json:"bold"`
					Italic        bool   `json:"italic"`
					Strikethrough bool   `json:"strikethrough"`
					Underline     bool   `json:"underline"`
					Code          bool   `json:"code"`
					Color         string `json:"color"`
				} `json:"annotations"`
				PlainText string      `json:"plain_text"`
				Href      interface{} `json:"href"`
			} `json:"text"`
			Checked bool `json:"checked"`
		} `json:"to_do,omitempty"`
		BulletedListItem struct {
			Text []struct {
				Type string `json:"type"`
				Text struct {
					Content string      `json:"content"`
					Link    interface{} `json:"link"`
				} `json:"text"`
				Annotations struct {
					Bold          bool   `json:"bold"`
					Italic        bool   `json:"italic"`
					Strikethrough bool   `json:"strikethrough"`
					Underline     bool   `json:"underline"`
					Code          bool   `json:"code"`
					Color         string `json:"color"`
				} `json:"annotations"`
				PlainText string      `json:"plain_text"`
				Href      interface{} `json:"href"`
			} `json:"text"`
		} `json:"bulleted_list_item,omitempty"`
		Toggle struct {
			Text []struct {
				Type string `json:"type"`
				Text struct {
					Content string      `json:"content"`
					Link    interface{} `json:"link"`
				} `json:"text"`
				Annotations struct {
					Bold          bool   `json:"bold"`
					Italic        bool   `json:"italic"`
					Strikethrough bool   `json:"strikethrough"`
					Underline     bool   `json:"underline"`
					Code          bool   `json:"code"`
					Color         string `json:"color"`
				} `json:"annotations"`
				PlainText string      `json:"plain_text"`
				Href      interface{} `json:"href"`
			} `json:"text"`
		} `json:"toggle,omitempty"`
	} `json:"results"`
	NextCursor interface{} `json:"next_cursor"`
	HasMore    bool        `json:"has_more"`
}

func Decode(resp *http.Response, args []string) {
	defer resp.Body.Close()

	if args[0] != "paragraph" && args[0] != "todo" && args[0] != "bullet" && args[0] != "toggle" {
		println("Please check args.")
		return
	}

	apistruct := NotionAPIStruct{}

	if err := json.NewDecoder(resp.Body).Decode(&apistruct); err != nil {
		panic(err)
	}
	for _, hit := range apistruct.Results {
		switch hit.Type {
		case "paragraph":
			for _, hit2 := range hit.Paragraph.Text {
				if len(args) == 0 || args[0] == "paragraph" {
					println(hit2.PlainText)
				}
			}
		case "to_do":
			for _, hit2 := range hit.ToDo.Text {
				if len(args) == 0 || args[0] == "todo" {
					println(hit2.PlainText)
				}
			}
		case "bulleted_list_item":
			for _, hit2 := range hit.BulletedListItem.Text {
				if len(args) == 0 || args[0] == "bullet" {
					println(hit2.PlainText)
				}
			}
		case "toggle":
			for _, hit2 := range hit.Toggle.Text {
				if len(args) == 0 || args[0] == "toggle" {
					println(hit2.PlainText)
				}
			}
		}
	}
}
