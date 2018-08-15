package main

import "fmt"
import "time"

func main() {
	go heartBeat()
	time.Sleep(time.Second * 5)
}
func heartBeat(){
    for range time.Tick(time.Second *1){
        fmt.Println("Foo")
    }
}