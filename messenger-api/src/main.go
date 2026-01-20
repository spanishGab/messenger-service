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
			"id": "8f5eadea-c459-41d9-ae74-dc7164d9483b",
		},
	}

	messageById := controller.GetMessageById(listById)
	if messageById.Error != nil {
		fmt.Printf("Error: %s\n", messageById.Error)
		return 
	}
	fmt.Printf("%s \n", string(*messageById.Value))

	// GetMessages
	list := handlers.Command{
		Type: handlers.List,
		Data: map[string]string{
			"content": "buy",
			"createdAtStart": "2026-01-19",
			"createdAtEnd": "2026-01-19",
			"timesSentValue": "2",
			"timesSentOperator": ">=",
			"page": "3",
			"pageSize": "1",
		},
	}

	messangeList := controller.GetMessages(list)
	if messangeList.Error != nil {
		fmt.Printf("Error: %s\n", messangeList.Error)
		return 
	}
	fmt.Printf("%s \n", string(*messangeList.Value))

	// DeleteMessage 
	deleteMessage := handlers.Command{
		Type: handlers.Delete,
		Data: map[string]string{
			"id": "45f88db2-d973-44cf-abd4-4bda898d8ef6",
		},
	}

	deleteResult := controller.DeleteMessage(deleteMessage)
	if deleteResult.Error != nil {
		fmt.Printf("Error: %s\n", deleteResult.Error)
		return 
	}
	fmt.Printf("%s \n", string(*deleteResult.Value))

	// InsertMessage
	insertMessage := handlers.Command{
		Type: handlers.Create,
		Data: map[string]string{
			"content": "Buy cake",
			"timesSent": "2",
		},
	}

	insertResult := controller.InsertMessage(insertMessage)
	if insertResult.Error != nil {
		fmt.Printf("Error: %s\n", insertResult.Error)
		return 
	}
	fmt.Printf("%s \n", string(*insertResult.Value))

	// UpdateMessage
	updateMessage := handlers.Command{
		Type: handlers.Update,
		Data: map[string]string{
			"id": "e6718f1b-d178-4f69-97a2-3b01b986fb3f",
			"content": "Talk about money",
			"timesSent": "1",
		},
	}

	updateResult := controller.UpdateMessage(updateMessage)
	if updateResult.Error != nil {
		fmt.Printf("Error: %s\n", updateResult.Error)
		return 
	}
	fmt.Printf("%s \n", string(*updateResult.Value))
}
