package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := ":8080"
	http.HandleFunc("/", handleDollarQuotation)
	fmt.Fprintf(os.Stderr, "Servidor rodando na porta %s\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao iniciar o servidor: %v\n", err)
	}
}

func handleDollarQuotation(w http.ResponseWriter, r *http.Request) {
	// Lógica para lidar com a cotação do dólar
	fmt.Fprintf(w, "Cotação do dólar: 5.30") // Exemplo estático
}
