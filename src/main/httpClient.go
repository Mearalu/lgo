package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"io"
)


/**
url 指定url
data post数据 "name=cjb"
*/

func httpPost(httpUrl string, data string) {
	resp, err := http.Post(httpUrl,
		"application/x-www-form-urlencoded",
		strings.NewReader(data))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

/**

 */
func httpPostForm(httpUrl string) {
	resp, err := http.PostForm(httpUrl,
		url.Values{"key": {"Value"}, "id": {"123"}})

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

}

func httpget(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	if resp.StatusCode == http.StatusOK {
		fmt.Println(resp.StatusCode)
	}
	defer resp.Body.Close()

	buf := make([]byte, 1024)
	f, err1 := os.OpenFile("path.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, os.ModePerm) //可读写，追加的方式打开（或创建文件）
	if err1 != nil {
		panic(err1)
		return
	}
	defer f.Close()

	for {
		n, _ := resp.Body.Read(buf)
		if 0 == n {
			break
		}
		//f.WriteString(string(buf[:n]))
		fmt.Println(string(buf[:n]))
	}
}
func main() {
	//httpget("http://www.lixinedu.com.cn/")

	//headers := map[string]string{"Cookie": "name=anny"}
	//httpDo("http://www.lixinedu.com.cn/", "", headers, "GET")

	//httpDo("http://www.oschina.net/", "", headers, "GET")
	//
	//d,_:=Decode([]byte(src))
	//fmt.Println(d)
	//fmt.Println(string(d))
	//LiXin("http://www.lixinedu.com.cn")
	Scrape()

	//regexTest()
}
func mahoniaTest() {
	src := "编码转换内容内容"

	enc := mahonia.NewEncoder("GBK")
	output := enc.ConvertString(src)
	fmt.Println(output)
}

func LiXin(httpurl string) {
	doc, err := goquery.NewDocument(httpurl)
	if err != nil {
		fmt.Println(err)
		return
	}
	doc.Find("div.block1_2_arc2").Each(func(i int, s *goquery.Selection) {
		//band := s.Find("li").Text()
		title := s.Find("a").Text()
		d := []byte(title)
		fmt.Printf("Review %d: - %s\n", i, d)
	})
}

func Scrape() {
	doc, err := goquery.NewDocumentFromReader(httpDo("http://www.lixinedu.com.cn/", "", nil, "GET"))
	if err != nil {
		fmt.Println(err)
		return
	}
	doc.Find("div.block1_2_arc2").Each(func(i int, s *goquery.Selection) {
		title := s.Find("a").Text()
		fmt.Printf("Review %d: - %s\n", i,  title)
	})
}
/**

 */
func httpDo(httpUrl string, data string, headers map[string]string, method string) (io.Reader) {
	client := &http.Client{}
	reg, _ := regexp.Compile(`<meta.+?charset=[^\w]?([-\w]+)`)
	charset := "utf-8"
	req, err := http.NewRequest(method, httpUrl, strings.NewReader(data))
	if err != nil {
		// handle error
	}
	if strings.EqualFold("POST", method) {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	//req.Header.Set("Cookie", "name=anny")

	resp, err := client.Do(req)
	rs := reg.FindSubmatch([]byte(resp.Header.Get("Content-Type")))
	if len(rs) > 0 {
		charset = string(rs[1])
	}
	defer resp.Body.Close()


	//缓冲读取
	//buf := make([]byte, 1024)
	//n, _ := resp.Body.Read(buf)
	//rs := reg.FindSubmatch(buf)
	//if len(rs) > 0 {
	//	charset = string(rs[1])
	//}
	////
	//fmt.Println("|", string(ToUTF8Byte(buf[:n],charset)), "|")
	//for {
	//	n, _ := resp.Body.Read(buf)
	//	if 0 == n {
	//		break
	//	}
	//
	//	//f.WriteString(string(buf[:n]))
	//	//fmt.Println(string(Decode(string(buf[:n]))))
	//	fmt.Println(string(GBKToUTF8Byte(buf[:n],charset)))
	//}


	//直接读取
	body, err := ioutil.ReadAll(resp.Body)
	rs = reg.FindSubmatch(body)
	if len(rs) > 0 {
		charset = string(rs[1])
	}
	rsss := ToUTF8Byte(body, charset)
	return bytes.NewReader(rsss);
}
func Decode(s []byte) (rs []byte) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		fmt.Println(e)
		return s
	}
	return d
}

func ToUTF8Byte(s []byte, charset string) ([]byte) {
	I := bytes.NewReader(s)
	var O io.Reader
	switch charset {
	case "gb2312", "GB2312":
		O = transform.NewReader(I, simplifiedchinese.HZGB2312.NewDecoder())
	case "gbk", "GBK":
		O = transform.NewReader(I, simplifiedchinese.GBK.NewDecoder())
	case "gb18030", "GB18030":
		O = transform.NewReader(I, simplifiedchinese.GB18030.NewDecoder())
	default:
		return s
	}
	d, e := ioutil.ReadAll(O)
	if e != nil {
		fmt.Println(e)
		return s
	}
	return d
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
