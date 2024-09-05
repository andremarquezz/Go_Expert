package main

import (
	"log"
	"net/http"
	"text/template"
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

	t := template.Must(template.New("content.html").ParseFiles(templates...))
	err := t.Execute(w, cursos)
	if err != nil {
		http.Error(w, "Erro ao processar o template", http.StatusInternalServerError)
	}
}
