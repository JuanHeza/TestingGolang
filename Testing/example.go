package main

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {

	widgets.NewQApplication(len(os.Args), os.Args)

	//create a window
	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Hello World Example")
	window.SetMinimumSize2(200, 200)

	//create a layout
	layout := widgets.NewQVBoxLayout()

	//create a widget and set the layout
	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(layout)

	//create a lineedit and add it to the layout
	input := widgets.NewQLineEdit(nil)
	input.SetPlaceholderText("1. write something")
	layout.AddWidget(input, 0, 0)

	//create a button and add it to the layout
	button := widgets.NewQPushButton2("2. click me", nil)
	button.ConnectClicked(func(checked bool) {
		widgets.QMessageBox_Information(nil, "OK", input.Text(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
	})
	layout.AddWidget(button, 0, 0)

	//add the widget as the central widget to the window
	window.SetCentralWidget(widget)

	//show the window
	window.Show()

	//enter the main event loop
	widgets.QApplication_Exec()
}