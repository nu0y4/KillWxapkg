package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"

	. "github.com/Ackites/KillWxapkg/internal/cmd"
	. "github.com/Ackites/KillWxapkg/internal/config"
	"github.com/Ackites/KillWxapkg/internal/restore"
)

func Execute(appID, input, outputDir, fileExt string, restoreDir bool, pretty bool, noClean bool, save bool, sensitive bool) {
	// 存储配置
	configManager := NewSharedConfigManager()
	configManager.Set("appID", appID)
	configManager.Set("input", input)
	configManager.Set("outputDir", outputDir)
	configManager.Set("fileExt", fileExt)
	configManager.Set("restoreDir", restoreDir)
	configManager.Set("pretty", pretty)
	configManager.Set("noClean", noClean)
	configManager.Set("save", save)
	configManager.Set("sensitive", sensitive)

	inputFiles := ParseInput(input, fileExt)

	if len(inputFiles) == 0 {
		log.Println("未找到任何文件")
		return
	}

	// 确定输出目录
	if outputDir == "" {
		outputDir = DetermineOutputDir(input, appID)
	}

	var wg sync.WaitGroup
	for _, inputFile := range inputFiles {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			err := ProcessFile(file, outputDir, appID, save)
			if err != nil {
				log.Printf("处理文件 %s 时出错: %v\n", file, err)
			} else {
				log.Printf("成功处理文件: %s\n", file)
			}

		}(inputFile)
	}
	wg.Wait()

	// 还原工程目录结构
	restore.ProjectStructure(outputDir, restoreDir)
}

// 获取指定目录及子目录中文件的最后修改时间
func getDirectoryModTimes(path string) (map[string]time.Time, error) {
	modTimes := make(map[string]time.Time)

	err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		modTimes[p] = info.ModTime()
		return nil
	})

	if err != nil {
		return nil, err
	}

	return modTimes, nil
}

// 提取出wx小程序的id
func extractWXString(path string) (string, error) {
	parts := strings.Split(path, `\`)
	re := regexp.MustCompile(`^wx[a-z0-9]{16}$`)

	// 遍历分割后的路径部分，查找符合条件的字符串
	for _, part := range parts {
		if re.MatchString(part) {
			return part, nil
		}
	}

	// 如果没有找到，返回错误(悲)
	return "", fmt.Errorf("没有找到符合条件的 wx 字符串")
}

// 创建文件夹并返回完整路径
func createFolderWithTimestamp(baseName string) (string, error) {
	currentTime := time.Now()
	timeSuffix := currentTime.Format("2006.01.02.15.04")
	folderName := fmt.Sprintf("%s.%s", baseName, timeSuffix)
	fullPath := filepath.Join("output", folderName)

	err := os.MkdirAll(fullPath, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("无法创建文件夹: %v", err)
	}

	return fullPath, nil
}

// 实时监控
func WatchDirectory(dirToWatch string) {
	// 获取初始的目录修改时间列表
	previousModTimes, err := getDirectoryModTimes(dirToWatch)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for {
		// 每秒检测一次
		time.Sleep(1 * time.Second)

		// 获取当前的目录修改时间列表
		currentModTimes, err := getDirectoryModTimes(dirToWatch)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// 比较当前修改时间和之前的修改时间
		for path, _ := range currentModTimes {
			if _, exists := previousModTimes[path]; !exists {
				if strings.HasSuffix(path, ".wxapkg") {
					fmt.Println("检测到wx小程序文件:", path)
					appID, _ := extractWXString(path)
					fmt.Println("检测到ID：:", appID)
					fmt.Println("========================开始解包文件========================")
					pathName, err := createFolderWithTimestamp(appID)
					fmt.Println(pathName)
					if err != nil {
						return
					}
					Execute(appID, path, pathName, "", false, true, false, false, false)

				}
			}
		}

		// 更新之前的修改时间列表
		previousModTimes = currentModTimes
	}

}
