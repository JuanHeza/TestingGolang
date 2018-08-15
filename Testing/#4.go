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
	NodoSolucion:=BuscarSolucion(Estado_inicial,Solucion)
	fmt.Println("NodoSolucion:",NodoSolucion)
	N:=NodoSolucion
	var aux Nodo
	for N.Get_Padre()!=Empty{
		Resultados=append(Resultados,N.Get_Datos())
		N.datos=N.Get_Padre()
		aux.datos=N.Get_Padre()
		N=BuscarPadre(aux.datos)
	}
	Resultados=append(Resultados,Estado_inicial)
	Reverse(Resultados)
	fmt.Println("Resultados:",Resultados)
	logger.Println("Final")
	fmt.Print(&buf)
}

func BuscarSolucion(Estado_inicial [4]int, Solucion [4]int) Nodo{
	var nodo, NodoInicial Nodo
	NodoInicial.datos=Estado_inicial
	Nodos_Frontera[len(Nodos_Frontera)]=NodoInicial
	for !Solucionado && len(Nodos_Frontera)!=0{
		nodo=Nodos_Frontera[len(Nodos_Frontera)-1]
		if nodo.datos == [4]int{4,3,2,1} {
			nodo=Nodos_Frontera[len(Nodos_Frontera)-2]
		}
		delete(Nodos_Frontera,len(Nodos_Frontera)-1)
		for Q:=0;Q<len(Nodos_Frontera);Q++{
			Nodos_Frontera[Q]=Nodos_Frontera[Q+1]
		}
		for I,X := range Nodos_Frontera{
			if X.datos==Empty{
				delete(Nodos_Frontera,I)
			}
		}
		Nodos_Visitados[len(Nodos_Visitados)]=nodo
		if nodo.Get_Datos()==Solucion{
			Solucionado=true
			break
		}else{
			Dato:=nodo.Get_Datos()
			
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
			nodo.Set_Hijos(Hijos)
			
			if !HijoIzquierdo.En_Lista(Nodos_Visitados) && !HijoIzquierdo.En_Lista(Nodos_Frontera){
				Nodos_Frontera[len(Nodos_Frontera)] = HijoIzquierdo
			}
			
			if !HijoCentral.En_Lista(Nodos_Visitados) && !HijoCentral.En_Lista(Nodos_Frontera){
				Nodos_Frontera[len(Nodos_Frontera)] = HijoCentral
			}
			
			if !HijoDerecho.En_Lista(Nodos_Visitados) && !HijoDerecho.En_Lista(Nodos_Frontera){
				Nodos_Frontera[len(Nodos_Frontera)] = HijoDerecho
			}
		}
	}
	return nodo
}