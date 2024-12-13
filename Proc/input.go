package Proc

import (
	"CipT/Logger"
	"bufio"
	"os"
	"strings"
	"sync"
)

type Input struct {
	Text  []string
	File  []string
	cache map[string][]string // 缓存文件内容
	mu    sync.Mutex          // 用于保护缓存的并发安全
}

// NewInput 创建并初始化 Input 实例
func NewInput(in []string) *Input {
	if len(in) > 0 {
		var input Input
		input.cache = make(map[string][]string) // 初始化缓存
		for _, a := range in {
			_, err := os.Stat(a)
			if err != nil {
				if os.IsNotExist(err) {
					input.Text = append(input.Text, a)
				} else {
					Logger.GLogger.Error.Println(err)
				}
			} else {
				input.File = append(input.File, a)
			}
		}
		return &input
	}
	return &Input{}
}

// readFileWithCache 从缓存中读取文件内容，若无缓存则读取文件并存入缓存
func (input *Input) readFileWithCache(fileName string) []string {
	input.mu.Lock()
	defer input.mu.Unlock()

	// 如果文件内容已缓存，直接返回
	if content, ok := input.cache[fileName]; ok {
		return content
	}

	// 缓存未命中，读取文件内容
	var content []string
	file, err := os.Open(fileName)
	if err != nil {
		Logger.GLogger.Error.Printf("Failed to open file %s: %v", fileName, err)
		return nil
	}
	defer file.Close()

	// 读取前三个字节来检查是否存在UTF-8 BOM
	skipFileBOM(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = append(content, strings.TrimSpace(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		Logger.GLogger.Error.Printf("Error reading file %s: %v", fileName, err)
	}

	// 将读取的内容存入缓存
	input.cache[fileName] = content
	return content
}

// GetContent 获取分页内容（包括文本和文件内容）
func (input *Input) GetContent(pageIndex int, pageSize int) []string {

	// 存储所有内容
	var allContent []string

	// 添加 text 内容
	allContent = append(allContent, input.Text...)

	if pageSize <= 0 || pageIndex < 0 {
		return allContent // 无效的分页参数
	}

	// 添加文件内容（从缓存中获取）
	for _, fileName := range input.File {
		content := input.readFileWithCache(fileName)
		allContent = append(allContent, content...)
	}

	// 实现分页
	start := pageIndex * pageSize
	if start >= len(allContent) {
		return nil // 页码超出范围
	}
	end := start + pageSize
	if end > len(allContent) {
		end = len(allContent)
	}
	return allContent[start:end]
}

// GetContentCount 返回 Input 中所有内容的总行数
func (input *Input) GetContentCount(filename string) int {
	if input == nil {
		return 0 // 输入为空时返回 0
	}

	switch filename {
	case "":
		return len(input.Text)
	default:
		return len(input.readFileWithCache(filename))
	}
}

// GetAllContentCount 返回 Input 中所有内容的总行数
func (input *Input) GetAllContentCount() int {
	if input == nil {
		return 0 // 输入为空时返回 0
	}

	// 计算 text 中的行数
	totalCount := input.GetContentCount("")

	// 计算 file 中的行数
	for _, filename := range input.File {
		totalCount += input.GetContentCount(filename)
	}

	return totalCount
}

// skipFileBOM 读取前三个字节来检查是否存在 UTF-8 BOM, 并跳过
func skipFileBOM(file *os.File) {
	buf := make([]byte, 3)
	n, _ := file.Read(buf)

	// 检查是否为UTF-8 BOM
	if n >= 3 && buf[0] == 0xEF && buf[1] == 0xBB && buf[2] == 0xBF {
		// 跳过BOM
	} else {
		// 如果没有BOM，则将指针回退到文件开始处
		file.Seek(0, 0)
	}
}
