package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"spanishGab/aula_camada_model/src/db"
	"spanishGab/aula_camada_model/src/repositories"

	"github.com/google/uuid"
)

func main() {
	cwd := os.Getenv("CWD")
	dbConnection := db.NewFileHandler(path.Join(cwd, "src", "db", "persons.json"))
	dbConnection.Connect()

	repo := repositories.NewPersonRepository(*dbConnection)
	person, err := repo.GetById(uuid.MustParse("93dbd07d-f050-447a-9ad7-ec52b8741c3f"))
	if person == nil {
		fmt.Println("Person not found", err)
		return
	}
	output, _ := json.MarshalIndent(person, "", "  ")
	fmt.Printf("%s", string(output))
}
