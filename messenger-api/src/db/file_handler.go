package db

import (
	"fmt"
	"os"
)

const (
	OwnerRWPermission = 0644
)

type FileHandler struct {
	FileName string
}

func NewFileHandler(filename string) *FileHandler {
	return &FileHandler{
		FileName: filename,
	}
}

func (fh *FileHandler) Connection() {
	fmt.Printf("Arquivo %s\n", fh.FileName)

	_, err := os.OpenFile(fh.FileName, os.O_RDWR|os.O_CREATE, OwnerRWPermission)
	if err != nil {
		fmt.Printf("erro ao abrir arquivo %v\n", err)
	} 
	// defer file.Close()
}

func (fh *FileHandler) Read() ([]byte, error) {
	file, err := os.ReadFile(fh.FileName)
	if err != nil {
		return nil, fmt.Errorf("não foi possivel ler o arquivo %s", err.Error())
	}
	return file, nil
}

func (fh *FileHandler) Write(data []byte) (int, error) {
	err := os.WriteFile(fh.FileName, data, OwnerRWPermission)
	if err != nil { 
		return 0, fmt.Errorf("não foi possivel escrever no arquivo %s", err.Error())
	}
	return len(data), nil
}
