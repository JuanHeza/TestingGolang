package main

import (
    "fmt"
    "math"
    "strconv"
)

func Multi(a string, xo string, m string) []string {
	var XN,p,K int
	var REC,PE float64
	var X string
	A, err:=strconv.Atoi(a)
	XO, err:=strconv.Atoi(xo)
	oro:=XO
	M, err:=strconv.Atoi(m)
		if err != nil {
			fmt.Print(err)
		}
	var Salida[] string
	if M%10==0{
		mexp:=M;
		for p=0;mexp>1;p++{
			mexp=mexp/10;
		}
		if p>=5{
			PE=math.Pow(10,float64(p-2));
			PE=PE*5;
		}else{
			PE=math.Pow(5,float64(p-1));
			PE=4*PE;
		}
	}else{
		mexp:=M;
		for p=0;mexp>1;p++{
			mexp=mexp/2;
		}
		PE=float64(M/4);
	}
	for i:=1.0;i<=PE;i++{
		XN=(A*XO)%M;
		K=(A*XO)/M;
		REC=float64(XN)/float64(M);
		X=fmt.Sprintf("\n%3.0f\t<%2d>\t%2d+%2d/%2d\t<%2d>\t[%.3f]",i,XO,K,XN,M,XN,REC)
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



