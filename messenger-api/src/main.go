package main

import (
	"encoding/json"
	"fmt"
	"messenger-api/src/db"
	"messenger-api/src/repositories"
	"os"
	"path"

	"github.com/google/uuid"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("failed to get current working directory: %v", err)
		return
	}

	dbConnection := db.NewFileHandler(path.Join(cwd, "db", "message.json"))
	dbConnection.Connection()

	repo := repositories.NewMessageRepository(*dbConnection)

	var id = uuid.MustParse("e6718f1b-d178-4f69-97a2-3b01b986fb3f")

	message, err := repo.GetById(id)
	if err != nil {
		fmt.Println("failed to retrieve message with ID '%w': %w", id, err)
		return 
	}

	output, _ := json.MarshalIndent(message, "", "  ")
	fmt.Printf("Message found:\n%+v\n", string(output))
}