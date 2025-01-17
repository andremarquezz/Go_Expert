package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	http.HandleFunc("/", TemplateHandler)
	port := ":8080"
	log.Printf("Servidor rodando na porta %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v\n", err)
	}

}

func TemplateHandler(w http.ResponseWriter, r *http.Request) {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}
	cursos := Cursos{
		Curso{"Go", 40},
		Curso{"Java", 60},
		Curso{"Python", 45},
	}

	t := template.New("content.html")
	t.Funcs(template.FuncMap{"ToUpper": ToUpper})
	t = template.Must(t.ParseFiles(templates...))
	err := t.Execute(w, cursos)
	if err != nil {
		http.Error(w, "Erro ao processar o template", http.StatusInternalServerError)
	}
}

func ToUpper(s string) string {
	return strings.ToUpper(s)
}
