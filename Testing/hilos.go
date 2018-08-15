package main

import (
    "os"
	"fmt"
    "bufio"
)

// var x = 1
type Data struct{
	n int
	st bool
}
func main(){
	x := make(chan int,5)
		x <- 1 //1
		x <- 2 //2
		x <- 3 //6
		x <- 4 //24
		x <- 5 //120
	ch := make(chan Data,5)
		go Factorial(<-x,ch,false)
		go Factorial(<-x,ch,false)
		   Factorial(<-x,ch,true)
		go Factorial(<-x,ch,false)
		   Factorial(<-x,ch,true)
	for i:=1;i<=5;i++{
		fmt.Println("\t",<-ch)
	}
	
  fmt.Println("Press 'Enter' to continue...")
  bufio.NewReader(os.Stdin).ReadBytes('\n') 
}

func Factorial(I int,ch chan Data,status bool){
	Fact:=1
	for i:=1;i<=I;i++{
		Fact=Fact*i
	}
	X := Data{Fact ,status}
	ch<-X
	// x++
}