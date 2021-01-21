package webpushr

import (
	"fmt"
	"testing"
)

func Test_getConfig(t *testing.T) {
	err := GetConfig("./config/conf.yaml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(conf)
}

func Test_webpush(t *testing.T) {
	err := GetConfig("./config/conf.yaml")
	if err != nil {
		fmt.Println(err)
	}
	err = webpush()
	if err != nil {
		fmt.Println(err)
	}
}
