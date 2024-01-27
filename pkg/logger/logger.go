package logger

import (
	"log"
	"os"
)

var Logger *log.Logger

func init() {
	Logger = log.New(os.Stdout, "UalaAPP: ", log.Ldate|log.Ltime|log.Lshortfile)
}
