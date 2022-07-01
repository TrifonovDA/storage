package storage

import (
	file2 "Storage/internal/file"
	//file "github.com/TrifonovDA/Storage/intenal/file"
)

type Storage struct{}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Upload(filename string, blob []byte) (*file2.File, error) {
	return file2.NewFile(filename, blob)
	//if err != nil {
	//	return nil, err
	//}
	//return newFile, nil
}
