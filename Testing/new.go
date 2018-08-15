package main

import (
	"fmt"
	"log"

	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

func main() {
	w, err := window.New(sciter.SW_TITLEBAR|sciter.SW_RESIZEABLE|sciter.SW_CONTROLS|sciter.SW_MAIN|sciter.SW_ENABLE_DEBUG, nil)
	if err != nil {
		log.Fatal("Create Window Error: ", err)
	}
	w.LoadFile("index.html")
	w.SetTitle("Callback")
	setEventHandler(w)
	w.Show()
	w.Run()
}
func setEventHandler(w *window.Window) {
	w.DefineFunction("getNetInformation", func(args ...*sciter.Value) *sciter.Value {
		if len(args)!=0{
			X:=fmt.Sprint(args[0]," Complete")
			log.Println(X)
			fn := args[1]
			fn.Invoke(sciter.NullValue(), "[Native Script]", sciter.NewValue(X))
		}else{
			log.Println("NO ARGS")
		}
		ret:= sciter.NewValue()
		ret.Set("ip", sciter.NewValue("127.0.0.1"))
		return ret
	})
}