package main

import (
	"fmt"
	"messenger-api/src/db"
	"messenger-api/src/entities"
	"messenger-api/src/repositories"
	"os"
	"path"
	"time"

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

	// // GetById
	// var id = uuid.MustParse("e6718f1b-d178-4f69-97a2-3b01b986fb3f")
	// message, err := repo.GetById(id)
	// if err != nil {
	// 	fmt.Println("failed to retrieve message with ID '%w': %w", id, err)
	// 	return 
	// }

	// output, _ := json.MarshalIndent(message, "", "  ")
	// fmt.Printf("Message found:\n%+v\n", string(output))

	// // GetMessages
	// count := int8(5)
	// filters := entities.Filters{
	// 	Content: "new",
	// 	CreatedAt: "2000-04-16",
	// 	TimesSent: &count,
	// }

	// message, err := repo.GetMessages(filters)
	// if err != nil {
	// 	fmt.Println("failed to retrieve message with content '%w': %w", filters.Content, err)
	// 	return 
	// }
	// output, _ := json.MarshalIndent(message, "", "  ")
	// fmt.Printf("Message found:\n%+v\n", string(output))

	// // DeleteMessages
	// var id = uuid.MustParse("0a5834b0-16b5-4a6a-b995-0292caace221")
	// err = repo.DeleteMessage(id)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// filters := entities.Filters{}
	// message, err := repo.GetMessages(filters)
	// if err != nil {
	// 	fmt.Println("failed to retrieve message with content '%w': %w", filters.Content, err)
	// 	return 
	// }
	// output, _ := json.MarshalIndent(message, "", "  ")
	// fmt.Printf("Messages found:\n%+v\n", string(output))

	// InsertMessage
	var id = uuid.New()
	createdAt := time.Now()

	message := entities.Message{
		Id: id,
		Content: "Buy boardgame",
		CreatedAt: createdAt,
		TimesSent: 2,
	}
	err = repo.InsertMessage(message)
	if err != nil {
		fmt.Println("failed to insert message '%w': %w", message)
		return 
	}
}
