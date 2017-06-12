package commonUtil

import (
	"net/http"
	"os"
	"log"
	"crypto/md5"
	"io"
	"encoding/hex"
	"strings"
	"fmt"
	"io/ioutil"
)

/*interface 转 string*/
func InterfaceToString(it interface{}) string {
	var str string
	var ok bool
	if str, ok = it.(string); !ok {
		str = ""
	}
	return str
}

/*byte 转 string*/
func ByteString(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}

/*判断是否空字符串*/
func IsEmpty(i interface{}) bool {
	var isEmpty bool
	switch i.(type) {
	case string:
		str := InterfaceToString(i)
		isEmpty = len(str) == 0
	}
	return isEmpty
}

/*获取 request参数,在调用此方法前要先 r.ParseForm */
func RequestGet(r *http.Request, paramName, defaultValue string) string {
	var returnValue string
	if len(r.Form[paramName]) == 0 {
		returnValue = defaultValue
	} else {
		returnValue = r.Form[paramName][0]
	}
	return returnValue
}

/*计算文件的md5*/
func MD5File(file_path *string) (string, error) {
	f, err := os.Open(*file_path)
	if err != nil {
		log.Println("Open", err)
		return "", err
	}

	defer f.Close()

	md5hash := md5.New()
	if _, err := io.Copy(md5hash, f); err != nil {
		log.Println("Copy", err)
		return "", err
	}

	md5 := md5hash.Sum(nil)
	return hex.EncodeToString(md5), nil
}

/*计算字符串的md5*/
func MD5(text *string) string {
	ctx := md5.New()
	ctx.Write([]byte(*text))
	return hex.EncodeToString(ctx.Sum(nil))
}

/*http 请求*/
func HttpPost(url string, params []string) ([]byte, error) {
	var param string
	if len(params) > 0 {
		param = strings.Join(params, "&")
	}

	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader(param))
	if err != nil {
		fmt.Println(err)
		return []byte{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err.Error())
		return []byte{}, err
	}
	return body, nil

}
