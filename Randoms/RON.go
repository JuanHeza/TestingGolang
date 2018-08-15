package main

import (
	//"strconv"
	"fmt"
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/math"
	"github.com/google/gxui/samples/flags"
)

///*	Generador Congruencial Mixto
func GenCogMix(theme gxui.Theme, overlay gxui.BubbleOverlay) gxui.Control {

	layout := theme.CreateLinearLayout()
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
//	adapter := gxui.CreateDefaultAdapter()
	list := theme.CreateList()
//	var Res[] string
	/*button := func(name string, action func()) gxui.Button {
		b := theme.CreateButton()
		b.SetText(name)
		b.OnClick(func(gxui.MouseEvent) { action()})
		layout.AddChild(b)
		return b
	}*/

	Final := theme.CreateLabel()
	/*button("Aceptar",
		func() { Res=Mixto(textoA.Text(),textoX.Text(),textoC.Text(),textoM.Text())
		Final.SetText(Res[len(Res)-1])
		Res=Res[:len(Res)-1]
			adapter.SetItems(Res)
			list.SetAdapter(adapter)
			list.SetOrientation(gxui.Vertical)
		},
	)*/
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
//*/

/* Generador Congruencial Multiplicativo
func GenCogMul(theme gxui.Theme, overlay gxui.BubbleOverlay) gxui.Control {

	layout := theme.CreateLinearLayout()
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
		func() { Res=Multi(textoA.Text(),textoX.Text(),textoM.Text())
		Final.SetText(Res[len(Res)-1])
		Res=Res[:len(Res)-1]
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
*/

/* Kolmogorov
func GenCogMul(theme gxui.Theme, overlay gxui.BubbleOverlay) gxui.Control {

	layout := theme.CreateLinearLayout()
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
		func() { Res=Multi(textoA.Text(),textoX.Text(),textoM.Text())
		Final.SetText(Res[len(Res)-1])
		Res=Res[:len(Res)-1]
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
*/

/* Frecuencias
func GenCogMul(theme gxui.Theme, overlay gxui.BubbleOverlay) gxui.Control {

	layout := theme.CreateLinearLayout()
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
		func() { Res=Multi(textoA.Text(),textoX.Text(),textoM.Text())
		Final.SetText(Res[len(Res)-1])
		Res=Res[:len(Res)-1]
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
*/

/* Promedios
func GenCogMul(theme gxui.Theme, overlay gxui.BubbleOverlay) gxui.Control {

	layout := theme.CreateLinearLayout()
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
		func() { Res=Multi(textoA.Text(),textoX.Text(),textoM.Text())
		Final.SetText(Res[len(Res)-1])
		Res=Res[:len(Res)-1]
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
*/

/* Poisson
func GenCogMul(theme gxui.Theme, overlay gxui.BubbleOverlay) gxui.Control {

	layout := theme.CreateLinearLayout()
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
		func() { Res=Multi(textoA.Text(),textoX.Text(),textoM.Text())
		Final.SetText(Res[len(Res)-1])
		Res=Res[:len(Res)-1]
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
*/

/* Uniforme
func GenCogMul(theme gxui.Theme, overlay gxui.BubbleOverlay) gxui.Control {

	layout := theme.CreateLinearLayout()
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
		func() { Res=Multi(textoA.Text(),textoX.Text(),textoM.Text())
		Final.SetText(Res[len(Res)-1])
		Res=Res[:len(Res)-1]
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
*/

/* Exponencial
func GenCogMul(theme gxui.Theme, overlay gxui.BubbleOverlay) gxui.Control {

	layout := theme.CreateLinearLayout()
	layout.SetDirection(gxui.TopToBottom)

	labelX := theme.CreateLabel()
	labelX.SetText("Numeros \"N\":")
	textoX := theme.CreateTextBox()
	textoX.Text()
	layout.AddChild(labelX)
	layout.AddChild(textoX)

	labelM := theme.CreateLabel()
	labelM.SetText("Lamda:")
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
		func() { Res=Multi(textoA.Text(),textoX.Text(),textoM.Text())
		Final.SetText(Res[len(Res)-1])
		Res=Res[:len(Res)-1]
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
*/

/* Volados
func GenCogMul(theme gxui.Theme, overlay gxui.BubbleOverlay) gxui.Control {

	layout := theme.CreateLinearLayout()
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
		func() { Res=Multi(textoA.Text(),textoX.Text(),textoM.Text())
		Final.SetText(Res[len(Res)-1])
		Res=Res[:len(Res)-1]
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
*/

/* Colas
func GenCogMul(theme gxui.Theme, overlay gxui.BubbleOverlay) gxui.Control {

	layout := theme.CreateLinearLayout()
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
		func() { Res=Multi(textoA.Text(),textoX.Text(),textoM.Text())
		Final.SetText(Res[len(Res)-1])
		Res=Res[:len(Res)-1]
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
*/

/*
*/
func appMain(driver gxui.Driver) {
	theme := flags.CreateTheme(driver)

	overlay := theme.CreateBubbleOverlay()

	holder := theme.CreatePanelHolder()
//	holder.AddPanel(GenCogMix(theme, overlay), "Generador Mixto") 			//:D
//	holder.AddPanel(GenCogMul(theme, overlay), "Generador Multiplicativo")	//:D
//	holder.AddPanel(GenCogMul(theme, overlay), "Generador")					//:D
//	holder.AddPanel(GenCogMul(theme, overlay), "Kolmogorov")				//:(
//	holder.AddPanel(GenCogMul(theme, overlay), "Frecuencias")				//:(
//	holder.AddPanel(GenCogMul(theme, overlay), "Promedios")					//:(
//	holder.AddPanel(GenCogMul(theme, overlay), "Poisson")					//:(
//	holder.AddPanel(GenCogMul(theme, overlay), "Uniforme")					//:D
	holder.AddPanel(GenCogMix(theme, overlay), "Exponencial")				//:D
//	holder.AddPanel(GenCogMul(theme, overlay), "Volados")					//:D
//	holder.AddPanel(GenCogMul(theme, overlay), "Colas")						//:(

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