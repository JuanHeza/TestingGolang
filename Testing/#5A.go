package main

import (
	"bytes"
	"fmt"
	"log"
)

var Solucionado = false
var Nodos_Visitados = make(map[int]Nodo)
var Nodos_Frontera = make(map[int]Nodo)
var Resultados [][4]int
var Estado_inicial = [4]int{4,2,3,1}
var Solucion = [4]int{1,2,3,4}
var Hijo[4]int
var Empty = [4]int{0,0,0,0}
	var buf bytes.Buffer

func main(){
	logger := log.New(&buf, "Time: ", log.Lmicroseconds)
	logger.Println("Inicio")
	fmt.Println("Estado_inicial:",Estado_inicial)
	fmt.Println("Solucion:",Solucion)
	NodoSolucion:=BuscarSolucion(Estado_inicial,Solucion,Nodos_Visitados)
	fmt.Println("NodoSolucion:",NodoSolucion)
	
	var aux Nodo
	for NodoSolucion.Get_Padre()!=Empty{
		Resultados=append(Resultados,NodoSolucion.Get_Datos())
		NodoSolucion.datos=NodoSolucion.Get_Padre()
		aux.datos=NodoSolucion.Get_Padre()
		NodoSolucion=BuscarPadre(aux.datos)
	}
	Resultados=append(Resultados,Estado_inicial)
	Reverse(Resultados)
	fmt.Println("Resultados:",Resultados)
	log.Println("Nodos Visitados:",len(Nodos_Visitados))
	if len(Nodos_Visitados)==24{
		fmt.Println("Todos Los Nodos Visitados")
		fmt.Println(Nodos_Visitados)
	}
	logger.Println("Final")
	fmt.Print(&buf)
}

func BuscarSolucion(Estado_inicial [4]int, Solucion [4]int, Nodos_Visitados map[int]Nodo) Nodo{
	var sol, NodoInicial Nodo
	NodoInicial.datos=Estado_inicial
	Nodos_Visitados[len(Nodos_Visitados)]=NodoInicial
	for i,x:= range Nodos_Visitados{
		if x.datos==NodoInicial.datos&&len(Nodos_Visitados)>1{
			y:=Nodos_Visitados[i-1]
			NodoInicial.Set_Padre(y.datos)
		}
	}
	// log.Println(Nodos_Visitados)
		if NodoInicial.Get_Datos()==Solucion{
			Solucionado=true
			return NodoInicial
		}else{
			Dato:=NodoInicial.Get_Datos()
			
			var HijoIzquierdo Nodo
			Hijo = [4]int{Dato[1],Dato[0],Dato[2],Dato[3]}
			HijoIzquierdo.datos = Hijo
			
			var HijoCentral Nodo
			Hijo = [4]int{Dato[0],Dato[2],Dato[1],Dato[3]}
			HijoCentral.datos = Hijo
			
			var HijoDerecho Nodo
			Hijo = [4]int{Dato[0],Dato[1],Dato[3],Dato[2]}
			HijoDerecho.datos = Hijo
			
			var Hijos = make(map[int]*Nodo)
			Hijos[len(Hijos)]=&HijoIzquierdo
			Hijos[len(Hijos)]=&HijoCentral
			Hijos[len(Hijos)]=&HijoDerecho
			NodoInicial.Set_Hijos(Hijos)
			// fmt.Println("NodoInicial:",NodoInicial)
			
			for _,NodoHijo := range NodoInicial.Get_Hijos(){
			var NH Nodo
			NH.datos=NodoHijo
				if !NH.En_Lista(Nodos_Visitados){
					sol=BuscarSolucion(NodoHijo,Solucion,Nodos_Visitados)
					if sol.datos==Empty&&sol.padre==Empty&&sol.hijo==nil{
						return sol
					}
				}
			}
		}
		// fmt.Println(">",NodoInicial,"<")
	return NodoInicial
}