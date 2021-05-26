package main

import (
	"encoding/json"
	"fmt"
)

//
//type jsonImageContent ImageContent
//
//type ImageContent struct {
//	Src  string `json:"src"`
//}
//
//func (c ImageContent) MarshalJSON() ([]byte, error) {
//	return json.Marshal(struct {
//		jsonImageContent
//		Type string `json:"type"`
//	}{
//		jsonImageContent: jsonImageContent(c),
//		Type:             "image",
//	})
//}
//
//type jsonLinkContent LinkContent
//
//type LinkContent struct {
//	Text string `json:"text"`
//	Link string `json:"link"`
//}
//
//func (c LinkContent) MarshalJSON() ([]byte, error) {
//	return json.Marshal(struct {
//		jsonLinkContent
//		Type string `json:"type"`
//	}{
//		jsonLinkContent: jsonLinkContent(c),
//		Type:            "link",
//	})
//}

type peo struct {
	Age    int    `json:"age"`
	MyName MyName `json:"name"`
}
type MyName struct {
	Name string
}

type a MyName

//
func (m MyName) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		a
		Other string `json:"other"`
	}{
		a:     a(m),
		Other: "111",
	})
}

func main() {
	p := peo{
		MyName: MyName{
			Name: "1",
		},
		Age: 18,
	}
	data, err := json.Marshal(p)
	fmt.Println(err)
	fmt.Printf("%s\n", data)
	//msg := other.Message{
	//	Title: "A apple",
	//	Content: []interface{}{
	//		other.TextContent{
	//			Text: "There is a apple.",
	//		},
	//		other.TextContent{
	//			Text: "It is red.",
	//		},
	//		//ImageContent{
	//		//	Src:  "http://example.com/apple.png",
	//		//},
	//		//LinkContent{
	//		//	Text: "more info",
	//		//	Link: "http://example.com/more.info.html",
	//		//},
	//	},
	//}
	//buffer, _ := json.Marshal(msg)
	//fmt.Printf("%s\n", buffer)
	// 输出：
	// {"title":"A apple","content":[{"text":"There is a apple.","type":"text"},{"text":"It is red.","type":"text"},{"src":"http://example.com/apple.png","type":"image"},{"text":"more info","link":"http://example.com/more.info.html","type":"link"}]}
}
