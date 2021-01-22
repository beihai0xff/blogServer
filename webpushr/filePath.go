package webpushr

import (
	"fmt"
	"os"
)

// 判断目录或文件是否存在
func PathExists(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) { // 文件或目录不存在
		return fmt.Errorf("文件或目录：%s 不存在，", path)
	}
	return err
}
