package store

import (
	"encoding/json"
	"os"

	"github.com/batatinha123/products-api/internal/entities"
)

type FileStoreProductsStub struct {
	FileName string
}

func (fs *FileStoreProductsStub) Read(data interface{}) error {
	var products []entities.Product
	jsonData, _ := json.Marshal(products)
	return json.Unmarshal(jsonData, &data)
}

func (fs *FileStoreProductsStub) Write(data interface{}) error {
	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FileName, fileData, 0644)
}

func NewFileStoreProductsStub(store Type, fileName string) Store {
	switch store {
	case FileType:
		return &FileStore{fileName}
	}
	return nil
}
