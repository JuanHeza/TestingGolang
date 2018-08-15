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
var buf bytes.Buffer
var Conexiones = make(map[string][]string)
var LOG int

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
	log.Println(LOG++)
	fmt.Println("Solucion:",Solucion)
	log.Println(LOG++)
	NodoSolucion:=BuscarSolucion(Conexiones,Estado_inicial,Solucion)
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
	// Reverse(Resultados)
	fmt.Println("Resultados:",Resultados)
	logger.Println("Final")
	fmt.Print(&buf)
}

func BuscarSolucion(Conexiones  map[string][]string,Estado_inicial string, Solucion string) Nodo{
	var nodo, NodoInicial Nodo
	log.Println(LOG++)
	NodoInicial.datos=Estado_inicial
	Nodos_Frontera[len(Nodos_Frontera)]=NodoInicial
	log.Println(LOG++)
	for !Solucionado && len(Nodos_Frontera)!=0{
		log.Println(LOG++)
		nodo=Nodos_Frontera[0]
		Nodos_Visitados[len(Nodos_Visitados)]=Nodos_Frontera[0]
		delete(Nodos_Frontera,0)
/* 		for Q:=0;Q<len(Nodos_Frontera);Q++{
			Nodos_Frontera[Q]=Nodos_Frontera[Q+1]
		}
		for I,X := range Nodos_Frontera{
			if X.datos==Empty{
				delete(Nodos_Frontera,I)
			}
		}
		
		Nodos_Visitados[len(Nodos_Visitados)]=nodo */
		if nodo.Get_Datos()==Solucion{
			Solucionado=true
			break
		}else{
			Dato:=nodo.Get_Datos()
			///////
			var lista_hijos = make(map[int]*Nodo)
			var hijo Nodo
			for _,un_hijo:= range Conexiones[Dato]{
				hijo.datos=un_hijo
				lista_hijos[len(lista_hijos)]=&hijo
				if !hijo.En_Lista(Nodos_Visitados) && !hijo.En_Lista(Nodos_Frontera){
					Nodos_Frontera[len(Nodos_Frontera)]=hijo
				}
				nodo.Set_Hijos(lista_hijos)
			}
			///////
		}
	}
	return nodo
}