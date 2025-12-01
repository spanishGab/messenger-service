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

func NewFileHandler(fileName string) *FileHandler {
	return &FileHandler{
		FileName: fileName,
	}
}

// CreateFile creates the file if it does not exist
func (fh *FileHandler) Connect() {
	_, err := os.OpenFile(fh.FileName, os.O_CREATE, 0666)
	if err != nil {
		panic(fmt.Sprintf("error while initializing the system: %s", err.Error()))
	}
}

func (fh *FileHandler) Write(p []byte) (int, error) {
	if err := os.WriteFile(fh.FileName, p, OwnerRWPermission); err != nil {
		return 0, fmt.Errorf("error while trying to write to file: %s", err.Error())
	}
	return len(p), nil
}

func (fh *FileHandler) Read() ([]byte, error) {
	var err error
	data, err := os.ReadFile(fh.FileName)
	if err != nil {
		return nil, fmt.Errorf("error while trying to read all file data: %s", err.Error())
	}
	return data, nil
}
