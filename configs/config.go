package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis de ambiente padrão")
	}

	required := []string{"DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME"}
	for _, v := range required {
		if os.Getenv(v) == "" {
			log.Fatalf("Variável de ambiente %s não configurada", v)
		}
	}
}
