package main

import (
  "fmt"
  "math"
  "bufio"
  "os"
  "strconv"
)

func main() {
var R[] float64
var K[] float64
var lamda, Sum, aux float64
	////////////////////
var N int
	fmt.Print("\tNumeros:\t")
	fmt.Scanf("%d\t",&N)
	R=Numeros(N);
	fmt.Print("\tDefine Lamda:\t")
	fmt.Scanf("%f\n",&lamda)
	for i:=0;i<len(R);i++{
		aux=(-1/lamda)*math.Log(R[i])
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