package main

import (
	"fmt"
	"messenger-api/src/db"
	"messenger-api/src/repositories"
	"os"
	"path"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("Start main.go")

	cwd := os.Getenv("CWD")
	dbConnection := db.NewFileHandler(path.Join(cwd, "src", "db", "message.json"))
	dbConnection.Connection()

	repo := repositories.NewMessageRepository(*dbConnection)

	var id = uuid.MustParse("e6718f1b-d178-4f69-97a2-3b01b986fb3f")

	message, err := repo.GetById(id)
	if message == nil {
		fmt.Println("não foi possivel encontrar message id: %s. Error: ", id, err)
		return 
	}

	fmt.Println(message)
}