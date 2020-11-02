package main

import (
	"log"
	"os"
)

func logToFile() {
	f, err := os.OpenFile("./app.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime)
	infoLog.Printf("New log message")
}
