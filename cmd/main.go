package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq" // Importa o driver PostgreSQL
)

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

	// Exibir tabelas disponíveis no banco
	checkTables(db)

	// Lógica principal
	listOrders(db)
}

// Função para verificar tabelas disponíveis
func checkTables(db *sql.DB) {
	rows, err := db.Query("SELECT tablename FROM pg_catalog.pg_tables WHERE schemaname = 'public'")
	if err != nil {
		log.Fatalf("Erro ao listar tabelas: %v", err)
	}
	defer rows.Close()

	fmt.Println("Tabelas disponíveis no banco de dados:")
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			log.Fatalf("Erro ao ler tabelas: %v", err)
		}
		fmt.Printf("- %s\n", tableName)
	}
	fmt.Println("Fim da listagem de tabelas.")
}

// Função para listar pedidos no banco de dados
func listOrders(db *sql.DB) {
	rows, err := db.Query("SELECT id, customer, total FROM orders")
	if err != nil {
		log.Fatalf("Erro ao executar a consulta: %v. Verifique se a tabela 'orders' existe.", err)
	}
	defer rows.Close()

	fmt.Println("Pedidos:")
	hasRows := false
	for rows.Next() {
		var id int
		var customer string
		var total float64

		err := rows.Scan(&id, &customer, &total)
		if err != nil {
			log.Fatalf("Erro ao ler os resultados: %v", err)
		}
		fmt.Printf("ID: %d, Cliente: %s, Total: %.2f\n", id, customer, total)
		hasRows = true
	}

	if !hasRows {
		fmt.Println("Nenhum pedido encontrado.")
	}

	// Confirmar que o loop foi concluído
	fmt.Println("Fim da listagem de pedidos.")
}
