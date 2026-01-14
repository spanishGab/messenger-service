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

	// DeleteMessage 
	deleteMessage := handlers.Command{
		Type: handlers.Delete,
		Data: map[string]string{
			"id": "8849e12f-f6a6-4f8e-ad58-d50f2b0a443e",
		},
	}

	deleteResult, err := controller.DeleteMessage(deleteMessage)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return 
	}
	fmt.Printf("%s \n", string(deleteResult))

	// InsertMessage
	insertMessage := handlers.Command{
		Type: handlers.Create,
		Data: map[string]string{
			"content": "Buy a new planner",
			"timesSent": "2",
		},
	}

	insertResult, err := controller.InsertMessage(insertMessage)
		if err != nil {
		fmt.Printf("Error: %s\n", err)
		return 
	}
	fmt.Printf("%s \n", string(insertResult))

	// UpdateMessage
	updateMessage := handlers.Command{
		Type: handlers.Update,
		Data: map[string]string{
			"id": "e6718f1b-d178-4f69-97a2-3b01b986fb3f",
			"content": "Talk about money",
			"timesSent": "10",
		},
	}

	updateResult, err := controller.UpdateMessage(updateMessage)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return 
	}
	fmt.Printf("%s \n", string(updateResult))
}
