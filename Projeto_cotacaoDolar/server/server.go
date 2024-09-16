package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type USDBRL struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

type QuotationResponse struct {
	USDBRL USDBRL `json:"USDBRL"`
}

func main() {
	port := ":8080"
	http.HandleFunc("/cotacao", handleDollarQuotation)
	fmt.Fprintf(os.Stderr, "Servidor rodando na porta %s\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Fprintf(os.Stderr, "erro ao iniciar o servidor: %v\n", err)
	}
}

func handleDollarQuotation(w http.ResponseWriter, r *http.Request) {
	data, err := getQuotation()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func getQuotation() (*USDBRL, error) {
	req, err := http.Get("https://economia.awesomeapi.com.br/json/last/USD-BRL")
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer a requisição: %w", err)
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler o corpo da requisição: %w", err)
	}

	var response QuotationResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("erro ao fazer o unmarshal: %w", err)
	}

	return &response.USDBRL, nil
}
