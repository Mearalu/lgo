package lhttp

import (
	"net/http"
	"net/url"
	"io/ioutil"
	"fmt"
	"regexp"
	"io"
	"strings"
	"lencoding"
	"bytes"
)

func LhttpDo(httpUrl string, data string, headers map[string]string, method string) (io.Reader) {
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
	cs := reg.FindSubmatch([]byte(resp.Header.Get("Content-Type")))
	if len(cs) > 0 {
		charset = string(cs[1])
		return lencoding.ToUTF8Reader(resp.Body,charset)
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
	cs = reg.FindSubmatch(body)
	if len(cs) > 0 {
		charset = string(cs[1])
	}
	rs:= lencoding.ToUTF8Byte(body, charset)
	return bytes.NewReader(rs);
}


func LhttpPostForm(httpUrl string) {
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