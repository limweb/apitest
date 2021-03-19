package config

import (
	_ "embed"
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

var (

	//go:embed ..\config.ini
	j []byte

	cfg *ini.File
)

func Setup() {
	configini, err := ini.Load(j)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	cfg = configini
}

func GetConfig() *ini.File {
	return cfg
}
