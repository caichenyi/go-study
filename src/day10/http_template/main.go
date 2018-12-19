package main

import (
	"fmt"
	"html/template"
	"os"
)

type Person struct {
	Name string
	Age string
}

func main() {
	t, err := template.ParseFiles("C:/Users/chenyi.cai/Workspace/code/go-study/src/day10/http_template/index.html")
	if err != nil {
		fmt.Println("parse file err: ", err)
		return
	}
	p := Person{Name: "Mary", Age: "31"}
	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println("There was an error: ", err.Error())
	}
}