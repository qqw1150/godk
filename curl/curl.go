package curl

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//
//  UploadImage 上传图片
//  @Description: 上传图片 只支持jpg，png，gif,svg格式
//  @param stream 图片流
//  @param url 上传地址
//  @param fieldName 上传图片表单字段名
//  @param mimeType 上传图片mine类型
//
func UploadImage(stream, url, fieldName, mimeType string) {
	fType := strings.Split(mimeType, "/")[1]
	if !strings.Contains("jpg,jpeg,png,gif,svg", fType) {
		panic("文件只支持：jpg,jpeg,png,gif,svg")
	}

	bound := "curl"
	crlf := "\n"
	data := ""
	data += "--" + bound + crlf
	data += fmt.Sprintf(`Content-Disposition: form-data; name="%s"; filename="%s.%s"%s`, fieldName, fieldName, fType, crlf)
	data += fmt.Sprintf(`Content-Type: %s%s`, mimeType, crlf)
	data += crlf
	data += stream + crlf
	data += "--" + bound + "--" + crlf

	client := http.Client{}
	request, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		panic(err)
	}
	defer request.Body.Close()

	request.Header.Add("Content-Type", "multipart/form-data; boundary="+bound)

	_, err = client.Do(request)
	if err != nil {
		panic(err)
	}
}

//
//  Curl 发送http请求
//  @Description: 发送http请求
//  @param method 请求方法
//  @param url
//  @param query 请求参数 支持（string,map[string]string）
//  @param headers 请求头
//  @return []byte
//  @return error
//
func Curl(method, url string, query interface{}, headers map[string]string) ([]byte, error) {
	methodList := map[string]interface{}{"post": nil, "get": nil}
	s := strings.ToLower(method)
	if _, ok := methodList[s]; !ok {
		return nil, errors.New("请求方法必须是post/get")
	}

	client := http.Client{}

	data := ""
	if query != nil {
		if params, ok := query.(map[string]string); ok {
			i := 0
			for k, v := range params {
				if i > 0 {
					data += "&"
				}
				data += fmt.Sprintf("%s=%s", k, v)
				i++
			}
		} else {
			data, ok = query.(string)
			if !ok {
				return nil, errors.New("参数类型不合法")
			}
		}
	}

	request, err := http.NewRequest(strings.ToUpper(method), url, strings.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer request.Body.Close()

	hasUserAgent := false
	if headers != nil {
		for hName, hVal := range headers {
			if strings.ToLower(hName) == "user-agent" {
				hasUserAgent = true
			}
			request.Header.Add(hName, hVal)
		}
	}

	if !hasUserAgent {
		request.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36")
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return res, nil
}
