package main

import (
	"log"
	"os"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	InfoLogger = log.New(file, "INFO: ", log.LstdFlags|log.LstdFlags)
	WarningLogger = log.New(file, "WARNING: ", log.LstdFlags|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.LstdFlags|log.Lshortfile)

}

func main() {
	InfoLogger.Println("This is some info")
	WarningLogger.Println("THis is probably Important ")
	ErrorLogger.Println("TSomethign went wrong")
}

// func main() {
// 	// process("order")
// 	// process("refund")
// 	log.SetFlags(log.LstdFlags | log.Lshortfile)
// 	log.Println("This is a log message")
// 	// log.Fatal("This is a log message")
// 	// log.Panic("This is a log message")
//
// 	log.SetOutput(file)
// 	log.Println("This is my log")
// }

// func process(item string){
// 	for i:=1;true;i++{

// 	}
// }
