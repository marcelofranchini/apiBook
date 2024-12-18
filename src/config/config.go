package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DbString = ""

	Porta = 0
)

func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 8000
	}

	DbString = os.Getenv("DB_SQL")

	// result := fmt.Sprintf("Meu nome é %s e eu tenho %d anos.", name, age)

}
