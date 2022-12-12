package utils

import (
	"os"
	"path/filepath"
)

// GetFiles 获取文件路径集合
// workdir 基础路径
// fuzzyPath 要赛选的路径
// pathList 查找到的路径对象
func GetFiles(workdir string, fuzzyPath []string, pathList []string) []string {
	files, _ := os.ReadDir(workdir)
	flag := false
	for _, file := range files {
		currentPath := workdir + "/" + file.Name()
		for _, path := range fuzzyPath {
			matched, err := filepath.Match(path, currentPath)
			flag = matched
			if matched && err == nil {
				pathList = append(pathList, currentPath)
			}
		}
		if file.IsDir() && !flag {
			pathList = GetFiles(currentPath, fuzzyPath, pathList)
		}
	}
	return pathList
}

// GetSuffixFiles 获取路径下的所有相同后缀文件
// workdir 基础路径
// suffixName 要赛选的路径
// pathList 查找到的路径对象
func GetSuffixFiles(workdir string, suffixName string, pathList []string) []string {
	files, _ := os.ReadDir(workdir)
	for _, file := range files {
		currentPath := workdir + "/" + file.Name()
		if file.IsDir() {
			pathList = GetSuffixFiles(currentPath, suffixName, pathList)
		} else {
			if filepath.Ext(currentPath) == suffixName {
				pathList = append(pathList, currentPath)
			}
		}
	}
	return pathList
}
