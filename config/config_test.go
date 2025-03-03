package config_test

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
)

func TestLoadConfig(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  Arquivo .env não encontrado, carregando variáveis do sistema")
	}
}
