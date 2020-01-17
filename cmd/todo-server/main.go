package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"

	"github.com/someDevDude/todo-server/api"
	"github.com/someDevDude/todo-server/database"
	"github.com/someDevDude/todo-server/util"
)

func main() {
	util.InitialiseLogger()

	database.Connect()
	defer database.DB.Close()

	r := mux.NewRouter()

	api.CreateHomeRoutes(r)
	api.CreateHealthCheckRoutes(r)
	api.CreateTodoRoutes(r)

	http.ListenAndServe(":8080", r)
}
