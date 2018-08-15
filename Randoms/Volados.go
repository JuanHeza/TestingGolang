package main

import (
  "fmt"
  "bufio"
  "os"
)

func main() {
var R[] float64
//var K[] float64
var A, B, C, X/*, Sum, aux, Quiebra, Meta*/ float64
	Q:="---"
	R=Numeros();
	fmt.Print("\tDefine La Apuesta:\t")
	fmt.Scanf("%f\n",&A)
	fmt.Print("\tDefine Cantidad Inicial:\t")
	fmt.Scanf("%f\n",&B)
	fmt.Print("\tDefine Meta:\t")
	fmt.Scanf("%f\n",&C)
	auxA:=A
	auxB:=B
	X=1
	fmt.Println("\tAleatorio\tCorrida\tInicial\tApuesta\t¿Gano?\tFinal\tMeta")
	for i:=0;i<len(R);i++{
		fmt.Print("\t",R[i],"\t\t",X,"\t",B,"\t",A)
		if R[i]<=0.5{
			B+=A
			A=auxA
		}else{
			B-=A
			A=A*2
			if A>B{
				A=B
			}
		}
		if B>=C{
			Q="Meta"
		}else if B<=0{
			Q="Quiebra"
		}else{
			Q="---"
		}
		fmt.Println("\t",R[i]<=0.5,"\t",B,"\t",Q)
		if B>=C||B<=0{
			B=auxB
			A=auxA
			X++
		}
	}
	
	/*
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
	*/	
	 bufio.NewReader(os.Stdin).ReadBytes('\n') 
}
