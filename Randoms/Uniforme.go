package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
)

func main() {
var R[] float64
var K[] float64
var A, B, Sum, aux float64
	
	R=Numeros();
	fmt.Print("\tDefine A:\t")
	fmt.Scanf("%f\n",&A)
	fmt.Print("\tDefine B:\t")
	fmt.Scanf("%f\n",&B)
	for i:=0;i<len(R);i++{
		aux=A+(B-A)*R[i]
		aux, err := strconv.ParseFloat(strconv.FormatFloat(aux, 'f', 5, 64), 64)
		if err== nil{
			K=append(K,aux)
		}
	Sum+=aux
	}
	fmt.Println("\n\t N\t    R\t\t    K")
	for i:=0;i<len(R);i++{
	fmt.Println("\t",i+1,"\t",R[i],"\t",K[i])
	}
	Sum, err := strconv.ParseFloat(strconv.FormatFloat(Sum, 'f', 5, 64), 64)
	if err== nil{
		fmt.Println("\tSumatoria:\t",Sum)	
	}	
	Prom:=Sum/float64(len(K))
	Prom, err = strconv.ParseFloat(strconv.FormatFloat(Prom, 'f', 5, 64), 64)
	if err== nil{
		fmt.Println("\tPromedio:\t",Prom)
	}	
	
	 bufio.NewReader(os.Stdin).ReadBytes('\n') 
}


