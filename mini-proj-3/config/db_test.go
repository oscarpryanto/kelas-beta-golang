package config_test

import (
	"fmt"
	"miniproject3/manajemen_buku_perpustakaan/config"
	"testing"

	"github.com/joho/godotenv"
)

func InitEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println(err)
	}

}

func TestConnection(t *testing.T) {
	InitEnv()
	config.OpenDb()

}
