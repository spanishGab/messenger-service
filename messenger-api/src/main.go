package main

import (
	"fmt"
	"messenger-api/src/db"
	"messenger-api/src/handlers"
	"messenger-api/src/repositories"
	"os"
	"path"
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
	controller := handlers.NewMessageHandle(*repo)

	// GetById
	listById := handlers.Command{
		Type: handlers.ListById,
		Data: map[string]string{
			"id": "b829bb89-01e3-4466-8138-452d8fbeaedf",
		},
	}

	messageById, err := controller.GetMessageById(listById)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return 
	}
	fmt.Printf("%s \n", string(messageById))

	// GetMessages
	list := handlers.Command{
		Type: handlers.List,
		Data: map[string]string{
			"content": "buy",
			"createdAtStart": "2025-12-11",
			"createdAtEnd": "2025-12-12",
			"timesSentValue": "2",
			"timesSentOperator": "=",
		},
	}

	messangeList, err := controller.GetMessages(list)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return 
	}
	fmt.Printf("%s \n", string(messangeList))

	// DeleteMessages
	deleteMessage := handlers.Command{
		Type: handlers.Delete,
		Data: map[string]string{
			"id": "8849e12f-f6a6-4f8e-ad58-d50f2b0a443e",
		},
	}

	result, err := controller.DeleteMessage(deleteMessage)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return 
	}
	fmt.Printf("%s \n", string(result))

	// // InsertMessage
	// message, err := entities.NewMessage("Buy Spanish book", 3)
	// if err != nil {
	// 	fmt.Println("failed to create new message:", err)
	// 	return 
	// }

	// if err := repo.InsertMessage(message); err != nil {
	// 	fmt.Println("failed to insert message '%w':", message)
	// 	return 
	// }

	// 

}
