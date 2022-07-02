package storage

import (
	"fmt"
	file "github.com/TrifonovDA/Storage/internal/file"
	"github.com/google/uuid"
)

type Storage struct {
	Files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		Files: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	newFile, err := file.NewFile(filename, blob)
	if err != nil {
		return nil, err
	}

	s.Files[newFile.ID] = newFile
	return newFile, nil
}

func (s *Storage) GetByID(FileID uuid.UUID) (*file.File, error) {
	foundFile, ok := s.Files[FileID]
	if ok == false {
		return nil, fmt.Errorf("File %v not found", FileID)
	}

	return foundFile, nil
}
