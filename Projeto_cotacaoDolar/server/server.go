package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

type quotation struct {
	ID         string `gorm:"type:char(36);primaryKey"`
	Code       string
	Codein     string
	Name       string
	High       string
	Low        string
	VarBid     string
	PctChange  string
	Bid        string
	Ask        string
	Timestamp  string
	CreateDate string
}

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func connectDB() (*gorm.DB, error) {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&quotation{})
	return db, nil
}

func main() {
	db, err := connectDB()
	if err != nil {
		fmt.Fprintf(os.Stderr, "erro ao conectar ao banco de dados: %v\n", err)
		return
	}

	port := ":8080"
	http.HandleFunc("/cotacao", func(w http.ResponseWriter, r *http.Request) {
		handleDollarQuotation(w, db)
	})
	fmt.Fprintf(os.Stderr, "Servidor rodando na porta %s\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Fprintf(os.Stderr, "erro ao iniciar o servidor: %v\n", err)
	}
}

func handleDollarQuotation(w http.ResponseWriter, db *gorm.DB) {
	data, err := getQuotation()
	if err != nil {
		errorResponse := ErrorResponse{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)

		if err := json.NewEncoder(w).Encode(errorResponse); err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao codificar JSON de erro: %v\n", err)
		}
		return
	}
	if err := saveQuotation(db, *data); err != nil {
		http.Error(w, "Erro ao salvar cotação no banco de dados", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)

}

func getQuotation() (*USDBRL, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()

	URL := "https://economia.awesomeapi.com.br/json/last/USD-BRL"

	req, err := http.NewRequestWithContext(ctx, "GET", URL, nil)
	if err != nil {
		return nil, fmt.Errorf("erro durante a prepração da requisição: %w", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return nil, fmt.Errorf("requisição cancelada por timeout")
		}
		return nil, fmt.Errorf("erro ao fazer a requisição: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("requisição retornou status %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler o corpo da requisição: %w", err)
	}

	var response QuotationResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("erro ao fazer o unmarshal: %w", err)
	}
	return &response.USDBRL, nil
}

func saveQuotation(db *gorm.DB, data USDBRL) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	q := quotation{
		ID:         uuid.New().String(),
		Code:       data.Code,
		Codein:     data.Codein,
		Name:       data.Name,
		High:       data.High,
		Low:        data.Low,
		VarBid:     data.VarBid,
		PctChange:  data.PctChange,
		Bid:        data.Bid,
		Ask:        data.Ask,
		Timestamp:  data.Timestamp,
		CreateDate: data.CreateDate,
	}
	err := db.WithContext(ctx).Create(&q).Error
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return fmt.Errorf("tempo excedido ao salvar cotação")
		}
		return fmt.Errorf("erro ao salvar cotação: %w", err)
	}
	return nil

}
