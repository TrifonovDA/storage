package main

import (
	"fmt"
	"github.com/TrifonovDA/Storage/internal/storage"
	"log"
)

func main() {
	st := storage.NewStorage()
	file, err := st.Upload("test.txt", []byte("Hello."))
	if err != nil {
		log.Fatal(err)
	}
	restoredFile, err := st.GetByID(file.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("It is restored", restoredFile)

}
