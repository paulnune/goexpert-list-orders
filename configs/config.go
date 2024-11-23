package configs

import (
	"log"
	"os"
)

func LoadConfig() {
	requiredVars := []string{"DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME"}

	for _, v := range requiredVars {
		if os.Getenv(v) == "" {
			log.Fatalf("Variável de ambiente %s não configurada", v)
		}
	}

	log.Println("Configurações carregadas com sucesso")
}
