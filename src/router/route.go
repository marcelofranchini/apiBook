package router

import (
	"apiBook/src/router/rotas"

	"github.com/gorilla/mux"
)

func Gerar() *mux.Router {

	router := mux.NewRouter()

	rotas.ConfigurarRotas(router)

	return router

}
