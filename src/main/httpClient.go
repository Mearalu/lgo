package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"lgo"
	"regexp"
)

func main() {
	//httpget("http://www.lixinedu.com.cn/")

	//headers := map[string]string{"Cookie": "name=anny"}
	//r:=lgo.HttpDo("http://www.lixinedu.com.cn/", "", headers, "GET")

	//r:=lgo.HttpDoString("http://www.oschina.net/", "", headers, "GET")
	//fmt.Println(r)
	//
	//d,_:=Decode([]byte(src))
	//fmt.Println(d)
	//fmt.Println(string(d))
	//LiXin("http://www.lixinedu.com.cn")

	Scrape("http://www.oschina.net/")
	LiXin()

	//regexTest()
}
func mahoniaTest() {
	src := "编码转换内容内容"

	enc := mahonia.NewEncoder("GBK")
	output := enc.ConvertString(src)
	fmt.Println(output)
}

func Scrape(httpUrl string) {
	//doc, err := goquery.NewDocumentFromReader(lgo.HttpDo(httpUrl, "", nil, "GET"))
	doc, err := goquery.NewDocumentFromResponse(lgo.HttpResp(httpUrl, "", nil, "GET"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(doc.Url)
	doc.Find("li.today  a").Each(func(i int, s *goquery.Selection) {
		//band := s.Find("li").Text()
		title := s.Text()
		url, _ := s.Attr("href")
		d := []byte(title)
		fmt.Printf("Review %d: url%s  - %s\n", i, url, d)
	})
}

func LiXin() {
	doc, err := goquery.NewDocumentFromResponse(lgo.HttpGet("http://www.lixinedu.com.cn/"))
	if err != nil {
		fmt.Println(err)
		return
	}
	doc.Find("div.b3_2_2 li").Each(func(i int, s *goquery.Selection) {
		title := s.Find("a").Text()
		fmt.Printf("Review %d: - %s\n", i, title)
	})
}

func regexTest() {
	str := ` <!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">

<html xmlns="http://www.w3.org/1999/xhtml">
<head>


<meta http-equiv="X-UA-Compatible" content="IE=edge">`

	fmt.Println(regexp.Match(`<meta.+?charset=[^\w]?([-\w]+)`, []byte(str)))

	reg, _ := regexp.Compile(`<meta.+?charset=[^\w]?([-\w]+)`)
	fmt.Println(reg.FindString(str))
	rs := reg.FindSubmatch([]byte(str))
	fmt.Println(len(rs))
	fmt.Println("FindAllSubmatch", reg.FindAllSubmatch([]byte(str), -1))
	fmt.Println(reg.ReplaceAllString(str, "$1"))
}
