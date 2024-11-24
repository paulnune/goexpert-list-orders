package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq" // Importa o driver PostgreSQL
)

// Order representa a estrutura de um pedido
type Order struct {
	ID       int     `json:"id"`
	Customer string  `json:"customer"`
	Total    float64 `json:"total"`
}

func main() {
	// Exibir as variáveis de ambiente carregadas
	fmt.Printf("DB_HOST: %s\n", os.Getenv("DB_HOST"))
	fmt.Printf("DB_PORT: %s\n", os.Getenv("DB_PORT"))
	fmt.Printf("DB_USER: %s\n", os.Getenv("DB_USER"))
	fmt.Printf("DB_PASSWORD: %s\n", os.Getenv("DB_PASSWORD"))
	fmt.Printf("DB_NAME: %s\n", os.Getenv("DB_NAME"))

	// Montar a string de conexão com o banco de dados
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	fmt.Printf("String de conexão: %s\n", connStr)

	// Retry na conexão ao banco de dados
	var db *sql.DB
	var err error
	for retries := 0; retries < 10; retries++ {
		db, err = sql.Open("postgres", connStr)
		if err == nil {
			err = db.Ping()
			if err == nil {
				fmt.Println("Conexão com o banco de dados bem-sucedida!")
				break
			}
		}
		fmt.Printf("Erro ao conectar: %v. Tentando novamente...\n", err)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Fatalf("Não foi possível conectar ao banco de dados após várias tentativas: %v", err)
	}
	defer db.Close()

	// Configuração das rotas
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		// Configurar cabeçalhos para JSON
		w.Header().Set("Content-Type", "application/json")

		// Obter a lista de pedidos
		orders, err := listOrders(db)
		if err != nil {
			http.Error(w, fmt.Sprintf("Erro ao buscar pedidos: %v", err), http.StatusInternalServerError)
			return
		}

		// Retornar a lista como JSON
		json.NewEncoder(w).Encode(orders)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Iniciar o servidor HTTP
	port := "8080"
	fmt.Printf("Servidor rodando na porta %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// listOrders consulta a tabela de pedidos no banco de dados
func listOrders(db *sql.DB) ([]Order, error) {
	rows, err := db.Query("SELECT id, customer, total FROM orders")
	if err != nil {
		return nil, fmt.Errorf("erro ao executar a consulta: %v. Verifique se a tabela 'orders' existe.", err)
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var order Order
		if err := rows.Scan(&order.ID, &order.Customer, &order.Total); err != nil {
			return nil, fmt.Errorf("erro ao ler os resultados: %v", err)
		}
		orders = append(orders, order)
	}

	return orders, nil
}
