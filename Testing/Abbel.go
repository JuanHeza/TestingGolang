package main

import (
	"github.com/kyokomi/emoji"
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/samples/flags"
)

func appMain(driver gxui.Driver) {
	theme := flags.CreateTheme(driver)

	label1 := theme.CreateLabel()
	label1.SetColor(gxui.White)
	label1.SetText("Config")

	Config := theme.CreateLinearLayout()
	Config.SetBackgroundBrush(gxui.CreateBrush(gxui.Blue40))
	Config.SetHorizontalAlignment(gxui.AlignCenter)
	Config.AddChild(label1)

	label2 := theme.CreateLabel()
	label2.SetColor(gxui.White)
	label2.SetText("Chat")

	Chat := theme.CreateLinearLayout()
	Chat.SetBackgroundBrush(gxui.CreateBrush(gxui.Green40))
	Chat.SetHorizontalAlignment(gxui.AlignCenter)
	Chat.AddChild(label2)

	label3 := theme.CreateLabel()
	label3.SetColor(gxui.White)
	label3.SetText("Users")

	Users := theme.CreateLinearLayout()
	Users.SetBackgroundBrush(gxui.CreateBrush(gxui.Red40))
	Users.SetHorizontalAlignment(gxui.AlignCenter)
	Users.AddChild(label3)
	
	label4 := theme.CreateLabel()
	label4.SetColor(gxui.White)
	x := emoji.Sprint("I like a :pizza: and :sushi:!!")
	label4.SetText(x)
	
	textbutton := theme.CreateTableLayout()
	textbutton.SetGrid(4,1)
	textbox := theme.CreateTextBox()
	textbox.SetDesiredWidth(15) //65
	textbox.SetMultiline(true)
	Button := theme.CreateButton()
	Button.SetText("Enviar")
	textbutton.SetChildAt(0,0,3,1, textbox)
	textbutton.SetChildAt(3,0,1,1, Button)
	
	Textbox := theme.CreateLinearLayout()
	Textbox.SetBackgroundBrush(gxui.CreateBrush(gxui.Red40))
	Textbox.SetHorizontalAlignment(gxui.AlignCenter)
	Textbox.AddChild(label4)
	Textbox.AddChild(textbutton)

	table := theme.CreateTableLayout()
	table.SetGrid(6, 6) // columns, rows

	// row, column, horizontal span, vertical span
	table.SetChildAt(0, 0, 2, 5, Users)
	table.SetChildAt(0, 5, 2, 1, Config)
	table.SetChildAt(2, 5, 4, 1, Textbox)
	table.SetChildAt(2, 0, 4, 5, Chat)

	window := theme.CreateWindow(800, 600, "Table")
	window.AddChild(table)
	window.OnClose(driver.Terminate)
}

func main() {
	gl.StartDriver(appMain)
}