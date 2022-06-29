package utils

import (
	"log"
	"os"
)

var (
	InfoLog    *log.Logger
	WarningLog *log.Logger
	ErrorLog   *log.Logger
)

func InitLogger() {
	flags := log.Default().Flags()
	InfoLog = log.New(os.Stdout, "[INFO]\t", flags|log.Llongfile)
	WarningLog = log.New(os.Stdout, "[WARNING]\t", flags|log.Llongfile)
	ErrorLog = log.New(os.Stdout, "[ERROR]\t", flags|log.Llongfile)
}
