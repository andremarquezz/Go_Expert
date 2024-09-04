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
	for _, cep := range os.Args[1:] {
		req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer a requisição %v\n", err)
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler o corpo da resposta %v\n", err)
		}

		var data ViaCEP
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer o unmarshal %v\n", err)
		}

		file, err := os.Create("cep_" + cep + ".txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar o arquivo %v\n", err)
		}
		defer file.Close()
		formattedString := fmt.Sprintf("CEP: %s\nLogradouro: %s\nBairro: %s\nLocalidade: %s\nUF: %s\nEstado: %s\nRegião: %s\nIBGE: %s\nDDD: %s\nSIAFI: %s\n\n", data.Cep, data.Logradouro, data.Bairro, data.Localidade, data.Uf, data.Estado, data.Regiao, data.Ibge, data.Ddd, data.Siafi)
		_, err = file.WriteString(formattedString)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao escrever no arquivo %v\n", err)
		}
		fmt.Printf("Arquivo cep_%s.txt criado com sucesso\n", cep)
	}
}
