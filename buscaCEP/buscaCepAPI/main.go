package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	port := ":8080"
	http.HandleFunc("/", BuscaCEPHandler)

	fmt.Fprintf(os.Stderr, "Servidor rodando na porta %s\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao iniciar o servidor: %v\n", err)
	}

}

func BuscaCEPHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/" {
		http.Error(w, "Endpoint não encontrado", http.StatusNotFound)
		return
	}
	cep := r.URL.Query().Get("cep")
	if cep == "" {
		http.Error(w, "CEP não informado", http.StatusBadRequest)
		return
	}
	data, err := BuscaCEP(cep)
	if err != nil {
		http.Error(w, "Erro ao buscar o CEP", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(data)
}

func BuscaCEP(cep string) (*ViaCEP, error) {
	req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, fmt.Errorf("Erro ao fazer a requisição %v", err)
	}
	defer req.Body.Close()
	reqBody, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Errorf("Erro ao ler o corpo da resposta viaCEP %v", err)
	}
	var data ViaCEP
	err = json.Unmarshal(reqBody, &data)
	if err != nil {
		return nil, fmt.Errorf("Erro ao fazer o unmarshal %v", err)
	}
	return &data, nil
}
