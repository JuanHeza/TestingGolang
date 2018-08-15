package main

import (
  "fmt"
  "os"
)

func Numeros(N int) []float64 {
	f, err := os.Open("NAU.txt")
	if err != nil {
		fmt.Println(err)
	}
	var NAU[10][8][5] float64
	i:=0
	x:=0
	k:=0
	j:=0
	for {
		var flt float64
		var n int
		n, err = fmt.Fscanln(f, &flt)
		if n ==0 {
			break
		}
		NAU[i][k][j]=flt
		i++
		if(err==nil){
			x++
			k=x/5
			i=0
			j=(x)%5
		}
	}
	var Num[]float64
	var col, fil int
	//fmt.Print("\tNumeros:\t")
	//fmt.Scanf("%d\t",&N)
	R:=(N-1)/40
	for C:=0;C<=(R);C++{
		fmt.Print("\tColumna:\t")
		fmt.Scanf("%d\t",&col)
		fmt.Print("\tFila:\t\t")
		fmt.Scanf("%d\t",&fil)
		if N>40{
			for	nums:=1;nums<=8||nums<=8;nums++ {
				for i:=0;i<5;i++{
					Num=append(Num,NAU[col-1][fil-1][i])
				}
				fil++			
			}	
			N-=40
		}else{
			for	nums:=1;nums<=(N/5);nums++ {
				for i:=0;i<5;i++{
					Num=append(Num,NAU[col-1][fil-1][i])
				}
				fil++			
			}
		}
	}
	return Num
}