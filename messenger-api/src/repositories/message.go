package repositories

import (
	"encoding/json"
	"fmt"
	"messenger-api/src/db"
	"messenger-api/src/entities"

	"github.com/google/uuid"
)

// todo: db/message.json validar se ele suportar inteiro ou so string


type MessageRespository struct {
	dbConnection db.FileHandler
}

func NewMessageRepository(dbConnection db.FileHandler) *MessageRespository {
	return &MessageRespository{
		dbConnection: dbConnection,
	}
}

func (m *MessageRespository) GetById(id uuid.UUID) (*entities.Message ,error) {
	var messages []entities.Message

	file, err := m.dbConnection.Read()
	if err != nil {
		return nil, fmt.Errorf("não foi possivel ler o arquivo para obter os dados")
	}

	err = json.Unmarshal(file, &messages)
	if err != nil {
		return nil, fmt.Errorf("não foi possivel decodificar o json")
	}

	for _, message := range messages {
		if message.Id == id {
			return &message, nil
		}
	}
	return nil, fmt.Errorf("não foi possivel encontrar o id %s", id) 
}
