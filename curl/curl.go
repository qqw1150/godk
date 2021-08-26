package curl

import (
	"fmt"
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
