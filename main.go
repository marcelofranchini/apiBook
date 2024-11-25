package main

import (
	"apiBook/src/config"
	"apiBook/src/db"
	"apiBook/src/router"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	config.Carregar()

	// Conectar ao banco de dados
	connection, err := db.ConectarDb()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco: %v", err)
	}
	defer connection.Close()
	fmt.Println("Conexão bem-sucedida com o banco de dados!")

	router := router.Gerar()
	fmt.Println("rodando na porta:" + strconv.Itoa(config.Porta))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(config.Porta), router))
}
