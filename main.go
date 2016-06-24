package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	"encoding/json"
)

type jsonobject struct {
	Categorias []Categoria
}

type Categoria struct {
	Name   string
	Cursos []Curso
}

type Curso struct {
	Name string
	Url  string
}

func main() {
	file, e := ioutil.ReadFile("./cursos.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	var jsontype jsonobject
	json.Unmarshal(file, &jsontype)

	tmpl, err := template.New("README.tmpl").ParseFiles("README.tmpl")

	f2, err := os.Create("./README.md")
	defer f2.Close()

	w2 := bufio.NewWriter(f2)
	defer w2.Flush()

	err = tmpl.Execute(w2, jsontype)

	if err != nil {
		panic(err)
	}
}
