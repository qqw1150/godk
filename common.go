package main

import (
	"godk/str"
	"io/ioutil"
	"regexp"
	"strings"
)

//
//  LoadEnv 加载.env配置文件
//  @Description: 加载.env配置文件
//  @return map[string]string
//  @return error
//
func LoadEnv() (map[string]string, error) {
	bytes, err := ioutil.ReadFile("./.env")
	if err != nil {
		return nil, err
	}

	pattern, err := regexp.Compile("\n|\r\n")
	if err != nil {
		return nil, err
	}

	data := make(map[string]string, 0)
	res := pattern.Split(string(bytes), -1)
	for _, v := range res {
		v = strings.TrimSpace(v)

		length := 0
		if index := strings.Index(v, "#"); index == -1 {
			length = len(v)
		} else if index == 0 {
			continue
		} else {
			length = index
		}

		v = str.Substr(v, 0, length)

		kv := strings.Split(v, "=")
		k := strings.TrimSpace(kv[0])
		v := strings.TrimSpace(kv[1])
		data[k] = v
	}

	return data, nil
}
