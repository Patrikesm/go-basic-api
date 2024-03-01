package controllers

import (
	"api-rest/database"
	"api-rest/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// quando bate liga o template se tivesse
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	// aqui conseguimos adicionar que o content type é do tipo json, basicamente eu defini um header padrão
	//w.Header().Set("content-type", "application/json")

	// quando utilizamos o gorm podemos apenas definir uma váriavel array daquele modelo que queremos retornar
	// ele vai fazer automaticamente o retorno e meio que escanear dentro do da variavel
	var p []models.Personalidade

	database.DB.Find(&p)

	json.NewEncoder(w).Encode(p)
}

func RetornaUmaPersonalidades(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Personalidade

	database.DB.First(&p, id)

	json.NewEncoder(w).Encode(p)

	// nesse caso é necessário sempre passar o endereço de memória pra que ele entre na função
	// e faça a alteração na variavel que definimos
}

func CriarNovaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var personsalidade models.Personalidade

	json.NewDecoder(r.Body).Decode(&personsalidade)

	database.DB.Create(&personsalidade)

	json.NewEncoder(w).Encode(personsalidade)

}

func AtualizarPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personsalidade models.Personalidade

	database.DB.First(&personsalidade, id)

	//busquei a personalidade

	json.NewDecoder(r.Body).Decode(&personsalidade)

	//decodei os valores no objeto

	database.DB.Save(&personsalidade)

	//salvei o objeto

	json.NewEncoder(w).Encode(personsalidade)

}

func DeletaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Personalidade

	database.DB.Delete(&p, id)

	json.NewEncoder(w).Encode(p)
}
