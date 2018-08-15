package main

import (
  "fmt"
  "time"
  "strconv"
  )

func main() {
  var Confianza bool
  fmt.Println(Confianza)
    for Confianza== false{
    K:=30
    fmt.Println(K)
    Z:=time.Now()
    X:=Z.Second()
    fmt.Println(X)
    Y:=Z.Nanosecond()
    fmt.Println(Y)
    if X==0{
      return 
    }
    random := (K*Y) % X
    for i:=0;i<K;i++{
      Aux:=random
      Q:=float64(random)/float64(X)
      if Q<=0{
        Confianza=false
      }
      fmt.Println("[",i+1,"]",random,"<",strconv.FormatFloat(Q, 'f', 5, 64),">")
      random = (K*random) % X
      if Aux!=random{
        Confianza=true
      }
    }
  }
  if Confianza==false{
    fmt.Println("Generador No Confiable")
  }else{
    fmt.Println("Generador Confiable")
  }
}