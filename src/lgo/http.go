package lgo

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"lgo/encode"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

var (
	metaCharsetReg, _ = regexp.Compile(`<meta.+?charset=[^\w]?([-\w]+)`)
	charsetReg, _     = regexp.Compile(`charset=[^\w]?([-\w]+)`)
)

func HttpDoString(httpUrl string, data string, headers map[string]string, method string) string {
	r := HttpDo(httpUrl, data, headers, method)
	d, e := ioutil.ReadAll(r)
	if e != nil {
		fmt.Println(e)
		return ""
	}
	return string(d)
}

func HttpGet(httpUrl string) (res *http.Response) {
	return HttpResp(httpUrl, "", nil, "GET")
}
func httpDo(httpUrl string, data string, headers map[string]string, method string) (res *http.Response, err error) {
	client := &http.Client{}
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

	return client.Do(req)
}

func HttpResp(httpUrl string, data string, headers map[string]string, method string) (res *http.Response) {
	charset := "utf-8"
	resp, err := httpDo(httpUrl, data, headers, method)
	if err != nil {
		//
	}
	cs := charsetReg.FindSubmatch([]byte(resp.Header.Get("Content-Type")))
	if len(cs) > 0 {
		charset = string(cs[1])
		resp.Body = ioutil.NopCloser(encode.ToUTF8Reader(resp.Body, charset))
		return resp
	}
	defer resp.Body.Close()
	////缓冲读取
	//buf := make([]byte, 1024)
	//for {
	//	n, _ := resp.Body.Read(buf)
	//	if 0 == n {
	//		break
	//	}
	//	//f.WriteString(string(buf[:n]))
	//	fmt.Println(string(encode.ToUTF8Byte(buf[:n],charset)))
	//}
	//直接读取
	body, err := ioutil.ReadAll(resp.Body)
	cs = metaCharsetReg.FindSubmatch(body)
	if len(cs) > 0 {
		charset = string(cs[1])
	}
	rs := encode.ToUTF8Byte(body, charset)
	resp.Body = ioutil.NopCloser(bytes.NewReader(rs))
	return resp
}
func HttpDo(httpUrl string, data string, headers map[string]string, method string) io.Reader {
	charset := "utf-8"
	resp, err := httpDo(httpUrl, data, headers, method)
	if err != nil {
		//
	}
	cs := charsetReg.FindSubmatch([]byte(resp.Header.Get("Content-Type")))
	if len(cs) > 0 {
		charset = string(cs[1])
		return encode.ToUTF8Reader(resp.Body, charset)
	}
	defer resp.Body.Close()
	//直接读取
	body, err := ioutil.ReadAll(resp.Body)
	cs = metaCharsetReg.FindSubmatch(body)
	if len(cs) > 0 {
		charset = string(cs[1])
	}
	rs := encode.ToUTF8Byte(body, charset)
	return bytes.NewReader(rs)
}

func HttpPostForm(httpUrl string) {
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
	f, err1 := os.OpenFile("path.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm) //可读写，追加的方式打开（或创建文件）
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
