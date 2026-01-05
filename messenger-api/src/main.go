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
	// content := "message"

	// startDate, _ := time.Parse(shared.ShortDateFormat, "2025-12-11")
	// endDate, _ := time.Parse(shared.ShortDateFormat, "2025-12-12")

	// dateRange := entities.DateRange{
	// 	Start: startDate,
	// 	End: endDate,
	// }

	// timesSent := entities.TimesSent{
	// 	Value: uint8(2),
	// 	Operator: "=",
	// }

	// filters := entities.Filters{
	// 	Content: &content,
	// 	DateRange: &dateRange,
	// 	TimesSent: &timesSent,
	// }

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

	message, err := entities.NewMessage("Buy Spanish book", 3)
	if err != nil {
		fmt.Println("failed to create new message:", err)
		return 
	}

	if err := repo.InsertMessage(message); err != nil {
		fmt.Println("failed to insert message '%w':", message)
		return 
	}
}
