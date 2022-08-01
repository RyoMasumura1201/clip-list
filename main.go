/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/RyoMasumura1201/clip-list/cmd"
)

func main() {
	curDir, _ := os.Getwd()
	filePath := filepath.Join(curDir, ".clipList")

	file, err := os.OpenFile(filePath, os.O_CREATE, 0666)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	cmd.Execute(filePath)
}
