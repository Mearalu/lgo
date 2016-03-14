package encode

import (
	"io"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
	"bytes"
	"io/ioutil"
	"fmt"
)

func ToUTF8Reader(r io.Reader, charset string)(io.Reader){
	switch charset {
	case "gb2312", "GB2312":
		return transform.NewReader(r, simplifiedchinese.HZGB2312.NewDecoder())
	case "gbk", "GBK":
		return transform.NewReader(r, simplifiedchinese.GBK.NewDecoder())
	case "gb18030", "GB18030":
		return transform.NewReader(r, simplifiedchinese.GB18030.NewDecoder())
	default:
		return r
	}
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
