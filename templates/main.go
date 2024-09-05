package main

import (
	"log"
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	cursos := Cursos{
		Curso{"Go", 40},
		Curso{"Java", 60},
		Curso{"Python", 45},
	}

	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := t.Execute(os.Stdout, cursos)
	if err != nil {
		log.Fatal(err)
	}
}
