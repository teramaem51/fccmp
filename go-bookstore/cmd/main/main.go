package main

import (
	"go-bookstore/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

/* quick tests on terminal ---

curl -X POST -H "Content-Type: application/json" -d '{"Name": "Zero to one","Author": "Peter Thiel","Publication": "Penguin"}' http://localhost:8080/book/ | jq '.'

curl -X POST -H "Content-Type: application/json" -d '{"Name": "The startup way","Author": "Eric Ries","Publication": "Penguin"}' http://localhost:8080/book/ | jq '.'

curl http://localhost:8080/book/ | jq '.'

curl http://localhost:8080/book/1 | jq '.'

curl -X PUT -H "Content-Type: application/json" -d '{"Name": "The startup way","Author": "Eric Ries","Publication": "Orion"}'  http://localhost:8080/book/1 | jq '.'

curl -X DELETE http://localhost:8080/book/2 | jq '.'

--- */
