package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.jianshu.com/p/2f0b4659e598")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(data))
	if err != nil {
		fmt.Println(err)
	}
	// 标题
	title := doc.Find("._1RuRku").Text()
	fmt.Println("title", title)
	// 正文
	body := doc.Find("._2rhmJa").Children()
	body.Each(func(i int, selection *goquery.Selection) {
		if selection.Is("p") {
			fmt.Println("p", selection.Text())
		} else {
			imgSelector := selection.Find("img")
			if imgSelector.Length() > 0 {
				fmt.Println("111")
				if v, ex := imgSelector.Attr("data-original-src"); ex {
					fmt.Println("img", fmt.Sprintf("https:%s", v))
				}
			}
		}
	})
}
