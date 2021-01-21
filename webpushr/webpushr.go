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

func GetConfig(path string) error {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, &conf)
	return err
}

func webpush() error {
	client := &http.Client{}
	var data = strings.NewReader(`{"title":"beihai blog","message":"新文章发布","target_url":"https://www.wingsxdu.com"}`)
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
