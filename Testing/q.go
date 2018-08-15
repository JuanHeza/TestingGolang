package main

import (
	"os"
	"text/template"
)

func main() {
	// Define a template.
	const tmpl = 
`The name is {{.Id}}.{{range .Name}}
His email id is {{.}}{{end}}
Value {{.X}}`

	// Prepare some data to insert into the template.
	type User struct {
		Id int
		Name []string
		X int
	}
	
	type UserList []User
	var myuserlist User  = User {1, []string{"a", "b", "c"},19}

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("tmpl").Parse(tmpl))
	
	t.Execute(os.Stdout, myuserlist)
}
/*
The name is 1.
 His email id is a
 His email id is b
 His email id is c
Value 19
*/