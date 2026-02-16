package main

import (
	"os"
	"path"
	"spanishGab/aula_camada_model/src/cmd"
	"spanishGab/aula_camada_model/src/db"
	"spanishGab/aula_camada_model/src/handlers"
	"spanishGab/aula_camada_model/src/repositories"
)

func main() {
	cwd := os.Getenv("CWD")
	dbConnection := db.NewFileHandler(path.Join(cwd, "src", "db", "persons.json"))
	dbConnection.Connect()

	repo := repositories.NewPersonRepository(*dbConnection)
	controller := handlers.NewPersonHandler(*repo)
	router := cmd.NewPersonRouter(*controller)

	router.Route(os.Args)
}
