package main

/* import (
	"fmt"
	"bufio"
	"log"
	"os"
) */
type Nodo struct {
	datos 	[4]int
	padre	[4]int
	hijo	map[int][4]int
	coste	int
//	set_hijos(hijos)
}

func (n *Nodo) Set_Hijos(hijos map[int]*Nodo){
	n.hijo=make(map[int][4]int)
	for _,K:=range hijos{
		n.hijo[len(n.hijo)]=K.datos
	}
	// fmt.Println(n)
	if len(n.hijo)!=0{
		for _,Y:=range hijos{
			Y.padre=n.datos
			// log.Println(Y)
		}
	}
	 // bufio.NewReader(os.Stdin).ReadBytes('\n') 
}

func (n *Nodo) Set_Padre(Padre [4]int){
	n.padre = Padre
}

func (n *Nodo) Set_Datos(Datos [4]int){
	n.datos = Datos
}

func (n *Nodo) Set_Coste(Coste int){
	n.coste = Coste
}

func (n *Nodo) Get_Hijos() map[int][4]int{
	return n.hijo
}

func (n *Nodo) Get_Padre() [4]int{
	return n.padre
}

func (n *Nodo) Get_Datos() [4]int{
	return n.datos
}

func (n *Nodo) Get_Coste() int{
	return n.coste
}

func (n *Nodo) igual(nodo Nodo) bool{
	if n.Get_Datos()==nodo.Get_Datos(){
		return true
	}else{
		return false
	}
}

func (n *Nodo) En_Lista(lista map[int]Nodo) bool{
	en_la_lista:=false
	for _,X:=range lista{
		if n.igual(X){
			en_la_lista=true
		}
	} 
	return en_la_lista
}

func Reverse(Resultados [][4]int){
	for i, j := 0, len(Resultados)-1; i < j; i, j = i+1, j-1 {
		Resultados[i], Resultados[j] = Resultados[j], Resultados[i]
	}	
}

func BuscarPadre(hijo [4]int) Nodo{
	var padre Nodo
	for _,x := range (Nodos_Visitados){
		if x.datos==hijo{
			padre=x
			break
		}
	}
	return padre
}