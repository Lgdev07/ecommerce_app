package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadDotEnv() {
	if _, err := os.Stat(".env"); !os.IsNotExist(err) {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error getting env, %v", err)
		} else {
			fmt.Println("We are getting values")
		}
	}
}
