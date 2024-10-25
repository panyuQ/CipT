package main

import (
	"CipT/NoKey"
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Gui() {

	// 创建应用程序
	myApp := app.New()
	myWindow := myApp.NewWindow("Hello")

	// 创建一个标签
	label := widget.NewLabel("Hello World!")

	// 将标签放入容器中
	myWindow.SetContent(container.NewVBox(label))

	// 显示窗口并运行应用程序
	myWindow.ShowAndRun()
}

func Test(str string, method string) {
	x := NoKey.Cip{Before: []byte(str), Method: method}

	var err error

	fmt.Printf("[*] 原: %s\n%s\n", str, method)
	fmt.Print("[+] 加密后: ")
	if err = x.Encode(); err == nil {
		fmt.Println(string(x.After))
	} else {
		fmt.Println(err)
	}
	fmt.Print("[+] 解密后: ")
	if err = x.Decode(); err == nil {
		fmt.Println(string(x.Before))
	} else {
		fmt.Println(err)
	}
}

// 测试 XXEncode 编码和解码
func main() {
	Test("Hello", "Base64")
}
