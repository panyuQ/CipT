package Window

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Run() {

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
