package main

import (
	"bytes"
	"fmt"
	"log"
)

var Solucionado = false
var Nodos_Visitados = make(map[int]Nodo)
var Nodos_Frontera = make(map[int]Nodo)
var Resultados []string
var Estado_inicial = "Malaga"
var Solucion = "Santiago"
var Hijo[4]int
var Empty = ""
var sol, NodoInicial Nodo
	var buf bytes.Buffer
var Conexiones = make(map[string][]string)

func main(){
	Conexiones = map[string][]string{
		"Malaga":{"Salamanca","Madrid","Barcelona"},
		"Sevilla":{"Santiago","Madrid"},
		"Granada":{"Valencia"},
		"Valencia":{"Barcelona"},
		"Madrid":{"Salamanca","Sevilla","Malaga","Barcelona","Santander"},
		"Salamanca":{"Madrid","Malaga"},
		"Santiago":{"Sevilla","Santander","Barcelona"},
		"Santander":{"Santiago","Madrid"},
		"Zaragoza":{"Barcelona"},
		"Barcelona":{"Zaragoza","Santiago","Madrid","Malaga","Valencia"},
	}
	logger := log.New(&buf, "Time: ", log.Lmicroseconds)
	logger.Println("Inicio")
	fmt.Println("Estado_inicial:",Estado_inicial)
	fmt.Println("Solucion:",Solucion)
	NodoSolucion:=DFS_Iter(Estado_inicial,Solucion)
	fmt.Println("NodoSolucion:",NodoSolucion)
	if NodoSolucion.datos==Empty&&NodoSolucion.padre==Empty&&NodoSolucion.hijo==nil{
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
	}else{
		log.Println("Solucion No Encontrada")
	}
	logger.Println("Final")
	fmt.Println(&buf)
}

func DFS_Iter(Estado_inicial string, Solucion string) Nodo{
var X Nodo
	for i:=0;i<100;i++{
		sol=BuscarSolucion(Estado_inicial,Solucion,Nodos_Visitados, i)
		if sol.datos==Empty&&sol.padre==Empty&&sol.hijo==nil{
			X = sol
		}
	}
	return X
}

func BuscarSolucion(NodoHijo string,Solucion string,Nodos_Visitados map[int]Nodo, limite int) Nodo{
	var lista_hijos = make(map[int]*Nodo)
	var hijo, X Nodo 
	if limite > 0{
		NodoInicial.datos=NodoHijo
		Nodos_Visitados[len(Nodos_Visitados)]=NodoInicial
		if NodoInicial.Get_Datos()==Solucion{
			X = NodoInicial
		}else{
			Dato:=NodoInicial.Get_Datos()
			for _,un_hijo:= range Conexiones[Dato]{
				hijo.datos=un_hijo
				if !hijo.En_Lista(Nodos_Visitados) {
					lista_hijos[len(lista_hijos)]=&hijo
				}
			}
		}
		NodoInicial.Set_Hijos(lista_hijos)
		for _,NodoHijo := range NodoInicial.Get_Hijos(){
			var NH Nodo
			NH.datos=NodoHijo
			if !NH.En_Lista(Nodos_Visitados){
				sol=BuscarSolucion(NodoHijo,Solucion,Nodos_Visitados, limite-1)
				if sol.datos==Empty&&sol.padre==Empty&&sol.hijo==nil{
					return sol
				}
			}
		}
	}
	return X
}