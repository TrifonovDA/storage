package storage

import (
	file2 "Storage/internal/file"
	"fmt"
	"github.com/google/uuid"
)

type Storage struct {
	Files map[uuid.UUID]*file2.File
}

func NewStorage() *Storage {
	return &Storage{
		Files: make(map[uuid.UUID]*file2.File),
	}
}

func (s *Storage) Upload(filename string, blob []byte) (*file2.File, error) {
	newFile, err := file2.NewFile(filename, blob)
	if err != nil {
		return nil, err
	}

	s.Files[newFile.ID] = newFile
	return newFile, nil
}

func (s *Storage) GetByID(FileID uuid.UUID) (*file2.File, error) {
	foundFile, ok := s.Files[FileID]
	if ok == false {
		return nil, fmt.Errorf("File %v not found", FileID)
	}

	return foundFile, nil
}
