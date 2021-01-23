package webpushr

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"gopkg.in/yaml.v3"
)

type webpushConfig struct {
	Key   string `yaml:"webpushrKey"`
	Token string `yaml:"webpushrAuthToken"`
}

var conf = webpushConfig{}

// 从 yaml 文件获取 webpushr 的授权验证信息
func GetConfig(path string) error {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	return err
}

// 发起一次 webpushr 推送
func webpush(info pageInfo) error {
	client := &http.Client{}
	var data = strings.NewReader(fmt.Sprintf(`{"title":"beihai blog","message":"%s","target_url":"%s"}`, info.title, info.url))
	req, err := http.NewRequest("POST", "https://api.webpushr.com/v1/notification/send/all", data)
	if err != nil {
		return err
	}
	req.Header.Set("webpushrKey", conf.Key)
	req.Header.Set("webpushrAuthToken", conf.Token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", bodyText)
	return nil
}
