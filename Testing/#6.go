package main

import (
	"bytes"
	"fmt"
	"log"
)

var variables = [2]int {0,0}
var optimo = [2]int {0,0}
var sol int
var rango_variales [2][]int 
var buf bytes.Buffer


func main(){
	rango_variales[0]= makeRange(0, 50)
	rango_variales[1]= makeRange(0, 75)
	logger := log.New(&buf, "Time: ", log.Lmicroseconds)
	logger.Println("Inicio")
	sol:=backtracking(variables,rango_variales,optimo,0)
	log.Println("Mejor Solucion:")
	log.Println(sol[0],"Pantalones",sol[1],"Camisetas")
	log.Println("Beneficio:",evalua_solucion(sol))
	logger.Println("Final")
	fmt.Println(&buf)
}

func backtracking(variables [2]int, rango_variales [2][]int, optimo[2]int, profundidad int) [2]int{
	min:=rango_variales[profundidad][0]
	max:=rango_variales[profundidad][len(rango_variales[profundidad])-1]
	for v:=min ; v <= max ; v++{
		variables[profundidad]=v
		if profundidad < len(variables)-1{
			log.Println(escompletable(variables),"Variables:",variables)
			if escompletable(variables){
				optimo = backtracking(variables,rango_variales,optimo, profundidad+1)
				log.Println(optimo)
			}else{
				fmt.Println("else",variables)
				sol = evalua_solucion(variables)
				if sol > evalua_solucion(optimo) && escompletable(variables){
					optimo = [2]int {variables[0],variables[1]}
				}
			}
		}
	}
	return optimo
}

func evalua_solucion(variables [2]int) int{
	x1 := variables[0]
	x2 := variables[1]
	val := (12-6)*x1+(8-4)*x2
	// log.Println("Val := (12-6)*",x1,"+(8-4)*",x2,":",val)
	return val
}

func escompletable(variables [2]int) bool{
	x1 := variables[0]
	x2 := variables[1]
	val1 := 7*x1 + 4*x2
	val2 := 6*x1 + 5*x2
	if val1<=150 && val2<=160{
	fmt.Println(val1,"<=150 && ",val2,"<=160")
		return true
	}else{
		return false
	}
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
