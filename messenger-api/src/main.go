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


	// message, err := repo.GetMessages(filters)
	// if err != nil {
	// 	fmt.Println("failed to retrieve message with content '%w': %w", filters.Content, err)
	// 	return 
	// }
	// output, _ := json.MarshalIndent(message, "", "  ")
	// fmt.Printf("Message found:\n%+v\n", string(output))

	// // DeleteMessages
	// var id = uuid.MustParse("30bcc896-deba-44da-813c-5d52c9de42b9")
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
