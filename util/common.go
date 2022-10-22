package util

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// LoopNoOps 阻止进程退出
func LoopNoOps() {
	select {}
}
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	//将\替换成/
	return strings.Replace(dir, "\\", "/", -1)
}
