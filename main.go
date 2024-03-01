// escrever pkgm e dar enter cria essa parada
package main

import (
	"api-rest/database"
	"api-rest/routes"
	"fmt"
)

func main() {
	database.ConnectDB()

	fmt.Println("teste")
	routes.HandleRequest()
}
