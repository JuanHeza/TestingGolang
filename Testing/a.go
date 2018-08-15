package main

import (
    "fmt"

    "github.com/fatih/structs"
)

type A struct {
    Foo string
    Bar int
}

func main() {
	X:=A{Foo:"hola", Bar:16}
    names := structs.Names(X)
    fmt.Println(names) // ["Foo", "Bar"]
	for i:= range names{
		fmt.Printf("%v - %t\n",names[i],names[i])
	}
    field := structs.Values(X)
	for i:= range field{
		fmt.Printf("%v - %t\n",field[i],field[i])
	}
}

/* 
    "github.com/fatih/structs"
    q := structs.Names(w)
	fmt.Println(q)
	i := structs.IsStruct(w.Settings)
	if i==true{
		names:= structs.Names(w.Settings)
			fmt.Println("Variables de \"W\"")
			for i:= range names{
				fmt.Printf("%v - %t\n",names[i],(names[i]))
			}
		field := structs.Values(w.Settings)
			fmt.Println("\nValores de \"W\"")
			for i:= range field{
				fmt.Printf("%v - %t\n",field[i],field[i])
			}
	}else{
		fmt.Println("No Struct")
		fmt.Printf("%v - %t",w.Settings,w.Settings)
	} */