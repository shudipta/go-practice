package main

import (
	"os"
	"text/template"
	"fmt"
)

type Person struct {
	Name string
	NonExportedAgeField string //because it doesn't start with a capital letter
}

func main() {
	p:= Person{Name: "Mary", NonExportedAgeField: "31"}

	t := template.New("nonexported template demo")
	t, _ = t.Parse("hello {{.Name}}! Age is {{.NonExportedAgeField}}.\n")
	err := t.Execute(os.Stdout, p)
	if err != nil {
		fmt.Println("There was an error:", err.Error())
	}

	// ==================================
	//
	//tOk := template.New("first")
	//template.Must(tOk.Parse(" some static text /* and a comment */")) //a valid template, so no panic with Must
	//tOk.Execute(os.Stdout, nil)
	//fmt.Println("The first one parsed OK.")
	//
	//template.Must(template.New("second").Parse("some static text {{ .Name }}"))
	//fmt.Println("The second one parsed OK.")
	//
	//fmt.Println("The next one ought to fail.")
	//tErr := template.New("check parse error with Must")
	//template.Must(tErr.Parse(" some static text {{ .Name }")) // due to unmatched brace, there should be a panic here

	// ==================================

	tEmpty := template.New("template test")
	tEmpty = template.Must(tEmpty.Parse("Empty pipeline if demo: {{if ``}} Will not print. {{end}}\n")) //empty pipeline following if
	tEmpty.Execute(os.Stdout, nil)

	tWithValue := template.New("template test")
	tWithValue = template.Must(tWithValue.Parse("Non empty pipeline if demo: {{if `anything`}} Will print. {{end}}\n")) //non empty pipeline following if condition
	tWithValue.Execute(os.Stdout, nil)

	tIfElse := template.New("template test")
	tIfElse = template.Must(tIfElse.Parse("if-else demo1: {{if ``}} Print IF part. {{else}} Print ELSE part.{{end}}\n")) //non empty pipeline following if condition
	tIfElse.Execute(os.Stdout, nil)
	tIfElse = template.Must(tIfElse.Parse("if-else demo2:  {{define \"a\"}} {{with $x := `hello`}} {{$x}}{{end}} {{if ``}} Print IF part. {{else}} Print ELSE part.{{end}}{{end}}\n")) //non empty pipeline following if condition
	tIfElse.Execute(os.Stdout, nil)

	// ==================================

	t, _ = template.New("test").Parse("{{with `hello`}}{{end}}!\n")
	t.Execute(os.Stdout, nil)

	t1, _ := template.New("test").Parse("{{with `hello`}}{{.}} {{with `Mary`}}{{.}}{{end}}{{end}}!\n") //when nested, the dot takes the value according to closest scope.
	t1.Execute(os.Stdout, nil)

	// ==================================
	// ==================================
	// ==================================
	// ==================================
	// ==================================
	// ==================================
	// ==================================
}

