package routes

import (
	"api-rest/controllers"
	"api-rest/middleware"
	"log"
	"net/http"

	//gorilla mux e praticamente o que o http faz so que melhor
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// inicia o listen and serve
func HandleRequest() {
	r := mux.NewRouter()

	r.Use(middleware.ContentTypeMiddleware)
	// aqui utilizamos o middleware com o mux

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("GET")
	r.HandleFunc("/api/personalidades", controllers.CriarNovaPersonalidade).Methods("POST")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaPersonalidade).Methods("DELETE")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.AtualizarPersonalidade).Methods("PUT")

	// basicamente aqui precisamos adicionar o cors para que outras aplicações possam realizar a leitura
	// não sei como funcionaria em um ambiente homologado mas com base naquilo que vi os projetos podem ser acessados
	// normalmente sem restrição.

	// para isso utilizamos a biblioteca github.com/gorilla/handlers  do proprio gorilla mux
	log.Fatal(http.ListenAndServe(":8000", r), handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r))
}
