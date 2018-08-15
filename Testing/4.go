package main

import "fmt"

var Hijos =make(map[int]*Son)
type Son struct{
  A,B,C,D int
}
func main() {
  HijoIzquierdo := Son{5,6,7,8}
  HijoCentral := Son{9,4,1,2}
  HijoDerecho := Son{1,2,3,4}
  HijoNulo := Son{1,5,7,3}
  var base Son
	Hijos[len(Hijos)]=&HijoIzquierdo
	Hijos[len(Hijos)]=&HijoCentral
	Hijos[len(Hijos)]=&HijoDerecho
	Hijos[len(Hijos)]=&HijoNulo
	for i,x := range Hijos{
	  fmt.Println(*x)
	  if i==0{
	    base.A=x.A
	  }
	  if i==1{
	    base.B=x.B
	  }
	  if i==2{
	    base.C=x.C
	  }
	  if i==3{
	    base.D=x.D
	  }		
	  fmt.Println("<",base,">")
	}
}