package main

import (
	"Storage/internal/storage"
	"fmt"
)

func main() {
	st := storage.NewStorage()
	fmt.Println(st)
}
