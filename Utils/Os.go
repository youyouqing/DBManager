package Utils

import (
	"os"
	"path/filepath"
	"time"
)

func CreateDateDir(dirPaths string, mode os.FileMode) string {
	folderName := time.Now().Format("20060102")
	folderPath := filepath.Join(dirPaths, folderName)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		_ = os.MkdirAll(folderPath, mode) //0777也可以os.ModePerm
		_ = os.Chmod(folderPath, mode)
	}
	return folderPath
}
