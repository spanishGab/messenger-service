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
	_, err := os.OpenFile(fh.FileName, os.O_RDWR|os.O_CREATE, OwnerRWPermission)
	if err != nil {
		panic(fmt.Sprintf("failed to open file '%s': %s", fh.FileName, err))
	} 
}

func (fh *FileHandler) Read() ([]byte, error) {
	file, err := os.ReadFile(fh.FileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	return file, nil
}

func (fh *FileHandler) Write(data []byte) (int, error) {
	err := os.WriteFile(fh.FileName, data, OwnerRWPermission)
	if err != nil { 
		return 0, fmt.Errorf("failed to write to file: %w", err)
	}
	return len(data), nil
}
