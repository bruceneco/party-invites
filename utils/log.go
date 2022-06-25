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
	InfoLog = log.New(os.Stdout, "[INFO]\t", flags)
	WarningLog = log.New(os.Stdout, "[WARNING]\t", flags)
	ErrorLog = log.New(os.Stdout, "[ERROR]\t", flags)

}
