package main

import (
	//"strconv"
	"fmt"
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/math"
	"github.com/google/gxui/samples/flags"
)

		
// Number picker uses the gxui.DefaultAdapter for driving a list
func GenCogMix(theme gxui.Theme, overlay gxui.BubbleOverlay) gxui.Control {

	layout := theme.CreateRelativeLayout()
	layout.SetDirection(gxui.TopToBottom)

	labelX := theme.CreateLabel()
	labelX.SetText("Semilla Inicial \"Xo\":")
	textoX := theme.CreateTextBox()
	textoX.Text()
	layout.AddChild(labelX)
	layout.AddChild(textoX)

	labelA := theme.CreateLabel()
	labelA.SetText("Constante Multiplicativa \"A\":")
	textoA := theme.CreateTextBox()
	textoA.Text()
	layout.AddChild(labelA)
	layout.AddChild(textoA)

	labelC := theme.CreateLabel()
	labelC.SetText("Constante Aditiva \"C\":")
	textoC := theme.CreateTextBox()
	textoC.Text()
	layout.AddChild(labelC)
	layout.AddChild(textoC)

	labelM := theme.CreateLabel()
	labelM.SetText("Modulo \"M\":")
	textoM := theme.CreateTextBox()
	textoM.Text()
	
	layout.AddChild(labelM)
	layout.AddChild(textoM)
	
	//labelQ := theme.CreateLabel()
	adapter := gxui.CreateDefaultAdapter()
	list := theme.CreateList()
	var Res[] string
	button := func(name string, action func()) gxui.Button {
		b := theme.CreateButton()
		b.SetText(name)
		b.OnClick(func(gxui.MouseEvent) { action()})
		layout.AddChild(b)
		return b
	}

	Final := theme.CreateLabel()
	button("Aceptar",
		func() { Res=Mixto(textoA.Text(),textoX.Text(),textoC.Text(),textoM.Text())
		Final.SetText(Res[len(Res)-1])
		Res=Res[:len(Res)-1]
		/*labelQ.SetMultiline(false)
		labelQ.SetText("N\tXO\t(AXo+C) mod M\tXn+1\tRECTANGULAR\n")
		labelQ.SetMultiline(true)
			if len(Res)>0{
				for i:=0;i<len(Res)-1;i++{
					//labelQ.SetText(labelQ.Text()+Res[i])
				}
				
			}*/
			adapter.SetItems(Res)
			list.SetAdapter(adapter)
			list.SetOrientation(gxui.Vertical)
		},
	)
	layout.AddChild(Final)
	fmt.Println(Final.Color())
	
	resultado := theme.CreateScrollLayout ()
	resultado.SetScrollAxis(false,true)
	//resultado.SetChild(labelQ)
	resultado.SetChild(list)
	table := theme.CreateTableLayout()
	table.SetGrid(3, 2) // columns, rows
	table.SetChildAt(0, 0, 1, 1, layout)
	table.SetChildAt(1, 0, 1, 2, resultado)
	return table
}


// Color picker uses the customAdapter for driving a list
func colorPicker(theme gxui.Theme) gxui.Control {
	layout := theme.CreateLinearLayout()
	layout.SetDirection(gxui.TopToBottom)
	
	
	return layout
}
func appMain(driver gxui.Driver) {
	theme := flags.CreateTheme(driver)

	overlay := theme.CreateBubbleOverlay()

	holder := theme.CreatePanelHolder()
	holder.AddPanel(GenCogMix(theme, overlay), "Generador Mixto")
	holder.AddPanel(colorPicker(theme), "Custom adapter")

	window := theme.CreateWindow(800, 600, "Lists")
	window.SetScale(flags.DefaultScaleFactor)
	window.AddChild(holder)
	window.AddChild(overlay)
	window.OnClose(driver.Terminate)
	window.SetPadding(math.Spacing{L: 10, T: 10, R: 10, B: 10})
}

func main() {
	gl.StartDriver(appMain)
}