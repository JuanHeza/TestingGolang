/* package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	go func (){ ch <- 2}()
	go func (){ ch <- 1}()
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
 */
 
 package main

import "fmt"
import "log"
 var x int
func sum(c chan int,p int){
	switch p {
		case 1:
			log.Println("Numero:",<-c)
		case 2:
			fmt.Println("Numero:",<-c)
		}
	x++
	c<-x
}

func main(){
	c:=make(chan int, 2)
	x=0
	c<-x
	for x<10{
		go sum(c,1)
		 sum(c,2)
	}
	close(c)
}
/* func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
 */