package Proc

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

func TestInput(t *testing.T) {
	// 创建临时文件
	file1 := "test1.txt"
	file2 := "test2.txt"

	// 写入测试内容
	createTestFile(t, file1, "Line1\nLine2\nLine3\nLine4\n")
	createTestFile(t, file2, "A\nB\nC\nD\nE\n")
	defer os.Remove(file1)
	defer os.Remove(file2)

	// 模拟命令行参数
	os.Args = []string{"program", file1, "nonexistent", file2}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flag.Parse()

	// 初始化 Input
	input := NewInput()
	if input == nil {
		t.Fatal("Expected a valid Input object, got nil")
	}

	// 测试分页获取
	content := input.GetContent(2, 1)
	expected := []string{"nonexistent", "Line1"}
	if !compareSlices(content, expected) {
		t.Errorf("Expected %v, got %v", expected, content)
	}

	content = input.GetContent(3, 2)
	expected = []string{"Line3", "Line4", "A"}
	if !compareSlices(content, expected) {
		t.Errorf("Expected %v, got %v", expected, content)
	}

	// 测试缓存性能（重复读取文件）
	for i := 0; i < 10; i++ {
		_ = input.GetContent(2, 1) // 触发缓存读取
	}

	// 获取内容总行数
	totalCount := input.GetAllContentCount()
	expectedCount := 10 // 1 texts + 4 lines (file1) + 5 lines (file2)
	if totalCount != expectedCount {
		t.Errorf("Expected total count %d, got %d", expectedCount, totalCount)
	}

	fmt.Println("All tests passed!")
}

// 创建测试文件
func createTestFile(t *testing.T, filename, content string) {
	file, err := os.Create(filename)
	if err != nil {
		t.Fatalf("Failed to create test file %s: %v", filename, err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		t.Fatalf("Failed to write content to file %s: %v", filename, err)
	}
}

// 比较切片内容
func compareSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
