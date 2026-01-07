package main

import (
	"fmt"
	"os"
	"path"
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

	command := handlers.Command{
		Type: handlers.List,
		Data: map[string]string{
			"limit":  "1",
			"offset": "-1",
		},
	}
	result, err := controller.GetPersons(command)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("%s", string(result))
}
