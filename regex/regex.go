package regex

import (
	"fmt"
	"regexp"
)

type Regex struct {
	expr string
}

//func Compile(expr string) (*Regex, error) {
//	re, err := regexp.Compile(expr)
//	if err != nil {
//		return nil, err
//	}
//	return &Regex{expr: expr}, nil
//}
//
//func MustCompile(expr string) *Regex {
//	re := regexp.MustCompile(expr)
//	return &Regex{expr: expr}
//}

func main() {
	// 使用 Compile 编译正则表达式
	re, err := regexp.Compile(`\d+`) // 匹配数字
	if err != nil {
		fmt.Println("正则表达式编译失败:", err)
		return
	}

	// 使用 MustCompile 编译正则表达式
	re2 := regexp.MustCompile(`[a-z]+`) // 匹配小写字母

	fmt.Println(re.FindString("hello world! 123456789"))
	fmt.Println(re2.FindString("hello123")) // 输出: hello
}
