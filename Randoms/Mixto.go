package main

import (
	"fmt"
    "strconv"
)

func Mixto(a string, xo string, c string, m string) []string {
	var REC float64
	var XN,K int
	var X string
	A, err:=strconv.Atoi(a)
	XO, err:=strconv.Atoi(xo)
	oro:=XO
	C, err:=strconv.Atoi(c)
	M, err:=strconv.Atoi(m)
	if err != nil {
			fmt.Print(err)
		}
	var Salida[] string
	for i:=1;i<=M;i++{
		XN=(A*XO+C)%M;
		K=(A*XO+C)/M;
		REC=float64(XN)/float64(M);
		X=fmt.Sprintf("\n%3d\t<%2d>\t%2d+%2d/%2d\t<%2d>\t[%.3f]",i,XO,K,XN,M,XN,REC)
		Salida=append(Salida,X)
		XO=XN;
	}
	if oro==XN{
		Salida=append(Salida,"\nGenerador Confiable")
	}else{
		Salida=append(Salida,"\nGenerador No Confiable")

	}
	 return Salida
}



