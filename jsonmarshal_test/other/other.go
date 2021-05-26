package other

import "encoding/json"

type Message struct {
	Title   string        `json:"title"`
	Content []interface{} `json:"content"`
}

type jsonTextContent TextContent

type TextContent struct {
	Text string `json:"text"`
}

func (c TextContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		jsonTextContent
		Type string `json:"type"`
	}{
		jsonTextContent: jsonTextContent(c),
		Type:            "text",
	})
}
